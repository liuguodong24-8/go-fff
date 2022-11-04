package router

import (
	"blockExplore/internal/controller/block"
	"blockExplore/internal/controller/index"
	"blockExplore/internal/controller/transaction"
	"github.com/gin-gonic/gin"
)

// Router 全局 router
var Router *gin.Engine

// Init 初始化执行
func Init() {
	Router = gin.Default()
	Router.GET(`/api/test`, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	Router.GET(`/api/index`, index.GetIndex)
	Router.GET(`/api/blocks`, block.GetBlockList)
	Router.GET(`/api/blocks/:hash`, block.GetBlockDetail)
	Router.GET(`/api/transactions`, transction.GetTransactionList)
	Router.GET(`/api/erc20-token/transactions`, transction.GetTransactionListErc20)
	Router.GET(`/api/transactions/:hash`, transction.GetTransactionLDetail)
}
