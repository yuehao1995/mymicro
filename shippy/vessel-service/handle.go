package main

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	pb "mymicro/shippy/vessel-service/proto/vessel"
)

// 实现微服务的服务端
type handler struct {
	session *mgo.Session
}

func (h *handler) GetRepo() Repository {
	return &VesselRepository{h.session.Clone()}
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	defer h.GetRepo().Close()
	v, err := h.GetRepo().FindAvailable(req)
	if err != nil {
		fmt.Println("获取可用船只失败", err)
		return err
	}
	resp.Vessel = v
	return nil
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) (err error) {
	defer h.GetRepo().Close()
	if err := h.GetRepo().Create(req); err != nil {
		fmt.Println("创建船只失败", err)
		return err
	}
	resp.Vessel = req
	resp.Created = true
	return nil
}
