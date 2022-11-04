package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/basic/admin"
	"github.com/swaggo/gin-swagger/example/basic/api"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// Router 全局 router
var Router *gin.Engine

// Init 初始化执行
// @title TestSwg API
// @version 1.0
// @host localhost:5003
// @BasePath /
func Init() {
	r := gin.New()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	groupRoute := r.Group("/admin/api")

	groupRoute.GET("/get-string-by-int/:some_id", api.GetStringByInt)
	groupRoute.GET("/get-struct-array-by-string/:some_id", api.GetStructArrayByString)

	groupRoute.GET("/get-string", admin.GetSting)
	groupRoute.GET("/get-int", admin.GetInt)

	r.Run()
}
