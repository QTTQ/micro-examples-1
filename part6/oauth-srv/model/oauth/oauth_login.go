package oauth

import (
	proto "github.com/entere/micro-examples/part6/oauth-srv/proto/oauth"
	"github.com/entere/micro-examples/part6/plugins/db"
	"github.com/micro/go-micro/util/log"
)

func (s *service) LoginByPassword(loginName string, loginPwd string) (ret *proto.LoginData, err error) {

	queryString := `SELECT user_id,login_name,login_password FROM auth_passwords WHERE login_name = ?`
	o := db.GetDB()
	ret = &proto.LoginData{}
	err = o.QueryRow(queryString, loginName).Scan(&ret.UserID, &ret.LoginName, &ret.LoginPwd)
	if err != nil {
		log.Logf("[LoginByPassword] 查询数据失败，err：%s", err)
		return
	}
	//cr := strings.Compare(loginPwd, ret.LoginPwd)
	//if cr != 0 {
	//	log.Log("[LoginByPassword] 密码错误")
	//
	//	return nil, errors.New("密码错误")
	//}
	return

}
