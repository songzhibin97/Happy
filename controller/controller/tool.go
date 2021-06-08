package controller

import (
	"Happy/model/model"
	"Happy/settings"
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// 存放一些共用的工具
// InitTrans 初始化翻译器

var trans ut.Translator

// InitTrans:gin初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// OtherInitTrans:grpc初始化翻译器
func OtherInitTrans(v *validator.Validate, local string) error {
	// 注册一个获取json tag的自定义方法
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器

	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(enT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'
	var ok bool
	var err error
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	trans, ok = uni.GetTranslator(local)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", local)
	}
	// 注册翻译器
	switch local {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return err
}

// GetOtherValidator:封装OtherInitTrans返回Validator对象
func GetOtherValidator() *validator.Validate {
	validate := validator.New()
	_ = OtherInitTrans(validate, settings.GetString("App.Language"))
	return validate
}

// IsVerifyError:判断是否为校验失败
func IsVerifyError(err error) (string, bool) {
	// 修改 原来是返回 validator.ValidationErrorsTranslations类型(map[string]string)
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return "", false
	}
	return RemoveTopStruct(errs.Translate(trans)), true

}

// RemoveTopStruct:清除返回错误map对应的结构体前缀
// 修改 原来是返回map[string]string
func RemoveTopStruct(fields map[string]string) string {
	//// 新增处理:只返回一个错误
	//// 只取第一个key:value键值对
	flag := 0
	//res := make(map[string]string, len(fields))
	res := ""
	//for field, err := range fields {
	for _, err := range fields {
		if flag == 1 {
			return res
		}
		// 特殊处理 最后一次.出现的位置进行剪切
		//res[field[strings.LastIndex(field, ".")+1:]] = err
		// 特殊处理 只返回第一个错误信息 不是map对
		res = err
		flag++
	}
	return res
}

// ParameterVerification:校验gin请求参数
func ParameterVerification(c *gin.Context, i interface{}) error {
	if err := c.ShouldBind(i); err != nil {
		// 校验失败
		// 判断error是否是校验失败的error
		errs, ok := IsVerifyError(err)
		if !ok {
			// 如果不是校验失败的错误就返回异常 标记服务器异常
			zap.L().Error("IsVerifyError", zap.Error(err))
			model.ResponseErrorWithMsg(c, model.CodeServerBusy, err.Error())
			return model.CodeServerBusy.Err()
		}
		// 是参数校验的错误返回对应错误
		zap.L().Info("VerifyError", zap.Any("error", errs))
		model.ResponseErrorWithMsg(c, model.CodeInvalidParams, errs)
		return model.CodeInvalidParams.Err()
	}
	return nil
}

// StoA:string转int
// 如果error不为空将c中封装返回值
func StoA(s string, c *gin.Context) (int, error) {
	res, err := strconv.Atoi(s)
	if err != nil {
		model.ResponseErrorWithMsg(c, model.CodeInvalidParams, s)
	}
	return res, err
}

// GetToken:获取token并绑定到请求头上
func GetToken(c *gin.Context) (context.Context, error) {
	token, ok := c.Get("accessJWT")
	if !ok {
		zap.L().Error("Get(middleware.AccessToken)")
		model.ResponseErrorWithMsg(c, model.CodeServerBusy, "获取上下文Token失败")
		return nil, errors.New("获取上下文Token失败")
	}
	md := metadata.Pairs("authorization", fmt.Sprintf("bearer %s", token))
	ctx := metadata.NewOutgoingContext(c, md)
	return ctx, nil
}