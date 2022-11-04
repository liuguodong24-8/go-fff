package api

import (
	"blockExplore/pkgs"
	paginate "blockExplore/pkgs/mongdb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pagination struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	TotalPage int64 `json:"totalPage"`
}

func Created(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"code": pkgs.Success,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"code": pkgs.Success,
	})
}

func ServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": pkgs.ErrInternal,
		"msg":  message,
	})
}

func Unprocessable(c *gin.Context, message string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"code": pkgs.ErrUnprocessableEntity,
		"msg":  message,
	})
}

func MakePage(c *gin.Context, data interface{}, page paginate.PaginationData) {
	c.JSON(http.StatusOK, gin.H{
		"code": pkgs.Success,
		"data": data,
		"page": Pagination{
			Total:     page.Total,     //总数
			Page:      page.Page,      //当前页数
			PerPage:   page.PerPage,   //每页数量
			TotalPage: page.TotalPage, //总页数
		},
	})
}

func Make(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": pkgs.Success,
		"data": data,
	})
}
