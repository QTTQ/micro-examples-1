package handler

import (
	"context"
	"encoding/json"
	"github.com/go-log/log"

	"net/http"
	"time"

	auth "github.com/entere/micro-examples/part4/auth/srv/proto/auth"
	user "github.com/entere/micro-examples/part4/user/srv/proto/user"
	"github.com/micro/go-micro/client"
)

var (
	userClient user.UserService
	authClient auth.AuthService
)

func Init() {
	userClient = user.NewUserService("io.github.entere.srv.user", client.DefaultClient)
	authClient = auth.NewAuthService("io.github.entere.srv.auth", client.DefaultClient)

}

type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	rsp, err := userClient.QueryUserByName(context.TODO(), &user.Request{
		UserName: r.Form.Get("userName"),
	})
	log.Log(rsp)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	if rsp.User.Pwd == r.Form.Get("userPwd") {
		response["success"] = true
		rsp.User.Pwd = ""
		response["data"] = rsp.User

		rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
			UserId:   rsp.User.Id,
			UserName: rsp.User.Name,
		})
		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}
		response["token"] = rsp2.Token
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)
	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Logf("token获取失败")
		http.Error(w, "非法请求", 400)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.Request{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
