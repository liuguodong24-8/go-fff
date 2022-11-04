package mongdb

import (
	"blockExplore/internal/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDb mongoDb
type MongoDb struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// Config redis config
type Config struct {
	MongodbUri string
	Timeout    int
	Database   string
}

// NewEntity 实例化 mongodb
func NewEntity(ctx context.Context, cfg Config) (*MongoDb, error) {
	var (
		opts     *options.ClientOptions
		client   *mongo.Client
		err      error
		database *mongo.Database
	)

	// 连接数据库
	opts = options.Client().ApplyURI(config.Setting.MongoDb.MongodbUri) // opts
	if client, err = mongo.Connect(ctx, opts); err != nil {
		fmt.Println(err)
		return &MongoDb{}, err
	}

	//链接数据库和表
	database = client.Database(config.Setting.MongoDb.Database)

	return &MongoDb{
		Client:   client,
		Database: database,
	}, nil
}
