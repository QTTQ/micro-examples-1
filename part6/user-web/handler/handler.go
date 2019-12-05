package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"time"

	user "github.com/entere/micro-examples/part6/user-srv/proto/user"
	"github.com/micro/go-micro/client"
)

var (
	userClient user.UserService
)

func Init() {
	userClient = user.NewUserService("io.github.entere.srv.user", client.DefaultClient)
}

func QueryUserByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	// call the backend service

	rsp, err := userClient.QueryUserByName(context.TODO(), &user.QueryUserRequest{
		UserName: r.Form.Get("userName"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	response["data"] = rsp.Data
	response["error"] = rsp.Error

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func QueryUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	// call the backend service

	rsp, err := userClient.QueryUserByID(context.TODO(), &user.QueryUserRequest{
		UserID: r.Form.Get("userID"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	response["data"] = rsp.Data
	response["error"] = rsp.Error

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
