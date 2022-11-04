package wallet

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/fff-chain/go-fff/accounts/abi/bind"
	"github.com/fff-chain/go-fff/accounts/keystore"
	"github.com/fff-chain/go-fff/common"
	"github.com/fff-chain/go-fff/common/hexutil"
	math2 "github.com/fff-chain/go-fff/common/math"
	"github.com/fff-chain/go-fff/core/types"
	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/ethclient"
	"github.com/fff-chain/go-fff/global_config"
	"github.com/fff-chain/go-fff/global_config/token"
	"log"
	"math/big"
	"time"
)

var WEI = math2.BigPow(10, 9)

type Wallet struct {
	CurrClient *ethclient.Client
	RpcUrl     string
	RpcUrl2    string
}

func Init(RpcUrl string, RpcUrl2 string) *Wallet {
	newl := new(Wallet)
	newl.RpcUrl = RpcUrl
	newl.RpcUrl2 = RpcUrl2
	client, err := ethclient.Dial(RpcUrl)
	if err != nil || GetClientChainId(client) == 0 {
		log.Println("无法连接", RpcUrl, err, "等待五秒后重连")
		time.Sleep(time.Second * 5)
		client, err := ethclient.Dial(RpcUrl2)
		if err != nil || GetClientChainId(client) == 0 {
			log.Println("无法连接", RpcUrl2, err, "等待五秒后重连")
			time.Sleep(time.Second * 5)
			return Init(RpcUrl, RpcUrl2)
		} else {
			newl.CurrClient = client
			return newl
		}
	} else {
		newl.CurrClient = client
		return newl
	}
}
func (this *Wallet) ReInit() {
	for {
		client, err := ethclient.Dial(this.RpcUrl)
		if err != nil || GetClientChainId(client) == 0 {
			log.Println("无法连接", this.RpcUrl, err, "等待五秒后重连")
			time.Sleep(time.Second * 5)
			client, err := ethclient.Dial(this.RpcUrl2)
			if err != nil || GetClientChainId(client) == 0 {
				log.Println("无法连接", this.RpcUrl2, err, "等待五秒后重连")
				time.Sleep(time.Second * 5)
			} else {
				this.CurrClient = client
				return
			}
		} else {
			this.CurrClient = client
			return
		}
	}

}
func GetClientChainId(c *ethclient.Client) int64 {
	if c == nil {
		return 0
	}

	chainIdInt, err := c.ChainID(context.Background())
	if err != nil {
		log.Println(err)
		return 0
	}
	return chainIdInt.Int64()
}
func (this *Wallet) GetMiniBalance() float64 {
	gasLimit := uint64(21000)

	gasPrice, err2 := this.CurrClient.SuggestGasPrice(context.Background())

	if err2 != nil {
		gasPrice = big.NewInt(100)
	}
	count := new(big.Int).Mul(big.NewInt(int64(gasLimit)), gasPrice)
	return global_config.WeiToEth(count)

}
func FloatToWei(f float64) *big.Int {

	weiCount := 18

	beth := math2.BigPow(10, int64(weiCount))

	rs := new(big.Float).Mul(big.NewFloat(f), new(big.Float).SetInt(beth))

	w := big.NewInt(0)

	rs.Int(w)

	return w

}
func NewAddress() (pri string, pub string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Println(err)
		return "", ""
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	key := hexutil.Encode(privateKeyBytes)[2:]
	return key, GetPublicAddressFromKey(key)
}

func GetPublicAddressFromKey(userKey string) string {
	priKeyECD, err := crypto.HexToECDSA(userKey)
	if err != nil {
		log.Println("私钥异常")
		return ""
	}

	publicKeyECDSA, ok := priKeyECD.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return ""
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	if err != nil {
		log.Println("私钥异常")
		return ""
	}

	return address

}

func (this *Wallet) KeystoreToPrivateKey(keyJson, password string) (string, string, error) {
	/*keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		fmt.Println("read keyjson file failed：", err)
	}*/
	unlockedKey, err := keystore.DecryptKey([]byte(keyJson), password)
	if err != nil {
		return "", "", err
	}
	privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return privKey, addr.String(), nil
}

