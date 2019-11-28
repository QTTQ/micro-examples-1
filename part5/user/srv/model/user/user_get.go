package user

import (
	"github.com/entere/micro-examples/part5/plugins/db"
	proto "github.com/entere/micro-examples/part5/user/srv/proto/user"
	"github.com/micro/go-micro/util/log"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id,login_name,login_password FROM auth_passwords WHERE login_name = ?`
	o := db.GetDB()
	ret = &proto.User{}
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}

func (s *service) QueryUserByID(userID string) (ret *proto.User, err error) {
	queryString := `SELECT user_id,login_name FROM auth_passwords WHERE user_id = ?`
	o := db.GetDB()
	ret = &proto.User{}
	err = o.QueryRow(queryString, userID).Scan(&ret.Id, &ret.Name)
	if err != nil {
		log.Logf("[QueryUserByID] 查询数据失败，err：%s", err)
		return
	}
	return
}
