/******
** @创建时间 : 2020/8/16 10:34
** @作者 : SongZhiBin
******/
package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 定义响应内容
// ResponseStruct:返回响应结构体
/*
	Response:{
		code : 200  // 返回响应状态码
		msg  : Some errors or hints  // 携带一些错误或提示信息
		data : {	// 携带信息
			"key":"value"
		}
	}
*/
type ResponseStruct struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseErrorWithMsg:自定义error状态码以及错误信息
func ResponseErrorWithMsg(c *gin.Context, code ResCode, errMsg interface{}) {
	c.JSON(http.StatusOK, ResponseStruct{
		Code: code,
		Msg:  errMsg,
		Data: nil,
	})
}

// ResponseError:返回已知错误类型
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, ResponseStruct{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseSuccess:返回成功内容
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseStruct{
		Code: CodeOk,
		Msg:  CodeOk.Msg(),
		Data: data,
	})
}