func (this *Wallet) GetTokenBalance(tokenAddress string, own string) *big.Int {

	transtBnB, err := token.NewBEP20TokenCaller(common.HexToAddress(tokenAddress), this.CurrClient)
	if err != nil {
		return nil
	}
	currBnb, err := transtBnB.BalanceOf(nil, common.HexToAddress(own))
	if err != nil {
		return nil
	}
	return currBnb

}
func (this *Wallet) TransferToken(pri string, tokenAddress string, toAddress string, count *big.Int) string {
	transt, err := token.NewBEP20TokenTransactor(common.HexToAddress(tokenAddress), this.CurrClient)

	if err == nil {
		priKeyECD, err := crypto.HexToECDSA(pri)

		if err != nil {
			log.Println("私钥异常")
			return ""
		}
		chianId, err := this.CurrClient.ChainID(context.Background())
		if err != nil {
			log.Println("节点chianId异常", err)
			return ""
		}
		param, err := bind.NewKeyedTransactorWithChainID(priKeyECD, chianId)
		if err != nil {
			log.Println("节点Transactor异常", err)
			return ""
		}
		param.GasPrice, err = this.CurrClient.SuggestGasPrice(context.Background())
		if err != nil {
			log.Println("GAS异常", err)
		}

		if param.GasPrice == nil {
			param.GasPrice = big.NewInt(1000000000)
		}

		if param.GasPrice.Cmp(big.NewInt(1000000000)) < 0 {
			param.GasPrice = big.NewInt(1000000000)
		}
		log.Println("当前gas:", param.GasPrice)

		param.GasLimit = uint64(8000000)
		nonce, err := this.CurrClient.PendingNonceAt(context.Background(), common.HexToAddress(GetPublicAddressFromKey(pri)))
		if err != nil {
			log.Println("节点异常")
			return ""
		}
		param.Nonce = big.NewInt(int64(nonce))

		ts, err := transt.Transfer(param, common.HexToAddress(toAddress), count) //转账代币

		if err != nil {
			log.Println("转账代币异常")
			return ""
		}
		return ts.Hash().Hex()

	} else {
		log.Println("节点异常")
		return ""
	}
}
func (this *Wallet) Transfer(privateKey string, toAddress string, count *big.Int, transData string, gasP int64) string {

	if this.CurrClient == nil {
		return ""
	}

	priviteKey, err := crypto.HexToECDSA(privateKey)

	publicAddress := GetPublicAddressFromKey(privateKey)

	if len(publicAddress) == 0 {
		return ""
	}

	if err != nil {
		log.Println("chanid err", err)
		return ""
	}

	nonce, err := this.CurrClient.PendingNonceAt(context.Background(), common.HexToAddress(publicAddress))
	if err != nil {
		log.Println("err", err)
		return ""
	}
	gasLimit := uint64(21000)

	gasPrice, err2 := this.CurrClient.SuggestGasPrice(context.Background())

	if err2 != nil {
		log.Println("err", err2)
		return ""
	}

	if gasP > 0 {
		gasPrice = new(big.Int).Mul(big.NewInt(int64(gasP)), math2.BigPow(10, 9))
	}

	count = new(big.Int).Sub(count, new(big.Int).Mul(big.NewInt(int64(gasLimit)), gasPrice))
	data := []byte(transData)

	newTrans := types.NewTransaction(nonce, common.HexToAddress(toAddress), count, gasLimit, gasPrice, data)

	chainID, err := this.CurrClient.ChainID(context.Background())
	if err != nil {
		log.Println("chanid err", err)
		return ""
	}

	signer := types.NewEIP155Signer(chainID)

	signTx, err := types.SignTx(newTrans, signer, priviteKey)

	if err != nil {
		log.Println("sign err", err)
		return ""
	}

	errTrans := this.CurrClient.SendTransaction(context.Background(), signTx)

	if errTrans != nil {
		log.Println("err", errTrans)

		return ""
	}
	return signTx.Hash().String()
}
func (this *Wallet) GetBalance(address string) *big.Int {
	if this.CurrClient == nil {
		return nil
	}
	b, err := this.CurrClient.BalanceAt(context.Background(), common.HexToAddress(address), nil)

	if err != nil || b == nil {
		return nil
	}
	return b

}
func (this *Wallet) GetChainId() int64 {
	chainID, err := this.CurrClient.NetworkID(context.Background())
	if err != nil {
		log.Println("chanid err", err)
		return 0
	}
	return chainID.Int64()

}
func (this *Wallet) GetCurrGas() int64 {
	if this.CurrClient == nil {
		return 0
	}

	gasPrice, err2 := this.CurrClient.SuggestGasPrice(context.Background())
	if err2 != nil {
		log.Println(err2)
		return 0
	}
	gasInt := new(big.Int).Quo(gasPrice, WEI)
	return gasInt.Int64()

}
