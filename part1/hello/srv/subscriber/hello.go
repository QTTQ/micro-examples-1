package subscriber

import (
	"context"
	"github.com/entere/micro-examples/part1/hello/srv/proto/hello"
	"github.com/micro/go-micro/util/log"
)

type Hello struct{}

func (e *Hello) Handle(ctx context.Context, msg *io_github_entere_srv_hello.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *io_github_entere_srv_hello.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
