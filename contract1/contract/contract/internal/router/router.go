package router

import (
	"contract/internal/controller/contract"
	"github.com/gin-gonic/gin"
)

// Router 全局 router
var Router *gin.Engine

// Init 初始化执行
func Init() {
	Router = gin.Default()
	groupRoute := Router.Group("/api")
	groupRoute.POST("/mint-nft", contract.MintNft)
	groupRoute.POST("/trans-nft", contract.TransNft)
	groupRoute.POST("/check-address", contract.IsValidHexAddress)
}
