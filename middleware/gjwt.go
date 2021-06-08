package middleware

import (
	"Happy/dao/redis"
	"Happy/model/model"
	"Happy/pkg/jwt"
	"Happy/settings"
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
)

// gwt的中间件
func GVerificationJWT(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		zap.L().Info(token, zap.Error(err))
		return nil, err
	}
	auth, err := jwt.ParseJWT(token)
	if err != nil {
		if err != model.CodeJWTExpired.Err() {
			// 不是因为过期导致
			return nil, model.CodeInvalidToken.Err()
		}
		// 因为过期 判断模式是否是有refreshToken
		if settings.GetString("JWT.Mode") == "refresh" {
			// todo:如果配置文件是refresh 对应要提供一个接口用于刷新token
			// 表示是refresh模式
			// 告诉前端accessToken过期,需要携带refreshToken进行二次校验
			return nil, model.CodeAccessExpired.Err()
		}
		return nil, model.CodeJWTExpired.Err()
	}
	// 从redis获取token判断是否一致
	if !redis.GetToken(int64(auth.Uid), token) {
		return nil, model.CodeMultiTerminalLogin.Err()
	}
	newCtx := context.WithValue(ctx, ConTextUserID, auth.Uid)
	return newCtx, nil
}
