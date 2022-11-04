package nft

import (
	"context"
	"fmt"
	"github.com/fff-chain/go-fff/accounts/abi/bind"
	"github.com/fff-chain/go-fff/common"
	"github.com/fff-chain/go-fff/crypto"
	"github.com/fff-chain/go-fff/ethclient"
	"github.com/fff-chain/go-fff/global_config/wallet"
	"log"
	"testing"
)

func TestDeployAddress(t *testing.T) {

	d, _ := ethclient.Dial("http://87.118.86.2:8488")

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

	param.GasPrice, _ = d.SuggestGasPrice(context.Background())

	addr, _, _, _ := DeployFFFNFT(param, d)

	log.Println(addr, ownerAddress)

}

func TestFFFNFTTransactor_Mint(t *testing.T) {
	d, _ := ethclient.Dial("http://87.118.86.2:8488")

	pri := "c862dfb65d03054a5d98cd1cb02d0697d132effe7d5b5046fb8416072cb68554"

	priKeyECD, err := crypto.HexToECDSA(pri)

	ownerAddress := wallet.GetPublicAddressFromKey(pri)

	contractAddress := common.HexToAddress("FFF3k6ur4Vp4Qbe4iXjg2C1iLaNvPWRW1QjXHMLWJcp9zPrtFTHftLstRy")

	cl, _ := NewFFFNFTTransactor(contractAddress, d)
	chianId, err := d.ChainID(context.Background())
	if err != nil {
		log.Println("节点chianId异常", err)

	}
	param, err := bind.NewKeyedTransactorWithChainID(priKeyECD, chianId)
	if err != nil {
		log.Println("节点Transactor异常", err)

	}

	param.GasPrice, _ = d.SuggestGasPrice(context.Background())

	z, _ := cl.Mint(param, common.HexToAddress(ownerAddress), "1", "1", "1", nil)

	log.Println(z.Hash())
}

func TestFFFNFTCaller_GetApproved(t *testing.T) {

	d, _ := ethclient.Dial("http://87.118.86.2:8488")

	rc, _ := d.TransactionReceipt(context.Background(), common.HexToHash("0x318ea2d9e2af496e2b45d9f2cbf98eeab365375373251c400f65c75b9129268c"))

	fmt.Println(rc)
	//log.Println(rc.BlockNumber)

}
func TestDeployEnumerableMap(t *testing.T) {
	pri := "218ec8f94b29a96d38a21b19e0b866faf9963f48a1a0d575d91990fdf08f8936"

	ownerAddress := wallet.GetPublicAddressFromKey(pri)

	log.Println(ownerAddress)

}
