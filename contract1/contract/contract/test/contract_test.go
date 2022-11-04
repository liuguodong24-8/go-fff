package test

import (
	"contract/internal/config"
	"contract/internal/controller/contract"
	"contract/internal/util"
	"contract/pkgs/logger"
	"fmt"
	utils "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"testing"
)

func init() {
	config.Load("../config/config.toml")

	//注册日志
	if err := util.RegisterLogger(logger.Config{
		Channel:    config.Setting.Log.Channel,
		Level:      logger.Level(config.Setting.Log.Level),
		OutputFile: config.Setting.Log.Output,
		WithStack:  config.Setting.Log.Stack,
	}); nil != err {
		panic(fmt.Sprintf("注册日志失败:%s", err.Error()))
	}
	// 捕获异常信息
	defer util.CatchException()

	router := gin.Default() // 这需要写到init中，启动gin框架

	groupRoute := router.Group("/api")
	groupRoute.POST("/mint-nft", contract.MintNft)
	groupRoute.POST("/check-address", contract.IsValidHexAddress)
	utils.SetRouter(router) //把启动的engine 对象传入到test框架中
}

// 解析返回的错误类型
type OrdinaryResponse struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

// 真正的测试单元
func TestMintNftHandler(t *testing.T) {
	// 定义发送POST请求传的内容
	user := map[string]interface{}{
		"to_address":    "0x1D6d8c0ec7FF4E03d0Be661e5007B177daB9F69d",
		"nft_name":      "测试",
		"nft_content":   "NFT介绍",
		"nft_image_url": "https://dgss0.bdstatic.com/5bVWsj_p_tVS5dKfpU_Y_D3/res/r/image/2017-09-27/297f5edb1e984613083a2d3cc0c5bb36.png",
		"nft_cert_list": []string{"https://dgss0.bdstatic.com/5bVWsj_p_tVS5dKfpU_Y_D3/res/r/image/2017-09-27/297f5edb1e984613083a2d3cc0c5bb36.png"},
		"type":          "eth",
	}
	// 把返回response解析到resp中
	resp := OrdinaryResponse{}
	// 调用函数发起http请求
	err := utils.TestHandlerUnMarshalResp("POST", "/api/mint-nft", "json", user, &resp)
	if err != nil {
		t.Errorf("TestMintNftHandler: %v\n", err)
		return
	}
	// 得到返回数据结构体， 至此，完美完成一次post请求测试，
	// 如果需要benchmark 输出性能报告也是可以的
	fmt.Println("result:", resp)

}

// 真正的测试单元
func TestIsValidHexAddressHandler(t *testing.T) {
	// 定义发送POST请求传的内容
	user := map[string]interface{}{
		"address": "0x1D6d8c0ec7FF4E03d0Be661e5007B177daB9F69d",
	}
	// 把返回response解析到resp中
	resp := OrdinaryResponse{}
	// 调用函数发起http请求
	err := utils.TestHandlerUnMarshalResp("POST", "/api/check-address", "json", user, &resp)
	if err != nil {
		t.Errorf("IsValidHexAddress: %v\n", err)
		return
	}
	// 得到返回数据结构体， 至此，完美完成一次post请求测试，
	// 如果需要benchmark 输出性能报告也是可以的
	if err != nil {
		t.Errorf("IsValidHexAddress: %v\n", err)
		return
	}
	fmt.Println(resp)

}

// 真正的测试单元
func BenchmarkIsValidHexAddressHandler(b *testing.B) {
	// 定义发送POST请求传的内容
	user := map[string]interface{}{
		"address": "0x1D6d8c0ec7FF4E03d0Be661e5007B177daB9F69d",
	}
	// 把返回response解析到resp中
	resp := OrdinaryResponse{}
	// 调用函数发起http请求
	err := utils.TestHandlerUnMarshalResp("POST", "/api/check-address", "json", user, &resp)
	if err != nil {
		b.Errorf("IsValidHexAddress: %v\n", err)
		return
	}
	// 得到返回数据结构体， 至此，完美完成一次post请求测试，
	// 如果需要benchmark 输出性能报告也是可以的
	if err != nil {
		b.Errorf("IsValidHexAddress: %v\n", err)
		return
	}
	fmt.Println(resp)

}
