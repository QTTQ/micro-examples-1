package user

import (
	"fmt"
	proto "github.com/entere/micro-examples/part6/user-srv/proto/user"
	"sync"
)

var (
	m sync.Mutex
	s *service
)

type service struct {
}

type Service interface {
	QueryUserByName(userName string) (ret *proto.UserData, err error)
	QueryUserByID(userID string) (ret *proto.UserData, err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	s = &service{}

}
