package handler

import (
	"context"
	"encoding/json"
	oauth "github.com/entere/micro-examples/part6/oauth-srv/proto/oauth"
	"github.com/entere/micro-examples/part6/plugins/jwt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"time"
)

var (
	oauthClient oauth.OauthService
	jwtClient   = &jwt.Jwt{}
)

//type Error struct {
//	Code   string `json:"code"`
//	Detail string `json:"detail"`
//}

func Init() {
	oauthClient = oauth.NewOauthService("io.github.entere.srv.oauth", client.DefaultClient)
}

type Error struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	rsp, err := oauthClient.LoginByPassword(context.TODO(), &oauth.LoginRequest{
		LoginName: r.Form.Get("loginName"),
		LoginPwd:  r.Form.Get("loginPwd"),
	})
	if err != nil {
		log.Logf("oauthClient.LoginByPassword 请求失败")
		http.Error(w, err.Error(), 500)
		return
	}
	token, err := jwtClient.Encode(rsp.Data.UserID)
	if err != nil {
		log.Logf("jwtClient.Encode 解析失败")
		http.Error(w, err.Error(), 500)
		return
	}
	rsp.Data.Token = token
	// we want to augment the response
	response["error"] = &Error{
		Code: 200,
		Msg:  "success",
	}
	response["data"] = rsp.Data

	// 返回JSON结构
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}
