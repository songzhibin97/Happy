/******
** @创建时间 : 2020/8/22 09:34
** @作者 : SongZhiBin
******/
package grouter

import (
	"Happy/controller/gcontroller"
	pb2 "Happy/model/pmodel/jwt"
	pb "Happy/model/pmodel/user"
	"Happy/settings"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// :注册服务分块

// GrpcOption:func(server *grpc.Server)别名
type GrpcOption func(server *grpc.Server)

// GrpcOptions:sliceGrpcOption
type GrpcOptions []GrpcOption

var GrpcOptionsWares = make(GrpcOptions, 0)

// AddGrpcOptionsWares:添加
func (g *GrpcOptions) AddGrpcOptionsWares(grpcOptions ...GrpcOption) {
	*g = append(*g, grpcOptions...)
}

// LoadAll:加载所有服务
func (g *GrpcOptions) LoadAll(server *grpc.Server) {
	for _, v := range *g {
		v(server)
	}
}

// grpc router
func GrpcSetUp() {
	// 如果使用共存需要返回 *grpc.Server
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(settings.GetInt("APP.Port")+1))
	if err != nil {
		zap.L().Error("Listen Error", zap.Error(err))
		return
	}
	grpcServer := grpc.NewServer()
	grpcInternalAdd()
	GrpcOptionsWares.LoadAll(grpcServer)
	// 在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(grpcServer)

	// 协程启动
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			zap.L().Error("grpcServer Serve Error", zap.Error(err))
			panic(err)
		}
	}()
	// 自动优雅停止，tpc避免占用资源
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			grpcServer.GracefulStop()
			return
		}
	}
	//return grpcServer
}

// 相关路由
func grpcInternalAdd() {
	GrpcOptionsWares.AddGrpcOptionsWares(User)
	// 如果是refresh模式需要额外注册一个服务
	if settings.GetString("JWT.Mode") == "refresh" {
		GrpcOptionsWares.AddGrpcOptionsWares(Jwt)
	}
}

// ======== function =======

// User:用户相关grpc服务
func User(server *grpc.Server) {
	pb.RegisterUserServer(server, new(gcontroller.UserServer))
}

// Jwt:json-web-token相关grpc服务
func Jwt(server *grpc.Server) {
	pb2.RegisterJWTServer(server, new(gcontroller.JwtServer))
}
