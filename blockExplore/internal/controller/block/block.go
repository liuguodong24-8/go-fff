package block

import (
	"blockExplore/internal/model"
	"blockExplore/internal/util"
	"blockExplore/pkgs/api"
	paginate "blockExplore/pkgs/mongdb"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetBlockList 区块列表列表
func GetBlockList(c *gin.Context) {
	var request struct {
		PageSize int `json:"page_size" form:"page_size" binding:"required"`
		Page     int `json:"page" form:"page" binding:"required"`
	}
	if err := c.BindQuery(&request); err != nil {
		util.Logger.WithError(err).Error(fmt.Sprintf("获取区块链列表参数错误:%v", err))
		api.Unprocessable(c, fmt.Sprintf("获取区块链列表参数错误:%v", err))
		return
	}

	var blocks []model.TableBlock
	var page paginate.PaginationData
	blocks, page, err := model.FindAllPage(bson.M{}, request.PageSize, request.Page)
	if err != nil {
		if blocks == nil {
			api.MakePage(c, nil, page)
			return
		}
		util.Logger.Error(fmt.Sprintf("获取区块链列表失败:%v", err))
		api.ServerError(c, fmt.Sprintf("获取区块链列表失败:%v", err))
		return
	}
	api.MakePage(c, blocks, page)
}

// GetBlockDetail 区块列表详情
func GetBlockDetail(c *gin.Context) {
	hash := c.Param("hash")
	var block model.TableBlock
	block, err := model.Find(c, bson.M{"hash": common.HexToHash(hash)})
	if err != nil {
		util.Logger.Error(fmt.Sprintf("获取区块链详情失败:%v", err))
		api.ServerError(c, fmt.Sprintf("获取区块链详情失败:%v", err))
		return
	}
	api.Make(c, block)
}
