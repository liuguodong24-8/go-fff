package api

import (
	"contract/pkgs"
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

func Make(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": pkgs.Success,
		"data": data,
	})
}

func Success(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": pkgs.Success,
		"mag":  message,
	})
}
