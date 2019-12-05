package handler

import (
	"context"

	us "github.com/entere/micro-examples/part6/user-srv/model/user"
	user "github.com/entere/micro-examples/part6/user-srv/proto/user"
	"github.com/micro/go-micro/util/log"
)

var (
	userService us.Service
)

// Init 初始化handler
func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

type User struct{}

// 根据用户名获取用户信息
func (e *User) QueryUserByName(ctx context.Context, req *user.QueryUserRequest, rsp *user.QueryUserResponse) error {
	userData, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Error = &user.Error{
			Code: 500,
			Msg:  err.Error(),
		}

		return err
	}

	rsp.Data = userData
	rsp.Error = &user.Error{
		Code: 200,
		Msg:  "success",
	}
	return nil

}

// 根据用户ID获取用户信息
func (e *User) QueryUserByID(ctx context.Context, req *user.QueryUserRequest, rsp *user.QueryUserResponse) error {
	userData, err := userService.QueryUserByID(req.UserID)
	if err != nil {
		rsp.Error = &user.Error{
			Code: 500,
			Msg:  err.Error(),
		}

		return err
	}

	rsp.Data = userData
	rsp.Error = &user.Error{
		Code: 200,
		Msg:  "success",
	}
	return nil

}
