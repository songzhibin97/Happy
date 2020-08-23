/******
** @创建时间 : 2020/8/17 16:45
** @作者 : SongZhiBin
******/
package middleware

import (
	"Happy/dao/redis"
	"Happy/model"
	"Happy/pkg/jwt"
	"Happy/settings"
	"github.com/gin-gonic/gin"
)

// 关于jwt的中间件

const (
	ConTextUserID = "UID"
	AccessToken   = "accessJWT"
	RefreshToken  = "refreshJWT"
)

// VerificationJWT:判断JWT是否有效
func VerificationJWT(c *gin.Context) {
	// 从请求头获取jwt
	access := c.Request.Header.Get(AccessToken)
	if access == "" {
		// 未携带参数
		model.ResponseError(c, model.CodeInvalidToken)
		c.Abort()
		return
	}
	// 判断access的是否有效
	auth, err := jwt.ParseJWT(access)
	if err != nil {
		if err != model.CodeJWTExpired.Err() {
			// 不是因为过期导致
			model.ResponseError(c, model.CodeInvalidToken)
			c.Abort()
			return
		}
		// 因为过期 判断模式是否是有refreshToken
		if settings.GetString("JWT.Mode") == "refresh" {
			// todo:如果配置文件是refresh 对应要提供一个接口用于刷新token
			// 表示是refresh模式
			// 告诉前端accessToken过期,需要携带refreshToken进行二次校验
			model.ResponseError(c, model.CodeAccessExpired)
			c.Abort()
			return
		}
		model.ResponseError(c, model.CodeJWTExpired)
		c.Abort()
		return
	}
	// 从redis获取token判断是否一致
	if !redis.GetToken(int64(auth.Uid), access) {
		model.ResponseError(c, model.CodeMultiTerminalLogin)
		c.Abort()
		return
	}
	// jwt验证成功,将uid存储上下文
	c.Set(ConTextUserID, auth.Uid)
	// 放行
	c.Next()
}

// VerificationRefreshJWT:用于校验刷新token
func VerificationRefreshJWT(c *gin.Context) {
	access := c.Request.Header.Get(AccessToken)
	refresh := c.Request.Header.Get(RefreshToken)
	if access == "" || refresh == "" {
		// 未携带参数
		model.ResponseError(c, model.CodeInvalidToken)
		return
	}
	auth, err := jwt.ParseRFToken(access, refresh)
	if err != nil {
		if err != model.CodeJWTExpired.Err() {
			model.ResponseError(c, model.CodeJWTExpired)
			return
		}
		model.ResponseError(c, model.CodeJWTVerificationFailed)
		return
	}
	// 判断是否是最近一次的token
	ok := redis.GetToken(int64(auth.Uid), access)
	if !ok {
		model.ResponseError(c, model.CodeMultiTerminalLogin)
		return
	}
	// 生成新access token
	newToken, err := jwt.GetJWT(auth.Uid)
	if err != nil {
		model.ResponseError(c, model.CodeServerBusy)
		return
	}
	// 生成新的token放入redis进行缓存
	redis.SetToken(int64(auth.Uid), newToken)
	model.ResponseSuccess(c, gin.H{
		"accessToken": newToken,
	})
	return
}
