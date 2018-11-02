package main

import (
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"log"
	pb "mymicro/shippy/user-service/proto/user"
	"os"
)

func main() {

	cmd.Init()
	// 创建 user-service 微服务的客户端
	client := pb.NewUserServiceClient("UserService", microclient.DefaultClient)

	// 暂时将用户信息写死在代码中
	name := "zhang yuehao"
	email := "3313246150@11.com"
	password := "test123456"
	company := "zcm"

	resp, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("call Create error: %v", err)
	}
	log.Println("created: ", resp.User.Id)

	allResp, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("call GetAll error: %v", err)
	}
	for i, u := range allResp.Users {
		log.Printf("user_%d: %v\n", i, u)
	}

	authResp, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("auth failed: %v", err)
	}
	log.Println("token: ", authResp.Token)

	// 直接退出即可
	os.Exit(0)
}
