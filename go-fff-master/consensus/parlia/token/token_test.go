package token

import (
	"context"
	"github.com/fff-chain/go-fff/accounts/abi/bind"
	"github.com/fff-chain/go-fff/common"
	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/ethclient"
	"github.com/fff-chain/go-fff/global_config"
	"github.com/fff-chain/go-fff/global_config/wallet"
	"log"
	"math/big"
	"testing"
)

func TestDeployBEP20USDT(t *testing.T) { //发布代币

	d, _ := ethclient.Dial("www.3fchain.org:8489")

	pri := "c862dfb65d03054a5d98cd1cb02d0697d132effe7d5b5046fb8416072cb68554"

	priKeyECD, err := crypto.HexToECDSA(pri)

	ownerAddress := wallet.GetPublicAddressFromKey(pri)

	if err != nil {
		log.Println("私钥异常")

	}
	chianId, err := d.ChainID(context.Background())
	if err != nil {
		log.Println("节点chianId异常", err)

	}
	param, err := bind.NewKeyedTransactorWithChainID(priKeyECD, chianId)
	if err != nil {
		log.Println("节点Transactor异常", err)

	}

	addr, _, _, _ := DeployBEP20USDT(param, d, "DEUTSCHE MARK COIN CRYPTOCURRENCY", "DMC", global_config.EthToWei(100_0000_0000), "")

	log.Println(addr, ownerAddress)

}

func TestErc20Transfer(t *testing.T) { //代币转账

	d, _ := ethclient.Dial("http://87.118.86.2:8488")

	pri := "20a442166fda1598b32d85d0cf5b30fe5867bf23c943f8584013eab0ad49a88c"

	priKeyECD, err := crypto.HexToECDSA(pri)

	if err != nil {
		log.Println("私钥异常")

	}
	chianId, err := d.ChainID(context.Background())
	if err != nil {
		log.Println("节点chianId异常", err)

	}
	param, err := bind.NewKeyedTransactorWithChainID(priKeyECD, chianId)
	if err != nil {
		log.Println("节点Transactor异常", err)

	}

	tokenContract := common.HexToAddress("FFF3TKXUqGcZ8wnFi5RGVMXSkrujbWzWAmXnWtvQSQSRnMS9kNpMCx7ig8")

	tokenTransact, err := NewBEP20USDTTransactor(tokenContract, d)
	if err != nil {
		log.Println("节点Transactor异常", err)

	}

	ts, err := tokenTransact.Transfer(param, common.HexToAddress("FFF3rPN3k7M1EASWb6QU1QvJHXmrUGhPF8QS5U7hpDfkoWwGyb6aBoZhD1"), global_config.EthToWei(1))

	if err != nil {
		log.Println("节点Transactor异常", err)

	}

	log.Println(ts.Hash())

}
func TestGetErc20ContractInfo(t *testing.T) { //获取代币信息

	d, _ := ethclient.Dial("http://87.118.86.2:8488")

	ctAddress := common.HexToAddress("FFF3bU7YUTH4Pn6PqvzxZ5GeVHDWJwKYngkVX4CsCQUBqs9q8SjVgLE473")

	cl, _ := NewBEP20USDTCaller(ctAddress, d)

	name, _ := cl.Name(nil)

	log.Println("代币全民", name)

	symbol, _ := cl.Symbol(nil)

	log.Println("代币简称", symbol)

	Decimals, _ := cl.Decimals(nil)

	log.Println("代币精度", Decimals)

	total, _ := cl.TotalSupply(nil)

	log.Println("代币总流通", total)

	bl, _ := cl.BalanceOf(nil, common.HexToAddress("FFF3nx7A1Coqe5shVeQzSbC5TruewaDkWVScqrsES2dB8FsJtfcYDCvJ6c"))

	log.Println(global_config.WeiToEth(bl))

	log.Println(d.PendingBalanceAt(context.Background(), common.HexToAddress("FFF3nx7A1Coqe5shVeQzSbC5TruewaDkWVScqrsES2dB8FsJtfcYDCvJ6c")))

}
func TestGetErc20Transaction(t *testing.T) { //获取每个交易中包含的代币交易

	d, _ := ethclient.Dial("http://87.118.86.2:8488")

	log.Println(getTransactionErc20Transactions(d, "0xa6438cfe68f349743b9d7ed85df14510bea7938f36a7b1b467a124d1a6efbbbf"))

}

type Erc20Transaction struct {
	From            common.Address
	To              common.Address
	Value           *big.Int
	ContractAddress common.Address
}

func getTransactionErc20Transactions(d *ethclient.Client, tsHash string) []Erc20Transaction {
	var erc20Ts []Erc20Transaction
	rc, err := d.TransactionReceipt(context.Background(), common.HexToHash(tsHash))
	if err != nil {

		log.Println(err)
		return erc20Ts
	}
	logs := rc.Logs

	var ercTransferHash = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	for k := range logs {

		topics := logs[k].Topics

		if len(topics) > 2 && topics[0].Hex() == ercTransferHash.Hex() { //判断是不是ERC20转账

			newTrans := new(Erc20Transaction)
			newTrans.From = common.BytesToAddress(topics[1].Bytes())
			newTrans.To = common.BytesToAddress(topics[2].Bytes())
			newTrans.Value = new(big.Int).SetBytes(logs[k].Data)
			newTrans.ContractAddress = logs[k].Address
			erc20Ts = append(erc20Ts, *newTrans)
		}

	}

	return erc20Ts

}
