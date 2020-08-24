/******
** @创建时间 : 2020/8/23 13:18
** @作者 : SongZhiBin
******/
package gcontroller

import (
	pb "Happy/model/pmodel/jwt"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func TestJwtServer_VerificationRefreshJWT(t *testing.T) {
	// grpc.WithInsecure() 安全参数 可传可不传
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println("监听失败", err)
	}
	defer conn.Close()
	// 创建服务
	c := pb.NewJWTClient(conn)
	// 调用服务
	// c.Greet() .proto 生成go文件的服务方法
	r, err := c.VerificationRefreshJWT(context.TODO(), &pb.VerificationRefreshJWTRequest{Access: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo5MTIzNDUwMDYwOCwiZXhwIjoxNTk4MTYwOTQzLCJpYXQiOjE1OTgxNjAzNDMsImlzcyI6IkhhcHB5In0.uUZN8zN-sAoRsM2BDkRUxkxY5Rs8DoEfCa6L7cyNUiQ", Refresh: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTgxNjA5NDMsImlhdCI6MTU5ODE2MDM0MywiaXNzIjoiSGFwcHkifQ.ein25n82GERKAdrmdC6wUjzD9ApkC22Di8ccTCPsGdc"})
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Printf("%#v", r)
}
