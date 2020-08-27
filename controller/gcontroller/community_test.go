/******
** @创建时间 : 2020/8/26 21:17
** @作者 : SongZhiBin
******/
package gcontroller

import (
	pb "Happy/model/pmodel/community"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestCommunityServer_CommunityList(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务

	c := pb.NewCommunityClient(conn)
	// 调用服务
	if err != nil {
		fmt.Println("err", err)
	}
	// auth
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4NDk4NzQ4LCJpYXQiOjE1OTg0OTgxNDgsImlzcyI6IkhhcHB5In0.4tgI8CEaLrO85Ec51kuSvG4wb9d6dfiroaZifxjaeEI")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.CommunityList(ctx, &pb.CommunityListRequest{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", r)

}

func TestCommunityServer_CommunityDetail(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务

	c := pb.NewCommunityClient(conn)
	// 调用服务
	if err != nil {
		fmt.Println("err", err)
	}
	//// auth
	//md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4NDEyMjE5LCJpYXQiOjE1OTg0MTE2MTksImlzcyI6IkhhcHB5In0.xS__TWRRvXX-HKZasL0g2kVGfZ73sAx3k7y6YAJwt-I")
	//ctx := metadata.NewOutgoingContext(context.Background(), md)
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4NTI3Mjc5LCJpYXQiOjE1OTg1MjM2NzksImlzcyI6IkhhcHB5In0.vWNYyVzJn6pnbIRsBXydFtOBBwn0-jmtcOKvmr2XVhg")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.CommunityDetail(ctx, &pb.CommunityDetailRequest{ID: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", r)
}
