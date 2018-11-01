package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

// 创建与 MongoDB 交互的主回话
func CreateSession(host string) (*mgo.Session, error) {
	s, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("consignment-service连接mongo失败：", err)
		return nil, err
	}
	s.SetMode(mgo.Monotonic, true)
	return s, nil
}
