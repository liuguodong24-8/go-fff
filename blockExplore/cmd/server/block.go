package main

import (
	"blockExplore/internal/config"
	"blockExplore/internal/model"
	"blockExplore/internal/service"
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/websocket"
	"log"
	"math/big"
	"reflect"
	"runtime"
	"time"
)

const (
	TIMER       = 10
	BLOCKLENGTH = 10
)

func init() {
	var conf string
	flag.StringVar(&conf, "c", "", "指定配置文件位置")
	//flag.StringVar(&status, "f", "", "指定配置单个还是批量")
	flag.Parse()
	if conf == "" {
		panic("请指定配置文件位置")
	}
	config.Load(conf)

}

func main() {
	log.Println("开始同步区块数据")
	ctx := context.Background() // ctx
	err := service.SetMongoDb()
	if err != nil {
		log.Fatal("链接mongodb失败")
	}

	// 获取链上块
	client, err := ethclient.Dial(config.Setting.Chain.Address)
	if err != nil {
		log.Fatal(err)
	}

	// 获取链上块
	clientRpc, err := rpc.Dial(config.Setting.Chain.Address)
	if err != nil {
		log.Fatal(err)
	}
	//查询上次插入的区块高度
	start, end := 0, 0
	blocks, _, err := model.FindAllPage(bson.M{}, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	if blocks != nil {
		start = int(blocks[0].Number)
		end = start
	}

	t := time.NewTimer(TIMER * time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			go func() {
				//nowStats()
				//获取块高度
				header, err := client.HeaderByNumber(ctx, nil)
				if err != nil {
					log.Fatal(err)
				}
				//最新区块信息
				if (end + BLOCKLENGTH) < int(header.Number.Int64()) {
					end += BLOCKLENGTH
				} else {
					end = int(header.Number.Int64())
				}

				log.Printf("块总高:%d", header.Number.Int64())
				log.Printf("执行开始块高:%d", start)
				log.Printf("当前块高度:%d", end)

				err = batchTrans(ctx, *client, *clientRpc, start, end)
				if err != nil {
					log.Fatal(err)
				}

				log.Printf("同步区块数据完成%d--%d", start, end)

			}()

			start = end + 1
			runtime.GC()
			t.Reset(TIMER * time.Second)
			//continue
		}
	}
	log.Println("同步区块数据完成")
}

// batchTrans 批量写入交易信息
func batchTrans(ctx context.Context, client ethclient.Client, clientRpc rpc.Client, start, end int) error {
	//考虑到多条
	var transactions []interface{}
	var blocks []interface{}
	for i := start; i <= end; i++ {
		block, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			log.Print("number")
			log.Fatal(err)
		}

		var fee, number, difficulty int64
		if block.BaseFee() != nil {
			fee = block.BaseFee().Int64()
		}
		if block.Number() != nil {
			number = block.Number().Int64()
		}
		if block.Difficulty() != nil {
			difficulty = block.Difficulty().Int64()
		}

		blocks = append(blocks, model.TableBlock{
			Number:          number,
			Hash:            block.Hash(),
			ParentHash:      block.ParentHash(),
			Nonce:           block.Nonce(),
			LogsBloom:       block.Bloom(),
			StateRoot:       block.Root(),
			Difficulty:      difficulty,
			ExtraData:       block.Extra(),
			Size:            block.Size(),
			GasLimit:        block.GasLimit(),
			GasUsed:         block.GasUsed(),
			Timestamp:       block.Time(),
			Transactions:    block.Transactions(),
			TransactionsNum: block.Transactions().Len(),
			Uncles:          block.Uncles(),
			ReceivedAt:      block.ReceivedAt,
			ReceivedFrom:    block.ReceivedFrom,
			NumberU64:       block.NumberU64(),
			MixDigest:       block.MixDigest(),
			Coinbase:        block.Coinbase(),
			Root:            block.Root(),
			TxHash:          block.TxHash(),
			ReceiptHash:     block.ReceiptHash(),
			UncleHash:       block.UncleHash(),
			Extra:           block.Extra(),
			BaseFee:         fee,
			Header:          block.Header(),
			Body:            block.Body(),
		})

		//transaction, isPending, _ := client.TransactionByHash(context.Background(), common.HexToHash("0x079a529094efa177cb6f67acea8786040d4bf45bdeba758b850a728afd84b03e"))
		if block.Transactions().Len() > 0 {
			for _, v := range block.Transactions() {
				var transaction model.TableTransaction
				err := clientRpc.Call(&transaction, "eth_getTransactionByHash", v.Hash())
				if err != nil {
					log.Fatal(err)
				}
				if !reflect.DeepEqual(transaction, model.TableTransaction{}) {
					log.Printf("此块有交易数据， hash：%s", v.Hash())
					transactions = append(transactions, transaction)
				}
			}
		}
	}

	//fmt.Println(blocks)
	if err := model.Insert(ctx, blocks); err != nil {
		log.Print("insert block error")
		log.Fatal(err)
	}
	if transactions != nil {
		if err := model.InsertTransaction(ctx, transactions); err != nil {
			log.Print("insert transaction error")
			log.Fatal(err)
		}
	}

	log.Printf("区间数据保存成功， start：%d end：%d", start, end)
	return nil
}

