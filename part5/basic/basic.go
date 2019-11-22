package basic

import (
	config2 "github.com/entere/micro-examples/part5/basic/config"
	db2 "github.com/entere/micro-examples/part5/basic/db"
	redis2 "github.com/entere/micro-examples/part5/basic/redis"
)

func Init() {
	config2.Init()
	db2.Init()
	redis2.Init()
}
