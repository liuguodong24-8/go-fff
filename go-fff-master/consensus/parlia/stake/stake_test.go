package stake

import (
	"context"
	"fmt"
	"github.com/fff-chain/go-fff/accounts/abi/bind"
	"github.com/fff-chain/go-fff/common"
	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/ethclient"
	"github.com/fff-chain/go-fff/global_config"
	"github.com/fff-chain/go-fff/global_config/utils"
	"github.com/fff-chain/go-fff/global_config/wallet"
	"log"
	"math/big"
	"testing"
)

func TestDeploySafeMath(t *testing.T) {

	d, _ := ethclient.Dial("http://87.118.86.2:8489")
	/*d, _ := ethclient.Dial("http://87.118.86.2:8489")

	cl3, _ := NewBSCValidatorSetCaller(common.HexToAddress("FFF3Psbq3enwAmwXGa2QejWFdd9AwV1rczE6w1GPzs6WTPmJpKbmWiBrcX"), d)

	va, _ := cl3.GetAllStakeInfo(nil)

	log.Println(va)*/
	var currSubmit = 0

	cl2, _ := NewBSCValidatorSetTransactor(common.HexToAddress("FFF3Psbq3enwAmwXGa2QejWFdd9AwV1rczE6w1GPzs6WTPmJpKbmWiBrcX"), d)
	nonce, err := d.PendingNonceAt(context.Background(), common.HexToAddress("FFF3nyQJTh2PohN4Z1gBHXZ8vTLaoJcVJRyteoYEC4ZUToAov8KWPGgbxV"))
	if err != nil {
		log.Println("节点异常")
	}

	fmt.Println(nonce)

	for {
		pri := "0954ab7f6009370daf31a57c364b3251760dc9779058a40a532884043444f8ce"
		priKeyECD, err := crypto.HexToECDSA(pri)

		log.Println(wallet.GetPublicAddressFromKey(pri))

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

		param.Value = global_config.EthToWei(1)

		param.GasPrice, _ = d.SuggestGasPrice(context.Background())

		param.Nonce = big.NewInt(int64(nonce))

		log.Println(param)

		trans, err := cl2.StakeFFF(param)

		if err != nil {
			log.Println(err)
			continue
		}

		log.Println(trans.Hash(), err)
		currSubmit++
		nonce++

		if currSubmit > 0 {
			break
		}
	}
	//time.Sleep(time.Second * 1)

	TestDeploySafeMath(t)

}
func TestNewTransferHelperCaller2(t *testing.T) {
	log.Println(BSCValidatorSetABI)

}
func TestDeployTransferHelper(t *testing.T) {

}
func TestNewTransferHelperCaller(t *testing.T) {
	d, _ := ethclient.Dial("http://47.109.29.166:8488")

	bf, err := d.PendingBalanceAt(context.Background(), common.HexToAddress(wallet.GetPublicAddressFromKey("20a442166fda1598b32d85d0cf5b30fe5867bf23c943f8584013eab0ad49a88c")))

	log.Println(bf, err)

}
func TestNewBSCValidatorSet(t *testing.T) {
	log.Println(utils.FFFAddressEncode("0xea674fdde714fd979de3edf0f56aa9716b898ec8"))

}
func TestNewBSCValidatorSetCaller(t *testing.T) {
	log.Println(wallet.NewAddress())
}
