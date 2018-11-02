package main

import (
	"context"
	"github.com/micro/go-micro"
	"golang.org/x/crypto/bcrypt"
	"log"
	pb "mymicro/shippy/user-service/proto/user"
)

const topic = "user.created"

type handler struct {
	repo         Repository
	tokenService Authable
	Publisher    micro.Publisher
}

func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) (err error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPwd)
	if err := h.repo.Create(req); err != nil {
		return nil
	}
	resp.User = req

	if err := h.Publisher.Publish(ctx, req); err != nil {
		return err
	}
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
	// 在 part3 中直接传参 &pb.User 去查找用户
	// 会导致 req 的值完全是数据库中的记录值
	// 即 req.Password 与 u.Password 都是加密后的密码
	// 将无法通过验证
	u, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		log.Println("获取用户失败", err)
		return err
	}

	// 进行密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return err
	}
	t, err := h.tokenService.Encode(u)
	if err != nil {
		return err
	}
	resp.Token = t
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) (err error) {

	return nil
}
