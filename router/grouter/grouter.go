/******
** @创建时间 : 2020/8/22 09:34
** @作者 : SongZhiBin
******/
package grouter

import (
	"Happy/controller/gcontroller"
	"Happy/middleware"
	pb3 "Happy/model/pmodel/community"
	pb2 "Happy/model/pmodel/jwt"
	pb4 "Happy/model/pmodel/post"
	pb "Happy/model/pmodel/user"
	pb5 "Happy/model/pmodel/vote"
	"Happy/settings"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	NoAuthenticationRequire = "NoAuthenticationRequire" // 不需要认证
	AuthenticationRequire   = "AuthenticationRequire"   // 需要认证
)

// :注册服务分块

// GrpcOption:func(server *grpc.Server)别名
type GrpcOption func(server *grpc.Server)

// GrpcOptions:sliceGrpcOption
type GrpcOptions []GrpcOption

// 新增分类
type GrpcOptionsWare map[string]GrpcOptions

// 更改类型
var GrpcOptionsWares = make(GrpcOptionsWare)

// AddNoAuthenticationRequire:添加不用认证的grpc服务
func (g *GrpcOptionsWare) AddNoAuthenticationRequire(grpcOptions ...GrpcOption) {
	// 1.判断 NoAuthenticationRequire 是否初始化
	_, ok := (*g)[NoAuthenticationRequire]
	if !ok {
		(*g)[NoAuthenticationRequire] = make(GrpcOptions, 0)
	}
	(*g)[NoAuthenticationRequire] = append((*g)[NoAuthenticationRequire], grpcOptions...)
}

// AddAuthenticationRequire:添加需要认证的路由
func (g *GrpcOptionsWare) AddAuthenticationRequire(grpcOptions ...GrpcOption) {
	// 1.判断 NoAuthenticationRequire 是否初始化
	_, ok := (*g)[AuthenticationRequire]
	if !ok {
		(*g)[AuthenticationRequire] = make(GrpcOptions, 0)
	}
	(*g)[AuthenticationRequire] = append((*g)[AuthenticationRequire], grpcOptions...)
}

// LoadAll:加载所有服务
func (g *GrpcOptionsWare) LoadAll(server *grpc.Server, key string) {
	for _, v := range (*g)[key] {
		v(server)
	}
}

// grpc router
func GrpcSetUp() {
	// 如果使用共存需要返回 *grpc.Server
	nAuth, err := net.Listen("tcp", ":"+strconv.Itoa(settings.GetInt("GRPC.NoAuthPort")))
	if err != nil {
		zap.L().Error("Listen Error", zap.Error(err))
		return
	}
	Auth, err2 := net.Listen("tcp", ":"+strconv.Itoa(settings.GetInt("GRPC.AuthPort")))
	if err2 != nil {
		zap.L().Error("Listen Error", zap.Error(err))
		return
	}
	// grpcAuth:需要认证的grpcServer
	grpcNoAuth := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_opentracing.StreamServerInterceptor(),
		grpc_prometheus.StreamServerInterceptor,
		grpc_zap.StreamServerInterceptor(zap.L()),
		//grpc_auth.StreamServerInterceptor(middleware.GVerificationJWT),
		grpc_recovery.StreamServerInterceptor(),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(zap.L()),
			//grpc_auth.UnaryServerInterceptor(middleware.GVerificationJWT),
			grpc_recovery.UnaryServerInterceptor(),
		)))
	grpcAuth := grpc.NewServer(grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_opentracing.StreamServerInterceptor(),
		grpc_prometheus.StreamServerInterceptor,
		grpc_zap.StreamServerInterceptor(zap.L()),
		grpc_auth.StreamServerInterceptor(middleware.GVerificationJWT),
		grpc_recovery.StreamServerInterceptor(),
	)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(zap.L()),
			grpc_auth.UnaryServerInterceptor(middleware.GVerificationJWT),
			grpc_recovery.UnaryServerInterceptor(),
		)))
	grpcInternalAdd()

	GrpcOptionsWares.LoadAll(grpcNoAuth, NoAuthenticationRequire)

	GrpcOptionsWares.LoadAll(grpcAuth, AuthenticationRequire)
	// 在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(grpcNoAuth)
	reflection.Register(grpcAuth)
	// 协程启动
	go func() {
		if err := grpcNoAuth.Serve(nAuth); err != nil {
			zap.L().Error("grpcAuth Serve Error", zap.Error(err))
			panic(err)
		}
	}()
	go func() {
		if err := grpcAuth.Serve(Auth); err != nil {
			zap.L().Error("grpcAuth Serve Error", zap.Error(err))
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
			grpcNoAuth.GracefulStop()
			grpcAuth.GracefulStop()
			return
		}
	}
	//return grpcAuth
}

// 相关路由
func grpcInternalAdd() {
	GrpcOptionsWares.AddNoAuthenticationRequire(User)

	GrpcOptionsWares.AddNoAuthenticationRequire(Jwt)

	GrpcOptionsWares.AddAuthenticationRequire(Community, Post, Vote)
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

// Community:社区grpc服务
func Community(server *grpc.Server) {
	pb3.RegisterCommunityServer(server, new(gcontroller.CommunityServer))
}

// Post:帖子相关
func Post(server *grpc.Server) {
	pb4.RegisterPostServer(server, new(gcontroller.Post))
}

// Vote:投票相关
func Vote(server *grpc.Server) {
	pb5.RegisterVoteServer(server, new(gcontroller.Vote))
}
