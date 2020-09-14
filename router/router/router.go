/******
** @创建时间 : 2020/8/11 20:49
** @作者 : SongZhiBin
******/
package router

import (
	"Happy/controller/controller"
	_ "Happy/docs"
	"Happy/logger"
	"Happy/middleware"
	"Happy/model/model"
	"Happy/pkg/websocket"
	"Happy/settings"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
)

const (
	NoAuthenticationRequire = "NoAuthenticationRequire"
	AuthenticationRequire   = "AuthenticationRequire"
)

// Option:func(group *gin.RouterGroup)别名
type Option func(group *gin.RouterGroup)

// Options:sliceOption
type Options []Option

// OptionsWare:key-value存储Options
type OptionsWare map[string]Options

// 声明全局变量
var OptionsWares = make(OptionsWare)

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

// LoadAll:封装遍历增加
func (o *OptionsWare) LoadAll(group *gin.RouterGroup, key string) {
	for _, v := range (*o)[key] {
		v(group)
	}
}

// 注册路由
func SetUp() {
	controller.Init()
	// 如果使用共存需要返回 *gin.Engine
	if settings.GetString("APP.Mode") == "release" {
		// 发布版本
		gin.SetMode(gin.ReleaseMode)
	}
	// 避免重复注册翻译器
	//// 注册validator错误翻译器
	//err := controller.InitTrans(settings.GetString("APP.Language"))
	//if err != nil {
	//	zap.L().Error("controller.InitTrans Error", zap.Error(err))
	//}
	r := gin.New()
	// 注册中间件
	r.Use(logger.GinLogger(zap.L()), logger.GinRecovery(zap.L(), true))
	// todo 注册自己的业务路由
	InitOptionsWare()
	// 存放不需要验证身份的路由
	nar := r.Group("/")
	{
		OptionsWares.LoadAll(nar, NoAuthenticationRequire)
	}
	// 存放需要验证身份的路由
	ar := r.Group("/")
	ar.Use(middleware.VerificationJWT)
	{
		OptionsWares.LoadAll(ar, AuthenticationRequire)
	}
	r.NoRoute(func(c *gin.Context) {
		model.ResponseErrorWithMsg(c, model.Code404, gin.H{"访问异常": "没有相关资源,请检查URL"})
	})
	defer controller.CloseConn()
	if err := endless.ListenAndServe(":"+settings.GetString("APP.Port"), r); err != nil {
		zap.L().Error("listen error", zap.Error(err))
	}
	zap.L().Debug("Server exiting")
	//return r
}

// InitOptionsWare:初始化optionWare
func InitOptionsWare() {
	OptionsWares[NoAuthenticationRequire] = make(Options, 0)
	OptionsWares[AuthenticationRequire] = make(Options, 0)
	// 加入登录注册
	internalAdd()
	OtherApp()

}

// 用户相关路由

// internalAdd:加入注册登录
func internalAdd() {
	OptionsWares.AddNoAuthenticationRequire(Ws, User, GetVerificationCode)
	if settings.GetString("JWT.Mode") == "refresh" {
		OptionsWares.AddNoAuthenticationRequire(Refresh)
	}
	if settings.GetString("App.Mode") == "dev" {
		OptionsWares.AddNoAuthenticationRequire(Swagger)
	}
}

// OtherApp:其他加入
func OtherApp() {
	OptionsWares.AddAuthenticationRequire(Ping, Community, Post, Vote)
}

// =========== function ==========

// Ws:websocket
func Ws(e *gin.RouterGroup) {
	e.GET("WS", websocket.WsPage)
}

// Refresh:刷新路由
func Refresh(e *gin.RouterGroup) {
	e.POST("refresh", middleware.VerificationRefreshJWT)
}

// GetVerificationCode:获取验证码
func GetVerificationCode(e *gin.RouterGroup) {
	e.POST("VerificationCode", controller.VerificationCode)
}

// 用户相关

// User:用户相关
func User(e *gin.RouterGroup) {
	// 注册
	e.POST("SignUp", controller.SignUpHandler)
	e.POST("Login", controller.LoginHandler)
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

// 社区相关

// Community:社区相关
func Community(e *gin.RouterGroup) {
	e.POST("communityList", controller.CommunityList)
	e.POST("communityDetail", controller.CommunityDetail)
}

// Post:帖子相关
func Post(e *gin.RouterGroup) {
	e.GET("PostList", controller.PostList)
	e.GET("GetPostDetail", controller.GetPostDetail)
	e.POST("CreatePost", controller.CreatePost)
}

// Vote:投票相关
func Vote(e *gin.RouterGroup) {
	e.POST("vote", controller.Vote)
}

// Swagger:接口相关
func Swagger(e *gin.RouterGroup) {
	e.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
}
