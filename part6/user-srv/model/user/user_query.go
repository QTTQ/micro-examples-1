package user

import (
	"github.com/entere/micro-examples/part6/plugins/db"
	proto "github.com/entere/micro-examples/part6/user-srv/proto/user"
	"github.com/micro/go-micro/util/log"
)

func (s *service) QueryUserByName(userName string) (ret *proto.UserData, err error) {
	queryString := `SELECT user_id,login_name FROM auth_passwords WHERE login_name = ?`
	o := db.GetDB()
	ret = &proto.UserData{}
	err = o.QueryRow(queryString, userName).Scan(&ret.UserID, &ret.UserName)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}

func (s *service) QueryUserByID(userID string) (ret *proto.UserData, err error) {
	queryString := `SELECT user_id,login_name FROM auth_passwords WHERE login_name = ?`
	o := db.GetDB()
	ret = &proto.UserData{}
	err = o.QueryRow(queryString, userID).Scan(&ret.UserID, &ret.UserName)
	if err != nil {
		log.Logf("[QueryUserByID] 查询数据失败，err：%s", err)
		return
	}
	return
}
