package mongdb

import (
	"contract/internal/config"
	"contract/internal/util"
	"contract/pkgs/nft"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NftClient 以太坊客户端2
type NftClient struct {
	Client *ethclient.Client
	FffNft *nft.FFFNFTTransactor
}

// Config redis config
type Config struct {
	Address string
}

// NewEntity 实例化 ethClient
func NewEntity(cfg Config) (*NftClient, error) {
	client, err := ethclient.Dial(cfg.Address)
	if err != nil {
		util.Logger.WithError(err).Error(err.Error())
		return nil, nil
	}
	newNftTrac, err := nft.NewFFFNFTTransactor(common.HexToAddress(config.Setting.Chain.NftContractAddress), client)
	if err != nil {
		util.Logger.WithError(err).Error(err.Error())
		return nil, nil
	}

	return &NftClient{
		Client: client,
		FffNft: newNftTrac,
	}, nil
}
