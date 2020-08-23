/******
** @创建时间 : 2020/8/11 20:49
** @作者 : SongZhiBin
******/
package router

import (
	"Happy/controller"
	"Happy/logger"
	"Happy/middleware"
	"Happy/model"
	"Happy/pkg/websocket"
	"Happy/settings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	NoAuthenticationRequire = "NoAuthenticationRequire"
	AuthenticationRequire   = "AuthenticationRequire"
)

type Option func(group *gin.RouterGroup)

type Options []Option

type OptionsWare map[string]Options

var OptionsWares OptionsWare

// OptionsWare封装一些方法
// AddNoAuthenticationRequire:添加不用认证的路由
func (o *OptionsWare) AddNoAuthenticationRequire(options ...Option) {
	// 1.判断 NoAuthenticationRequire 是否初始化
	_, ok := (*o)[NoAuthenticationRequire]
	if !ok {
		(*o)[NoAuthenticationRequire] = make(Options, 0)
	}
	(*o)[NoAuthenticationRequire] = append((*o)[NoAuthenticationRequire], options...)
}

// AddAuthenticationRequire:添加需要认证的路由

func (o *OptionsWare) AddAuthenticationRequire(options ...Option) {
	// 1.判断 NoAuthenticationRequire 是否初始化
	_, ok := (*o)[AuthenticationRequire]
	if !ok {
		(*o)[AuthenticationRequire] = make(Options, 0)
	}
	(*o)[AuthenticationRequire] = append((*o)[AuthenticationRequire], options...)
}

// 注册路由
func SetUp() *gin.Engine {
	if settings.GetString("APP.Mode") == "release" {
		// 发布版本
		gin.SetMode(gin.ReleaseMode)
	}
	// 注册validator错误翻译器
	err := controller.InitTrans(settings.GetString("APP.Language"))
	if err != nil {
		zap.L().Error("controller.InitTrans Error", zap.Error(err))
	}
	r := gin.New()
	// 注册中间件
	r.Use(logger.GinLogger(zap.L()), logger.GinRecovery(zap.L(), true))
	// todo 注册自己的业务路由
	InitOptionsWare()
	// 存放不需要验证身份的路由
	nar := r.Group("/")
	{
		// 轮询开始添加
		for _, v := range OptionsWares[NoAuthenticationRequire] {
			v(nar)
		}
	}
	// 存放需要验证身份的路由
	ar := r.Group("/")
	ar.Use(middleware.VerificationJWT)
	if settings.GetString("JWT.Mode") == "refresh" {
		ar.POST("refresh", middleware.VerificationRefreshJWT)
	}
	{
		// 轮询开始添加
		for _, v := range OptionsWares[AuthenticationRequire] {
			v(ar)
		}
	}
	r.NoRoute(func(c *gin.Context) {
		model.ResponseSuccess(c, gin.H{})
	})
	return r
}

// 初始化optionWare
func InitOptionsWare() {
	OptionsWares = make(map[string]Options)
	OptionsWares[NoAuthenticationRequire] = make(Options, 0)
	OptionsWares[AuthenticationRequire] = make(Options, 0)
	// 加入登录注册
	internalAdd()
	OtherApp()
}

// 用户相关路由

// internalAdd:加入注册登录
func internalAdd() {
	OptionsWares.AddNoAuthenticationRequire(Ws, SignUp, Login, GetVerificationCode)
}

// OtherApp:其他加入
func OtherApp() {
	OptionsWares.AddAuthenticationRequire(Ping)
}

func Ws(e *gin.RouterGroup) {
	e.GET("WS", websocket.WsPage)
}

// SignUp:注册
func SignUp(e *gin.RouterGroup) {
	// 注册
	e.POST("SignUp", controller.SignUpHandler)
}

// Login:登录
func Login(e *gin.RouterGroup) {
	e.POST("Login", controller.LoginHandler)
}

// GetVerificationCode:获取验证码
func GetVerificationCode(e *gin.RouterGroup) {
	e.POST("VerificationCode", controller.VerificationCode)
}

// Ping:测试
func Ping(e *gin.RouterGroup) {
	e.GET("ping", func(c *gin.Context) {
		c.JSON(200, "Pong")
	})
	e.POST("ping", func(c *gin.Context) {
		c.JSON(200, "Pong")
	})
}
