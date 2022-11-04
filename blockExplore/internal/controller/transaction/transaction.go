package transction

import (
	"blockExplore/internal/model"
	"blockExplore/internal/util"
	"blockExplore/pkgs/api"
	paginate "blockExplore/pkgs/mongdb"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetTransactionList 区块列表列表
func GetTransactionList(c *gin.Context) {
	var request struct {
		PageSize int    `json:"page_size" form:"page_size" binding:"required"`
		Page     int    `json:"page" form:"page" binding:"required"`
		To       string `json:"to" form:"to"`
	}
	if err := c.BindQuery(&request); err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("获取交易列表参数错误:%v", err))
		api.Unprocessable(c, fmt.Sprintf("获取交易列表参数错误:%v", err))
		return
	}

	var transaction []model.TableTransaction
	var page paginate.PaginationData
	to := bson.M{}
	if request.To != "" {
		//to = bson.D{{"to", request.To}, {"from", request.To}}
		//或者查询
		to = bson.M{"$or": []bson.M{bson.M{"to": request.To}, bson.M{"from": request.To}}}
	}

	transaction, page, err := model.FindAllPageTransaction(to, request.PageSize, request.Page)
	if err != nil {
		if transaction == nil {
			api.MakePage(c, nil, page)
			return
		}
		util.Logger.Error(fmt.Sprintf("获取交易列表失败:%v", err))
		api.ServerError(c, fmt.Sprintf("获取交易列表失败:%v", err))
		return
	}
	api.MakePage(c, transaction, page)
}

// GetTransactionListErc20 Erc20区块交易列表列表
func GetTransactionListErc20(c *gin.Context) {
	var request struct {
		PageSize int    `json:"page_size" form:"page_size" binding:"required"`
		Page     int    `json:"page" form:"page" binding:"required"`
		To       string `json:"to" form:"to"`
	}
	if err := c.BindQuery(&request); err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("获取交易列表参数错误:%v", err))
		api.Unprocessable(c, fmt.Sprintf("获取交易列表参数错误:%v", err))
		return
	}

	var transaction []model.TableTransaction
	var page paginate.PaginationData
	to := bson.M{}
	if request.To != "" {
		//to = bson.D{{"to", request.To}, {"from", request.To}}
		//或者查询
		to = bson.M{"$or": []bson.M{bson.M{"to": request.To}, bson.M{"from": request.To}}}
	}

	transaction, page, err := model.FindAllPageTransaction(to, request.PageSize, request.Page)
	if err != nil {
		if transaction == nil {
			api.MakePage(c, nil, page)
			return
		}
		util.Logger.Error(fmt.Sprintf("获取交易列表失败:%v", err))
		api.ServerError(c, fmt.Sprintf("获取交易列表失败:%v", err))
		return
	}
	api.MakePage(c, transaction, page)
}

// GetTransactionLDetail 交易详情
func GetTransactionLDetail(c *gin.Context) {
	hash := c.Param("hash")
	if hash == "" {
		api.Unprocessable(c, fmt.Sprintf("hash不能为空"))
	}
	var transaction model.TableTransaction
	transaction, err := model.FindTransaction(c, bson.M{"hash": hash})
	if err != nil {
		util.Logger.Error(fmt.Sprintf("获取交易列表详情失败:%v", err))
		api.ServerError(c, fmt.Sprintf("获取交易详情失败:%v", err))
		return
	}
	api.Make(c, transaction)
}
