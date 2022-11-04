package service

import (
	"blockExplore/internal/config"
	"blockExplore/pkgs/mongdb"
	"context"
	"time"
)

var MongoDbEntity mongdb.MongoDb

// SetMongoDb 设置mongoDb
func SetMongoDb() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Setting.MongoDb.Timeout)*time.Millisecond) // ctx
	mGoEntity, err := mongdb.NewEntity(ctx, mongdb.Config{
		MongodbUri: config.Setting.MongoDb.MongodbUri,
		Timeout:    config.Setting.MongoDb.Timeout,
		Database:   config.Setting.MongoDb.Database,
	})

	if err != nil {
		return err
	}
	MongoDbEntity = *mGoEntity
	return nil
}
