package model

import (
	"blockExplore/internal/service"
	paginate "blockExplore/pkgs/mongdb"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// TableBlockName 表名
var TableBlockName = "block"

// TableBlock block结构
type TableBlock struct {
	/*	Id           primitive.ObjectID `json:"_id" bson:"_id"`*/
	Number          int64              `json:"number" bson:"number"`                   //块编号，挂起块为null
	Hash            common.Hash        `json:"hash" bson:"hash"`                       //块哈希，挂起块为null
	ParentHash      common.Hash        `json:"parentHash" bson:"parentHash"`           //父块的哈希
	Nonce           uint64             `json:"nonce" bson:"nonce"`                     //生成的pow哈希，挂起块为null
	LogsBloom       types.Bloom        `json:"logsBloom" bson:"logsBloom"`             //快日志的bloom过滤器，挂起块为null
	StateRoot       common.Hash        `json:"stateRoot" bson:"stateRoot"`             //块最终状态树的根节点
	Difficulty      int64              `json:"difficulty" bson:"difficulty"`           //难度
	ExtraData       []byte             `json:"extraData" bson:"extraData"`             //块额外数据
	Size            common.StorageSize `json:"size" bson:"size"`                       //本块字节数
	GasLimit        uint64             `json:"gasLimit" bson:"gasLimit"`               //本块允许的最大gas用量
	GasUsed         uint64             `json:"gasUsed" bson:"gasUsed"`                 //本块中所有交易使用的总gas用量
	Timestamp       uint64             `json:"timestamp" bson:"timestamp"`             //块时间戳
	Transactions    types.Transactions `json:"transactions" bson:"transactions"`       //交易对象数组，或32字节长的交易哈希数组 (数量)
	TransactionsNum int                `json:"transactionsNum" bson:"transactionsNum"` //交易数量
	Uncles          []*types.Header    `json:"uncles" bson:"uncles"`                   //叔伯哈希数组
	ReceivedAt      time.Time          `json:"receivedAt" bson:"receivedAt"`           //
	ReceivedFrom    interface{}        `json:"receivedFrom" bson:"receivedFrom"`       //
	NumberU64       uint64             `json:"numberU64" bson:"numberU64"`             //
	MixDigest       common.Hash        `json:"mixDigest" bson:"mixDigest"`             //
	Coinbase        common.Address     `json:"coinbase" bson:"coinbase"`               //
	Root            common.Hash        `json:"root" bson:"root"`                       //
	TxHash          common.Hash        `json:"txHash" bson:"txHash"`                   //
	ReceiptHash     common.Hash        `json:"receiptHash" bson:"receiptHash"`         //
	UncleHash       common.Hash        `json:"uncleHash" bson:"uncleHash"`             //
	Extra           []byte             `json:"extra" bson:"extra"`                     //
	BaseFee         int64              `json:"baseFee" bson:"baseFee"`                 //
	Header          *types.Header      `json:"header" bson:"header"`                   //
	Body            *types.Body        `json:"body" bson:"body"`                       //
	//SanityCheck  error              `json:"sanityCheck" bson:"sanityCheck,omitempty"` //
}

// Insert 批量新增新增
func Insert(ctx context.Context, b []interface{}) error {

	if _, err := service.MongoDbEntity.Database.Collection(TableBlockName).InsertMany(ctx, b); err != nil {
		return err
	}
	return nil
}

// InsertOne 新增一个
func InsertOne(ctx context.Context, b interface{}) error {

	if _, err := service.MongoDbEntity.Database.Collection(TableBlockName).InsertOne(ctx, b); err != nil {
		return err
	}
	return nil
}

// Find 查询
func Find(ctx context.Context, filter interface{}) (TableBlock, error) {
	var tableBlock TableBlock
	if err := service.MongoDbEntity.Database.Collection(TableBlockName).FindOne(ctx, filter).Decode(&tableBlock); err != nil {
		return tableBlock, err
	}
	return tableBlock, nil
}

// FindAllPage 查询
func FindAllPage(filter interface{}, limit, page int) ([]TableBlock, paginate.PaginationData, error) {
	var tableBlock []TableBlock
	collection := service.MongoDbEntity.Database.Collection(TableBlockName)
	projection := bson.D{
		{"_id", 1},
		{"number", 1},
		{"hash", 1},
		{"parentHash", 1},
		{"nonce", 1},
		{"logsBloom", 1},
		{"stateRoot", 1},
		{"difficulty", 1},
		{"extraData", 1},
		{"size", 1},
		{"gasLimit", 1},
		{"gasUsed", 1},
		{"timestamp", 1},
		{"transactions", 1},
		{"transactionsNum", 1},
		{"uncles", 1},
		{"receivedAt", 1},
		{"receivedFrom", 1},
		{"numberU64", 1},
		{"mixDigest", 1},
		{"coinbase", 1},
		{"root", 1},
		{"txHash", 1},
		{"receiptHash", 1},
		{"uncleHash", 1},
		{"extra", 1},
		{"baseFee", 1},
		{"header", 1},
		{"body", 1},
		//{"sanityCheck", 1},
	}

	paginatedData, err := paginate.New(collection).Limit(int64(limit)).Page(int64(page)).Sort("timestamp", -1).Select(projection).Filter(filter).Find()
	if err != nil {
		return nil, paginatedData.Pagination, err
	}
	for _, raw := range paginatedData.Data {
		var block *TableBlock
		marshallErr := bson.Unmarshal(raw, &block)
		if marshallErr == nil {
			tableBlock = append(tableBlock, *block)
		}
	}

	return tableBlock, paginatedData.Pagination, nil
}

// Update 更新
func Update(ctx context.Context, filter interface{}, b interface{}) error {
	if _, err := service.MongoDbEntity.Database.Collection(TableBlockName).UpdateOne(ctx, filter, b); err != nil {

		return err
	}
	return nil
}
