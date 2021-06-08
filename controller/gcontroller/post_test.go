package gcontroller

import (
	pb "Happy/model/pmodel/post"
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestPost_CreatePost(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务

	c := pb.NewPostClient(conn)
	// 调用服务
	if err != nil {
		fmt.Println("err", err)
	}
	// auth
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozMTMxNjM1MTM4NTYsImV4cCI6MTYwMDA3NTczMiwiaWF0IjoxNjAwMDcyMTMyLCJpc3MiOiJIYXBweSJ9._w_YWxhGMev6zGAV7-tOZmTAlzZKDcHNpOTd_rn1wOA")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4NDk4NzQ4LCJpYXQiOjE1OTg0OTgxNDgsImlzcyI6IkhhcHB5In0.4tgI8CEaLrO85Ec51kuSvG4wb9d6dfiroaZifxjaeEI")
	//ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.CreatePost(ctx, &pb.CreatePostRequest{
		CommunityID: 1,
		Title:       "Test1",
		Content:     "Test1111111",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", r)
}

func TestPost_PostList(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务

	c := pb.NewPostClient(conn)
	// 调用服务
	if err != nil {
		fmt.Println("err", err)
	}
	// auth
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNjAwMDUzNzMyLCJpYXQiOjE2MDAwNTAxMzIsImlzcyI6IkhhcHB5In0.M1tQvcF1LZc85wCFMXh-cXdPdy4J6csrXgkXozvsBJk")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.PostList(ctx, &pb.GetPostListRequest{
		Model: 1,
		ID:    &pb.GetPostListRequest_AuthorID{AuthorID: 91234500608},
		Page:  1,
		Max:   3,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", r)
}

func TestPost_GetPostDetail(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务

	c := pb.NewPostClient(conn)
	// 调用服务
	if err != nil {
		fmt.Println("err", err)
	}
	// auth
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozMTMxNjM1MTM4NTYsImV4cCI6MTYwMDA3NTczMiwiaWF0IjoxNjAwMDcyMTMyLCJpc3MiOiJIYXBweSJ9._w_YWxhGMev6zGAV7-tOZmTAlzZKDcHNpOTd_rn1wOA")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.GetPostDetail(ctx, &pb.GetPostDetailRequest{
		PostID: 4448263520387073,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", r)
}
