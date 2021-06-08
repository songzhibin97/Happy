package gcontroller

import (
	pb "Happy/model/pmodel/user"
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
)

// Token token认证
type Token struct {
	Value string
}

const headerAuthorize string = "authorization"

// GetRequestMetadata 获取当前请求认证所需的元数据
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{headerAuthorize: t.Value}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (t *Token) RequireTransportSecurity() bool {
	return true
}

// 注册
func TestUserServer_Register(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务
	c := pb.NewUserClient(conn)
	// 调用服务
	// c.Greet() .proto 生成go文件的服务方法
	r, err := c.Register(context.TODO(), &pb.RegisterRequest{UserName: "Test3", Password: "123456", ConfirmPassword: "123456", Email: "718428482@qq.com", VerificationCode: "ZLD3SlK4"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Printf("%#v", r)
}

// 登录
func TestUserServer_Login(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务

	c := pb.NewUserClient(conn)
	// 调用服务
	if err != nil {
		fmt.Println("err", err)
	}
	//// auth
	//md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4NDEyMjE5LCJpYXQiOjE1OTg0MTE2MTksImlzcyI6IkhhcHB5In0.xS__TWRRvXX-HKZasL0g2kVGfZ73sAx3k7y6YAJwt-I")
	//ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.Login(context.TODO(), &pb.LoginRequest{UserName: "Test2", Password: "123456"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Printf("%#v", r)
}

// 验证码
func TestUserServer_Verification(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务
	c := pb.NewUserClient(conn)
	// 调用服务
	// c.Greet() .proto 生成go文件的服务方法
	r, err := c.Verification(context.TODO(), &pb.VerificationRequest{Email: "718428482@qq.com"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Printf("%#v", r)
}
