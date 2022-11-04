package util

import (
	"context"
	"contract/pkgs"
	"contract/pkgs/logger"
	"runtime/debug"
)

// Logger 日志entity
var Logger Entity

// Entity ...
type Entity struct {
	*logger.Entity
}

// RegisterLogger 注册日志模块
func RegisterLogger(cfg logger.Config) error {
	entity, err := logger.NewLoggerEntity(cfg)
	if err != nil {
		return err
	}

	Logger = Entity{entity}

	return nil
}

// WithSleuthContext 日志记录ctx
func (e *Entity) WithSleuthContext(ctx context.Context) *Entity {
	entity := e.WithFields("context", logger.MakeFields(pkgs.GetSleuthCtx(ctx)))

	return &Entity{entity}
}

// CatchException 捕获异常panic情况
func CatchException() {
	errs := recover()
	if nil == errs {
		return
	}

	Logger.WithFields("exception", logger.Fields{
		"exception stack": string(debug.Stack()),
	}).Error("捕获异常信息")
}
