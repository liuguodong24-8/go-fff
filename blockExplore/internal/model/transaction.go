package model

import (
	"blockExplore/internal/service"
	"blockExplore/pkgs"
	paginate "blockExplore/pkgs/mongdb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

// TableTransactionName 表名
var TableTransactionName = "transaction"

// TableTransaction Transaction结构
type TableTransaction struct {
	/*Id         primitive.ObjectID `json:"_id" bson:"_id"`*/
	Hash             string         `json:"hash" bson:"hash"`         //交易hash
	Nonce            pkgs.HexInt    `json:"nonce" bson:"nonce"`       //本次交易之前发送方已经生成的交易数量]
	To               string         `json:"to" bson:"to"`             //交易接收方地址，对于合约创建交易，该值为null
	Value            pkgs.HexBigInt `json:"value" bson:"value"`       //发送的以太数量，单位：wei
	GasPrice         pkgs.HexBigInt `json:"gasPrice" bson:"gasPrice"` //发送方提供的gas价格，单位：wei
	Gas              pkgs.HexInt    `json:"gas" bson:"gas"`           //发送方提供的gas可用量
	BlockHash        string         `json:"blockHash" bson:"blockHash"`
	BlockNumber      pkgs.HexInt    `json:"blockNumber" bson:"blockNumber"`
	Input            string         `json:"input" bson:"input"`
	From             string         `json:"from" bson:"from"`
	TransactionIndex pkgs.HexInt    `json:"transactionIndex" bson:"transactionIndex"`
}

// InsertTransaction 批量新增新增
func InsertTransaction(ctx context.Context, b []interface{}) error {
	if _, err := service.MongoDbEntity.Database.Collection(TableTransactionName).InsertMany(ctx, b); err != nil {
		return err
	}
	return nil
}

// InsertTransactionOne 新增新增一个
func InsertTransactionOne(ctx context.Context, b interface{}) error {
	fmt.Println(b)
	if _, err := service.MongoDbEntity.Database.Collection(TableTransactionName).InsertOne(ctx, b); err != nil {
		return err
	}
	return nil
}

// FindTransaction 查询
func FindTransaction(ctx context.Context, filter interface{}) (TableTransaction, error) {
	var tableTransaction TableTransaction
	if err := service.MongoDbEntity.Database.Collection(TableTransactionName).FindOne(ctx, filter).Decode(&tableTransaction); err != nil {
		return tableTransaction, err
	}
	return tableTransaction, nil
}

// FindAllPageTransaction 查询
func FindAllPageTransaction(filter interface{}, limit, page int) ([]TableTransaction, paginate.PaginationData, error) {
	var tableTransaction []TableTransaction
	collection := service.MongoDbEntity.Database.Collection(TableTransactionName)
	projection := bson.D{
		{"_id", 1},
		{"hash", 1},
		{"nonce", 1},
		{"to", 1},
		{"value", 1},
		{"gasPrice", 1},
		{"gas", 1},
		{"blockHash", 1},
		{"blockNumber", 1},
		{"input", 1},
		{"from", 1},
		{"transactionIndex", 1},
	}
	paginatedData, err := paginate.New(collection).Limit(int64(limit)).Page(int64(page)).Select(projection).Filter(filter).Find()
	if err != nil {
		return nil, paginatedData.Pagination, err
	}
	for _, raw := range paginatedData.Data {
		var transaction *TableTransaction
		if marshallErr := bson.Unmarshal(raw, &transaction); marshallErr == nil {
			tableTransaction = append(tableTransaction, *transaction)
		}
	}
	return tableTransaction, paginatedData.Pagination, nil
}

// FindTransactionCount 查询交易量
func FindTransactionCount(ctx context.Context, filter interface{}) (int64, error) {
	var number int64
	collection := service.MongoDbEntity.Database.Collection(TableTransactionName)
	number, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// UpdateTransaction 更新
func UpdateTransaction(ctx context.Context, filter interface{}, b interface{}) error {
	if _, err := service.MongoDbEntity.Database.Collection(TableTransactionName).UpdateOne(ctx, filter, b); err != nil {

		return err
	}
	return nil
}
