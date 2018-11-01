package main

import (
	"context"
	pb "mymicro/shippy/user-service/proto/user"
)

type handler struct {
	repo Repository
}

func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) (err error) {
	if err = h.repo.Create(req); err != nil {
		return
	}
	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) (err error) {
	u, err := h.repo.Get(req.Id)
	if err != nil {
		return
	}
	resp.User = u
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) (err error) {
	users, err := h.repo.GetAll()
	if err != nil {
		return
	}
	resp.Users = users
	return nil
}

func (h *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) (err error) {
	_, err = h.repo.GetByEmailAndPassword(req)
	if err != nil {
		return
	}
	resp.Token = "`x_2nam"
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) (err error) {

	return nil
}
