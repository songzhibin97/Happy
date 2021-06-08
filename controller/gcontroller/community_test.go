package gcontroller

import (
	pb "Happy/model/pmodel/community"
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozMTMxNjM1MTM4NTYsImV4cCI6MTYwMDIzMTEzNCwiaWF0IjoxNjAwMjI3NTM0LCJpc3MiOiJIYXBweSJ9.LMfptXYXU8u9nMyxT-4M_YNaGMjSTs_iAie-mJ5HA0A")
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
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozMTMxNjM1MTM4NTYsImV4cCI6MTYwMDIzMTEzNCwiaWF0IjoxNjAwMjI3NTM0LCJpc3MiOiJIYXBweSJ9.LMfptXYXU8u9nMyxT-4M_YNaGMjSTs_iAie-mJ5HA0A")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.CommunityDetail(ctx, &pb.CommunityDetailRequest{ID: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", r)
}
