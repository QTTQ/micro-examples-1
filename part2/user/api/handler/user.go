package handler

import (
	"context"
	"encoding/json"
	client2 "github.com/entere/micro-examples/part2/user/api/client"
	"github.com/entere/micro-examples/part2/user/srv/proto/user"
	"github.com/micro/go-micro/util/log"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
)

type User struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// User.Info is called by the API as /user/call with post body {"name": "foo"}
func (e *User) Info(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received User.Info request")

	// extract the client from the context
	userClient, ok := client2.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("io.github.entere.api.user.user.info", "user client not found")
	}

	// make request
	response, err := userClient.QueryUserByID(ctx, &io_github_entere_srv_user.QueryRequest{
		UserId: extractValue(req.Post["user_id"]),
	})
	if err != nil {
		return errors.InternalServerError("io.github.entere.api.user.user.info", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
