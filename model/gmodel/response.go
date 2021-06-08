package gmodel

import (
	"Happy/model/model"
	pb "Happy/model/pmodel/user"
	"encoding/json"

	"go.uber.org/zap"
)

// grpc响应封装

// ResponseWithMsg:自定义error状态码以及错误信息
func ResponseWithMsg(code model.ResCode, errMsg interface{}) *pb.Response {
	msg, _ := json.Marshal(errMsg)
	return &pb.Response{
		Code: int32(code),
		Msg:  string(msg),
	}
}

// ResponseError:返回已知错误类型
func ResponseError(code model.ResCode) *pb.Response {
	return &pb.Response{
		Code: int32(code),
		Msg:  code.Msg(),
	}
}

// ResponseSuccess:返回成功内容
func ResponseSuccess(data map[string]string) *pb.Response {
	return &pb.Response{
		Code: int32(model.CodeOk),
		Msg:  model.CodeOk.Msg(),
		Data: data,
	}
}

// GinResponse:转化为gin的响应
func GinResponse(response *pb.Response) *model.ResponseStruct {
	zap.L().Info("Response", zap.Any("response", response))
	return &model.ResponseStruct{
		Code: model.ResCode(response.Code),
		Msg:  response.Msg,
		Data: response.Data,
	}
}
