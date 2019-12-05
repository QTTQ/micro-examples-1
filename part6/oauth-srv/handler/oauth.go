package handler

import (
	"context"
	"strings"

	oas "github.com/entere/micro-examples/part6/oauth-srv/model/oauth"
	oauth "github.com/entere/micro-examples/part6/oauth-srv/proto/oauth"
	"github.com/micro/go-micro/util/log"
)

var (
	oauthService oas.Service
)

// Init 初始化handler
func Init() {
	var err error
	oauthService, err = oas.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

type Oauth struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Oauth) LoginByPassword(ctx context.Context, req *oauth.LoginRequest, rsp *oauth.LoginResponse) error {
	loginName := strings.Trim(req.LoginName, " ")
	loginPwd := strings.Trim(req.LoginPwd, " ")
	ret, err := oauthService.LoginByPassword(loginName, loginPwd)
	cr := strings.Compare(loginPwd, ret.LoginPwd)
	if cr != 0 {
		rsp.Error = &oauth.Error{
			Code: 401,
			Msg:  "密码错误",
		}
		return err
	}
	if err != nil {
		rsp.Error = &oauth.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		return err
	}
	rsp.Error = &oauth.Error{
		Code: 200,
		Msg:  "success",
	}
	rsp.Data = ret
	return nil
}
