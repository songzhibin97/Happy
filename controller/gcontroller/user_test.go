/******
** @创建时间 : 2020/8/22 13:38
** @作者 : SongZhiBin
******/
package gcontroller

import (
	pb "Happy/model/pmodel/user"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

// 注册
func TestUserServer_Register(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务
	c := pb.NewUserClient(conn)
	// 调用服务
	// c.Greet() .proto 生成go文件的服务方法
	r, err := c.Register(context.TODO(), &pb.RegisterRequest{UserName: "Test1"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Printf("%#v", r)
}

// 登录
func TestUserServer_Login(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务
	c := pb.NewUserClient(conn)
	// 调用服务
	// c.Greet() .proto 生成go文件的服务方法
	r, err := c.Login(context.TODO(), &pb.LoginRequest{UserName: "Test1", Password: "123456"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Printf("%#v", r)
}

// 验证码
func TestUserServer_Verification(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
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
