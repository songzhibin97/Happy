/******
** @创建时间 : 2020/8/27 17:36
** @作者 : SongZhiBin
******/
package gcontroller

import (
	pb "Happy/model/pmodel/post"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
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
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4NTIyNTQxLCJpYXQiOjE1OTg1MjE5NDEsImlzcyI6IkhhcHB5In0.ikmeqTZE9VZ9qT3p2ldtvz9Vceufs3E3tyMV8V8qzJw")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4NDk4NzQ4LCJpYXQiOjE1OTg0OTgxNDgsImlzcyI6IkhhcHB5In0.4tgI8CEaLrO85Ec51kuSvG4wb9d6dfiroaZifxjaeEI")
	//ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.CreatePost(ctx, &pb.CreatePostRequest{
		CommunityID: 2,
		Title:       "test244",
		Content:     "contentTest244",
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
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4Njc0MjUwLCJpYXQiOjE1OTg2NzA2NTAsImlzcyI6IkhhcHB5In0.SjOVYNNWXmgHKGARmZF9w4_mO55hdODkVkH0J18D-j4")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.PostList(ctx, &pb.GetPostListRequest{
		Model: 2,
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
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4NTI3Mjc5LCJpYXQiOjE1OTg1MjM2NzksImlzcyI6IkhhcHB5In0.vWNYyVzJn6pnbIRsBXydFtOBBwn0-jmtcOKvmr2XVhg")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.GetPostDetail(ctx, &pb.GetPostDetailRequest{
		PostID: 248269242368,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", r)
}
