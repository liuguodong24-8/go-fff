package data

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type TrxGrpcClient struct {
	*client.GrpcClient
}

func NewTrxClient(endPoint string) (*TrxGrpcClient, func(), error) {
	conn := client.NewGrpcClient(endPoint)
	if err := conn.Start(grpc.WithInsecure()); err != nil {
		return nil, func() {}, err
	}

	return &TrxGrpcClient{conn}, func() {
		conn.Stop()
	}, nil
}

func (cli *TrxGrpcClient) SignTx(tx *core.Transaction, key string) (*core.Transaction, error) {
	rawData, err := proto.Marshal(tx.GetRawData())
	if err != nil {
		return nil, err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}
	tx.Signature = append(tx.Signature, signature)
	return tx, nil
}

func (cli *TrxGrpcClient) SendTx(signedTx *core.Transaction) error {
	_, err := cli.Broadcast(signedTx)
	return err
}