// insertOne 插入单条数据
func insertOne(ctx context.Context, client ethclient.Client) error {
	ws, err := websocket.Dial(config.Setting.Chain.WsAddress, "", config.Setting.Chain.Address)
	if err != nil {
		return err
	}
	var dataBlock model.TableBlock
	err = websocket.Message.Receive(ws, dataBlock)
	if err != nil {
		return err
	}

	block, err := model.Find(ctx, bson.M{"hash": dataBlock.Hash})
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(block, model.TableBlock{}) {
		//service.MongoDbEntity.Database.Collection(model.TableBlockName).DeleteOne(ctx, bson.M{"hash": block.Hash})
		model.Update(ctx, bson.M{"hash": block.Hash}, bson.M{"$set": block})
		if err := transInsert(ctx, client, block); err != nil {
			return err
		}

	} else {
		if err = model.InsertOne(ctx, model.TableBlock{
			Number:       dataBlock.Number,
			Hash:         dataBlock.Hash,
			ParentHash:   dataBlock.ParentHash,
			Nonce:        dataBlock.Nonce,
			LogsBloom:    dataBlock.LogsBloom,
			StateRoot:    dataBlock.Root,
			Difficulty:   dataBlock.Difficulty,
			ExtraData:    dataBlock.Extra,
			Size:         dataBlock.Size,
			GasLimit:     dataBlock.GasLimit,
			GasUsed:      dataBlock.GasUsed,
			Timestamp:    dataBlock.Timestamp,
			Transactions: dataBlock.Transactions,
			Uncles:       dataBlock.Uncles,
			ReceivedAt:   dataBlock.ReceivedAt,
			ReceivedFrom: dataBlock.ReceivedFrom,
			NumberU64:    dataBlock.NumberU64,
			MixDigest:    dataBlock.MixDigest,
			Coinbase:     dataBlock.Coinbase,
			Root:         dataBlock.Root,
			TxHash:       dataBlock.TxHash,
			ReceiptHash:  dataBlock.ReceiptHash,
			UncleHash:    dataBlock.UncleHash,
			Extra:        dataBlock.Extra,
			BaseFee:      dataBlock.BaseFee,
			Header:       dataBlock.Header,
			Body:         dataBlock.Body,
			//SanityCheck:  dataBlock.SanityCheck,
		}); err != nil {
			log.Fatal(err)
		}
		if err := transInsert(ctx, client, dataBlock); err != nil {
			return err
		}
	}
	return nil
}

// transInsert 交易信息添加
func transInsert(ctx context.Context, client ethclient.Client, dataBlock model.TableBlock) error {
	//考虑到多条
	var transactions []interface{}
	//transaction, isPending, _ := client.TransactionByHash(context.Background(), common.HexToHash("0x079a529094efa177cb6f67acea8786040d4bf45bdeba758b850a728afd84b03e"))
	if dataBlock.Transactions.Len() > 0 {
		for _, v := range dataBlock.Transactions {
			transaction, isPending, _ := client.TransactionByHash(context.Background(), v.Hash())
			if !isPending && transaction != nil {
				/*transactions = append(transactions, model.TableTransaction{
					Hash:       transaction.Hash(),
					Nonce:      hexutil.Uint64(transaction.Nonce()),
					To:         transaction.To(),
					Value:      transaction.Value().String(),
					GasPrice:   transaction.GasPrice().String(),
					Gas:        hexutil.Uint64(transaction.Gas()),
					AccessList: transaction.AccessList(),
					ChainId:    transaction.ChainId().String(),
					Cost:       transaction.Cost().String(),
					Data:       transaction.Data(),
					GasFeeCap:  transaction.GasFeeCap().String(),
					GasTipCap:  transaction.GasTipCap().String(),
					Protected:  transaction.Protected(),
				})*/
			}
		}
	}
	if err := model.InsertTransaction(ctx, transactions); err != nil {
		return err
	}
	return nil
	//只有一条数据时
	//if !isPending && transaction != nil {
	/*transDb, err := model.FindTransaction(ctx, bson.M{"hash": transactionAttr.Hash})
	if err != nil {
		return err
	}*/

	//更新
	/*if !reflect.DeepEqual(transDb, model.TableTransaction{}) {
		if err = model.UpdateTransaction(ctx, bson.M{"hash": transaction.Hash}, bson.M{"$set": transactionAttr}); err != nil {
			return err
		}
		//新增
	} else {*/
	/*}*/
	//}
}

func nowStats() string {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return fmt.Sprintf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes) NumGoroutine:%d", ms.Alloc, ms.HeapIdle, ms.HeapReleased, runtime.NumGoroutine())
}
