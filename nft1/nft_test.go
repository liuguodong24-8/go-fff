package nft

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
	"time"
)

func TestFFFNFTTransactor_Mint(t *testing.T) {

	client, err := ethclient.Dial("http://47.109.29.166:8488")
	if err != nil {
		log.Panicln(err)
	}

	nftContractOwnKey := "df7abadcc169e4a2a8ffdd32852f6fb34da033ed4abfa01fddecdc7aad9f0ec2" //NFT合约拥有者私钥
	nftContractAddress := common.HexToAddress("0x7aD8b67FAb5E77da46D23CcFc54C95c82927eFDf") //NFT合约地址
	priKeyECD, err := crypto.HexToECDSA(nftContractOwnKey)
	if err != nil {
		log.Println("私钥异常", err)
		return
	}
	newNftTrac, err := NewFFFNFTTransactor(nftContractAddress, client)
	if err != nil {
		log.Println(err)
		return
	}
	chainId, _ := client.ChainID(context.Background())
	param, _ := bind.NewKeyedTransactorWithChainID(priKeyECD, chainId)

	toAddress := "0x1D6d8c0ec7FF4E03d0Be661e5007B177daB9F69d" //给谁创建NFT

	nftName := "NFT111"                                                                                                             //NFT名字
	nftContent := "NFT介绍"                                                                                                           //NFT介绍
	nftImageUrl := "https://dgss0.bdstatic.com/5bVWsj_p_tVS5dKfpU_Y_D3/res/r/image/2017-09-27/297f5edb1e984613083a2d3cc0c5bb36.png" //NFT图片地址

	nftCertList := []string{"https://dgss0.bdstatic.com/5bVWsj_p_tVS5dKfpU_Y_D3/res/r/image/2017-09-27/297f5edb1e984613083a2d3cc0c5bb36.png"} //nft证书地址
	hash, err := newNftTrac.TransferFrom(param, common.HexToAddress(toAddress), nftName, nftContent, nftImageUrl, nftCertList)

	if err == nil {
		log.Println("创建NFT哈希已创建:", hash.Hash().Hex())
	} else {
		log.Println("创建NFT失败", err)
		return
	}

	var sumWait = 0
	var nftId = big.NewInt(0)
	for {

		recpt, err := client.TransactionReceipt(context.Background(), hash.Hash())
		if err == nil && recpt != nil && recpt.Status == 1 {
			nftId = recpt.Logs[0].Topics[3].Big()
			log.Println("创建NFT成功:", hash.Hash().Hex(), "确认，当前状态", recpt.Status, "nftId", nftId)
			break
		}
		if recpt == nil {
			log.Println("等待创建NFT:", hash.Hash().Hex(), "确认，当前状态", 0)
		} else {
			log.Println("等待创建NFT:", hash.Hash().Hex(), "确认，当前状态", recpt.Status)
		}
		sumWait++

		if sumWait > 60 {

			log.Println("超时未确认创建NFT失败")
			break
		}
		time.Sleep(time.Second * 2)
	}
	if sumWait <= 60 {
		log.Println("创建NFT成功", "ID:", nftId)

	}
}
