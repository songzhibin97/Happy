package gcontroller

import (
	"Happy/controller/controller"
	"Happy/model/gmodel"
	"Happy/model/model"
	pbUser "Happy/model/pmodel/user"

	"go.uber.org/zap"
)

// 通用函数

// verification
func _verification(request interface{}) (*pbUser.Response, error) {
	// 校验请求参数
	validate := controller.GetOtherValidator()
	err := validate.Struct(request)
	if err != nil {
		// 判断错误是否是校验失败
		errs, ok := controller.IsVerifyError(err)
		if !ok {
			// 如果不是校验失败的错误就返回异常
			zap.L().Error("IsVerifyError", zap.Error(err))
			return gmodel.ResponseWithMsg(model.CodeServerBusy, err), err
		}
		zap.L().Info("VerifyError", zap.Any("error", errs))
		return gmodel.ResponseWithMsg(model.CodeServerBusy, errs), err
	}
	return nil, nil
}
