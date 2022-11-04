package contract

import (
	"context"
	"contract/internal/config"
	"contract/internal/service"
	"contract/internal/util"
	"contract/pkgs/api"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"math/big"
	"regexp"
	"sync"
	"time"
)

// MintNft 创建nft
func MintNft(c *gin.Context) {
	var request struct {
		ID          int64    `json:"id" form:"id" binding:"required"`
		NftName     string   `json:"nft_name" form:"nft_name" binding:"required"`
		NftContent  string   `json:"nft_content" form:"nft_content" binding:"required"`
		NftImageUrl string   `json:"nft_image_url" form:"nft_image_url" binding:"required"`
		NftCertList []string `json:"nft_cert_list" form:"nft_cert_list" binding:"required"`
		Type        int      `json:"type" form:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("新增参数错误:%v", err))
		api.Unprocessable(c, fmt.Sprintf("新增参数错误:%v", err))
		return
	}

	client := service.NftEntity.Client
	newNftTrac := service.NftEntity.FffNft
	priKeyECD, err := crypto.HexToECDSA(config.Setting.Chain.NftContractOwnKey)
	if err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("私钥异常:%v", err))
		api.ServerError(c, fmt.Sprintf("私钥异常:%v", err))
		return
	}

	chainId, _ := client.ChainID(c)
	param, _ := bind.NewKeyedTransactorWithChainID(priKeyECD, chainId)

	hash, err := newNftTrac.Mint(param, common.HexToAddress(config.Setting.Chain.ToAddress), request.NftName, request.NftContent, request.NftImageUrl, request.NftCertList)
	if err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("创建NFT失败:%v", err))
		api.ServerError(c, fmt.Sprintf("创建NFT失败:%v", err))
		return
	}

	var wg sync.WaitGroup
	ch := make(chan int)
	ch1 := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	//defer cancel()

	wg.Add(1)
	go service.GetTransactionReceipt(ctx, client, hash, ch, ch1, &wg, request.ID, "创建")

	for {
		select {
		case <-ch1:
			close(ch)
			close(ch1)
			util.Logger.WithError(err).Error(fmt.Sprintf("创建NFT超时:%v", err))
			api.ServerError(c, "创建NFT超时")
			return
		case <-ch:
			close(ch)
			close(ch1)
			api.Success(c, "创建NFT哈希已创建")
			return
		}
	}
	wg.Wait()
}

// TransNft 体提现nft
func TransNft(c *gin.Context) {
	var request struct {
		Address string   `json:"address" form:"address" binding:"required"`
		NftId   *big.Int `json:"nft_id" form:"nft_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("新增参数错误:%v", err))
		api.Unprocessable(c, fmt.Sprintf("新增参数错误:%v", err))
		return
	}
	client := service.NftEntity.Client
	newNftTrac := service.NftEntity.FffNft
	priKeyECD, err := crypto.HexToECDSA(config.Setting.Chain.NftContractOwnKey)
	if err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("私钥异常:%v", err))
		api.ServerError(c, fmt.Sprintf("私钥异常:%v", err))
		return
	}
	chainId, _ := client.ChainID(c)
	param, _ := bind.NewKeyedTransactorWithChainID(priKeyECD, chainId)

	trans, err := newNftTrac.TransferFrom(param, common.HexToAddress(config.Setting.Chain.ToAddress), common.HexToAddress(request.Address), request.NftId)
	if err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("提现NFT失败:%v", err))
		api.ServerError(c, fmt.Sprintf("提现NFT失败:%v", err))
		return
	}

	var wg sync.WaitGroup
	ch := make(chan int)
	ch1 := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	//defer cancel()

	wg.Add(1)
	go service.GetTransactionReceipt(ctx, client, trans, ch, ch1, &wg, 0, "交易")

	for {
		select {
		case <-ch1:
			close(ch)
			close(ch1)
			util.Logger.WithError(err).Error(fmt.Sprintf("提现NFT超时:%v", err))
			api.ServerError(c, "提现NFT超时")
			return
		case <-ch:
			close(ch)
			close(ch1)
			api.Success(c, "提现NFT成功")
			return
		}
	}
	wg.Wait()
}

// IsValidHexAddress 判断是否是有效地址
func IsValidHexAddress(c *gin.Context) {
	var request struct {
		Address string `json:"address" form:"address" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("验证参数错误:%v", err))
		api.Unprocessable(c, fmt.Sprintf("验证参数错误:%v", err))
		return
	}
	var addressPattern = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	var zeroHash = regexp.MustCompile("^0?x?0+$")

	if zeroHash.MatchString(request.Address) || !addressPattern.MatchString(request.Address) {
		util.Logger.Error(fmt.Sprintf("验证失败，不是有效地址"))
		api.ServerError(c, "验证失败，不是有效地址")
		return
	}
	api.Success(c, "成功")
}
