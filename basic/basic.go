package basic

import (
	"github.com/entere/micro-examples/basic/config"
	"github.com/entere/micro-examples/basic/db"
	"github.com/entere/micro-examples/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
