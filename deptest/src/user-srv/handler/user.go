package handler

import (
	"golang.org/x/net/context"

	"log"
	"../../pb"
	"../entity"
	"../db"
)

type UserHandler struct {
	//logger *zap.Logger
}

// new一个UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		//logger: tlog.Instance().Named("UserHandler"),
	}
}

// 增
func (c *UserHandler) InsertUser(ctx context.Context, req *pb.InsertUserReq, rsp *pb.InsertUserResp) error {

	log.Println("InsertUser ...")
	user := &entity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}

	insertId, err := db.InsertUser(user)
	if err != nil {
		log.Fatal("user.db.InsertUser has a error")
		return err
	}
	rsp.Id = int32(insertId)

	return nil
}

// 删
func (c *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserReq, rsp *pb.DeleteUserResp) error {
	log.Println("DeleteUser ...")
	// 待操作
	err := db.DeleteUser(req.GetId())
	if err != nil {
		log.Fatal("user.db.DeleteUser has a error")
		return err
	}
	return nil
}

// 查
func (c *UserHandler) SelectUser(ctx context.Context, req *pb.SelectUserReq, rsp *pb.SelectUserResp) error {

	//c.logger.Info("SelectUser ... ")
	user, err := db.SelectUserById(req.GetId())
	if err != nil {
		log.Fatal("user.db.SelectUserByUId has a error")
		return err
	}
	if user != nil {
		rsp.Users = user.ToProtoUser()
	}
	return nil
}

//改
func (c *UserHandler) ModifyUser(ctx context.Context, req *pb.ModifyUserReq, rsp *pb.ModifyUserResp) error {

	user := &entity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Id:      req.Id,
	}

	log.Println("ModifyUser ...", req.GetId())
	err := db.ModifyUser(user)
	if err != nil {
		log.Fatal("user.db.SelectUserByUId通过id获取用户信息出错")
		return err
	}

	return nil
}
