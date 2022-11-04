package service

import (
	"blockExplore/internal/config"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

var EthClientEntity rpc.Client

// SetEthClient 设置eth
func SetEthClient() error {
	// 获取链上块
	clientRpc, err := rpc.Dial(config.Setting.Chain.Address)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return err
	}
	EthClientEntity = *clientRpc
	return nil
}
