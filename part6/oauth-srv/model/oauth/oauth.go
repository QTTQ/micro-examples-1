package oauth

import (
	"fmt"
	proto "github.com/entere/micro-examples/part6/oauth-srv/proto/oauth"
	"sync"
)

var (
	s *service
	m sync.RWMutex
)

// 使用用户密码和密码登录
type Service interface {
	LoginByPassword(loginName string, loginPwd string) (ret *proto.LoginData, err error)
}

type service struct {
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// 初始化服务信息
func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	s = &service{}
}
