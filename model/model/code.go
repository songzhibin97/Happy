package model

import "errors"

// 规范了返回的参数代码以及定义了一些常量

// ResCode:定义code类型
type ResCode int

const (
	CodeOk ResCode = 1000 + iota
	Code404
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeInvalidAuthFormat
	CodeNotLogin
	CodeUnKnowError

	CodeJWTVerificationFailed
	CodeJWTExpired
	CodeAccessExpired
	CodeMultiTerminalLogin

	CodeFrequentRequests
	CodeInvalidVerificationCode

	CodeGetListEmpty
	CodeGetListError

	CodePostError

	CodeRepeatVoting
	CodeAuthNoVote
)

// ResMsg:对应Code返回的Msg
var ResMsg = map[ResCode]string{
	CodeOk:                      "Success",
	Code404:                     "404",
	CodeInvalidParams:           "请求参数错误",
	CodeUserExist:               "用户名重复",
	CodeUserNotExist:            "用户不存在",
	CodeInvalidPassword:         "用户名或密码错误",
	CodeServerBusy:              "服务繁忙",
	CodeInvalidToken:            "无效的Token",
	CodeInvalidAuthFormat:       "认证格式有误",
	CodeNotLogin:                "未登录",
	CodeUnKnowError:             "未知错误",
	CodeJWTVerificationFailed:   "jwt验证失败",
	CodeJWTExpired:              "jwt已过期",
	CodeAccessExpired:           "AccessToken已过期",
	CodeMultiTerminalLogin:      "账户在其他设备上登录",
	CodeFrequentRequests:        "请求频发,请稍后再试",
	CodeInvalidVerificationCode: "验证码无效",
	CodeGetListEmpty:            "获取列表为空",
	CodeGetListError:            "获取列表失败",

	CodePostError: "发帖失败",

	CodeRepeatVoting: "重复投票",
	CodeAuthNoVote:   "作者不能投票",
}

// Msg:从Map获取Msg
func (r ResCode) Msg() string {
	msg, ok := ResMsg[r]
	if ok {
		return msg
	}
	return ResMsg[CodeServerBusy]
}

// Err:返回对应的error信息
func (r ResCode) Err() error {
	errs, ok := ResMsg[r]
	if !ok {
		return errors.New(ResMsg[CodeServerBusy])
	}
	return errors.New(errs)
}
