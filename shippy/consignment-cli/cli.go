package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"io/ioutil"
	"log"
	pb "mymicro/shippy/consignment-service/proto/consignment"
	"os"
)

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_INFO_FILE = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {
	sercice := micro.NewService(micro.Name("ShippingService"))
	sercice.Init()

	// 创建微服务的客户端，简化了手动 Dial 连接服务端的步骤
	client := pb.NewShippingServiceClient("ShippingService", sercice.Client())

	// 在命令行中指定新的货物信息 json 件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiZmQxZGYyYzEtYjEyOC00ZmJmLTkyNDEtOWY3ZDM3ODM3ZjFiIiwibmFtZSI6IkV3YW4gVmFsZW50aW5lIiwiY29tcGFueSI6IkJCQyIsImVtYWlsIjoiZXdhbi52YWx" +
		"lbnRpbmU4OUBnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCQ0bHZ2UXprcWJzMmhGMHVrU0gvVGp1SU5YWHl1VEFPY0hvWTZpRXBkcmQvTE1pU0M1dWZiZSJ9LCJleHAiOjE1NDE0MDYwMDIsImlzcyI6IlVzZXJTZXJ2aWNlIn0.OChJk_Vj6eiUWckgqDTM7GfIs1cofCxdkR6" +
		"R99hASLE"

	// 解析货物信息
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}
	tokenContext := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// 调用 RPC
	// 将货物存储到我们自己的仓库里
	resp, err := client.CreateConsignment(tokenContext, consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}

	// 新货物是否托运成功
	log.Printf("created: %t", resp.Created)
	log.Printf("resp: %v", resp)

	// 列出目前所有托运的货物
	resp, err = client.GetConsignments(tokenContext, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}
	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}
}
