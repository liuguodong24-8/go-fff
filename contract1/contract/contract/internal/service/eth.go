package service

import (
	"context"
	"contract/internal/config"
	"contract/internal/util"
	ethClient "contract/pkgs/eth"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"sync"
	"time"
)

var NftEntity ethClient.NftClient

// SetEth 设置meth
func SetEth() error {
	ethClients, err := ethClient.NewEntity(ethClient.Config{
		Address: config.Setting.Chain.Address,
	})

	if err != nil {
		return err
	}
	NftEntity = *ethClients
	return nil
}

// GetTransactionReceipt 获取交易状态
func GetTransactionReceipt(ctx context.Context, client *ethclient.Client, hash *types.Transaction, ch chan int, ch1 chan int, wg *sync.WaitGroup, id int64, str string) {
	defer wg.Done()
loop:
	for {
		var nftId = big.NewInt(0)
		receipt, err := client.TransactionReceipt(ctx, hash.Hash())
		if err == nil && receipt != nil && receipt.Status == 1 {
			nftId = receipt.Logs[0].Topics[3].Big()

			//fmt.Println(nftId)

			switch id {
			case 0:

			default:
				//err = postParamToJavaInsertNft(id, nftId)
			}
			/*bytesData, _ := json.Marshal(map[string]interface{}{
				"id":               id,
				"nftId":            nftId,
				"nftContract":      config.Setting.Chain.NftContractAddress,
				"platformHoldAddr": config.Setting.Chain.ToAddress,
				"createAddress":    config.Setting.Chain.ToAddress,
			})
			err = util.Post(bytesData)*/

			/*if err != nil {
				util.Logger.WithError(err).Error(fmt.Sprintf("请求nft保存接口错误:%s", err))
				log.Println(fmt.Sprintf("请求nft保存接口错误:%s", err))
			}*/

			util.Logger.Info(fmt.Sprintf("%sNFT成功:%s 确认，当前状态:%d nftId:%s", str, hash.Hash().Hex(), receipt.Status, nftId))
			ch <- 1
			break loop
		}
		if receipt == nil {
			log.Println(fmt.Sprintf("等待%sNFT:%s 确认，当前状态: %d", str, hash.Hash().Hex(), 0))

			util.Logger.WithError(err).Error(fmt.Sprintf("等待%sNFT:%s 确认，当前状态: %d", str, hash.Hash().Hex(), 0))
		} else {
			log.Println(fmt.Sprintf("等待%sNFT:%s 确认，当前状态: %d", str, hash.Hash().Hex(), receipt.Status))

			util.Logger.WithError(err).Error(fmt.Sprintf("等待%sNFT:%s 确认，当前状态: %d", str, hash.Hash().Hex(), receipt.Status))
		}

		select {
		case <-ctx.Done():
			ch1 <- 1
			break loop
		default:
		}
		time.Sleep(time.Second * 2)
	}
}

// postParamToJavaInsertNft 新增链上传信息到java
func postParamToJavaInsertNft(id int64, nftId *big.Int) error {
	bytesData, _ := json.Marshal(map[string]interface{}{
		"id":               id,
		"nftId":            nftId,
		"nftContract":      config.Setting.Chain.NftContractAddress,
		"platformHoldAddr": config.Setting.Chain.ToAddress,

		"createAddress": config.Setting.Chain.ToAddress,
	})
	return util.Post(bytesData)
}

// postParamToJavaTransform 交易上传信息到java
func postParamToJavaTransform(id int64, nftId big.Int) error {
	bytesData, _ := json.Marshal(map[string]interface{}{
		"id":               id,
		"nftId":            nftId,
		"nftContract":      config.Setting.Chain.NftContractAddress,
		"platformHoldAddr": config.Setting.Chain.ToAddress,
		"createAddress":    config.Setting.Chain.ToAddress,
	})
	return util.Post(bytesData)
}
