package gcontroller

import (
	"Happy/dao/redis"
	"Happy/middleware"
	"Happy/model/gmodel"
	"Happy/model/model"
	pb "Happy/model/pmodel/jwt"
	"Happy/pkg/jwt"
	"Happy/settings"
	"context"
)

// JwtServer:定义jwtServer
type JwtServer struct{}

// VerificationRefreshJWT 刷新JWT接口
func (j *JwtServer) VerificationRefreshJWT(ctx context.Context, request *pb.VerificationRefreshJWTRequest) (*pb.Response, error) {
	// 校验请求参数
	res, err := _verification(request)
	if err != nil {
		return (*pb.Response)(res), nil
	}
	auth, err := jwt.ParseRFToken(request.Access, request.Refresh)
	if err != nil {
		if err != model.CodeJWTExpired.Err() {
			return (*pb.Response)(gmodel.ResponseError(model.CodeJWTExpired)), nil

		}
		return (*pb.Response)(gmodel.ResponseError(model.CodeJWTVerificationFailed)), nil
	}
	// 判断是否是最近一次的token
	ok := redis.GetToken(int64(auth.Uid), request.Access)
	if !ok {
		return (*pb.Response)(gmodel.ResponseError(model.CodeMultiTerminalLogin)), nil
	}
	// 生成新access token
	newToken, err := jwt.GetJWT(auth.Uid)
	if err != nil {
		return (*pb.Response)(gmodel.ResponseError(model.CodeServerBusy)), nil
	}
	// 生成新的token放入redis进行缓存
	redis.SetToken(int64(auth.Uid), newToken)
	return (*pb.Response)(gmodel.ResponseSuccess(map[string]string{middleware.AccessToken: newToken})), nil
}

// VerificationJWT 校验JWT
func (j *JwtServer) VerificationJWT(ctx context.Context, request *pb.VerificationJWTRequest) (*pb.VerificationJWTResponse, error) {
	// 校验请求参数
	_, err := _verification(request)
	if err != nil {
		// 未通过校验
		return &pb.VerificationJWTResponse{State: pb.VerificationJWTResponse_NotPass}, nil
	}
	// 判断是否有效
	auth, err := jwt.ParseJWT(request.Access)
	if err != nil {
		if err != model.CodeJWTExpired.Err() {
			// 不是因为过期导致
			return &pb.VerificationJWTResponse{State: pb.VerificationJWTResponse_NotPass}, nil
		}
		// 因为过期 判断模式是否是有refreshToken
		if settings.GetString("JWT.Mode") == "refresh" {
			// todo:如果配置文件是refresh 对应要提供一个接口用于刷新token
			// 表示是refresh模式
			// 告诉前端accessToken过期,需要携带refreshToken进行二次校验
			return &pb.VerificationJWTResponse{State: pb.VerificationJWTResponse_ExpiredJump}, nil
		}
		return &pb.VerificationJWTResponse{State: pb.VerificationJWTResponse_Expired}, nil
	}
	// 从redis获取token判断是否一致
	if !redis.GetToken(int64(auth.Uid), request.Access) {
		return &pb.VerificationJWTResponse{State: pb.VerificationJWTResponse_MultiTerminalLogin}, nil
	}
	// 上下文存入
	return &pb.VerificationJWTResponse{State: pb.VerificationJWTResponse_Pass, Uid: int64(auth.Uid)}, nil
}
