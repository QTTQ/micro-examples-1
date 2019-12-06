package handler

import (
	"fmt"

	"github.com/entere/micro-examples/part6/plugins/jwt"
	"github.com/micro/go-micro/util/log"
	"net/http"
)

var (
	cfg = &jwt.Jwt{}
)

func JWTAuthWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO 从配置中心动态获取白名单URL
		if r.URL.Path == "/oauth/login" || r.URL.Path == "/oauth/register" || r.URL.Path == "/oauth/test" {
			h.ServeHTTP(w, r)
			return
		}

		header := r.Header.Get("Authorization")
		var tk string
		fmt.Sscanf(header, "Bearer %s", &tk)

		//token.Init()
		//tokenService, err := token.GetService()
		//if err != nil {
		//	log.Log(err)
		//}
		cc, err := cfg.Decode(tk)
		if err != nil {
			log.Log(err)
		}
		log.Log(cc)
		log.Logf("token:" + tk)
		w.Header().Add("uid", cc.UserID)
		h.ServeHTTP(w, r)
	})
}
