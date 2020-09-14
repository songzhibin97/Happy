/******
** @创建时间 : 2020/9/14 15:06
** @作者 : SongZhiBin
******/
package gcontroller

import (
	pb "Happy/model/pmodel/vote"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestVote_Vote(t *testing.T) {
	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务

	c := pb.NewVoteClient(conn)
	// 调用服务
	if err != nil {
		fmt.Println("err", err)
	}
	// auth
	md := metadata.Pairs("authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNjAwMDc1NzI3LCJpYXQiOjE2MDAwNzIxMjcsImlzcyI6IkhhcHB5In0.wlZqWhvjvYbB3_Wup4ns1vapemQMBlKK4ZrOFI4szQE")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.Vote(ctx, &pb.VoteRequest{
		PostID: 4451385391185921,
		Mode:   0,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", r)
}
