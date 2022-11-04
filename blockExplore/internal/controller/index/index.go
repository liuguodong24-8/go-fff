package index

import (
	"blockExplore/internal/model"
	"blockExplore/internal/service"
	"blockExplore/internal/util"
	"blockExplore/pkgs"
	"blockExplore/pkgs/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetIndex 区块列表主页详情
func GetIndex(c *gin.Context) {

	type Index struct {
		Number int64
		Price  pkgs.HexBigInt
	}
	var index Index
	number, err := model.FindTransactionCount(c, bson.M{})
	if err != nil {
		util.Logger.Error(fmt.Sprintf("获取交易数量失败:%v", err))
		api.ServerError(c, fmt.Sprintf("获取交易数量失败:%v", err))
		return
	}
	index.Number = number

	var price pkgs.HexBigInt
	err = service.EthClientEntity.Call(&price, "eth_gasPrice")
	if err != nil {
		util.Logger.Error(fmt.Sprintf("获取gas价格失败:%v", err))
		api.ServerError(c, fmt.Sprintf("获取gas价格失败:%v", err))
		return
	}
	index.Price = price
	api.Make(c, index)
}
