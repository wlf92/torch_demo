package main

import (
	"context"
	"encoding/json"
	"time"
	"torch_demo/assets/constant"
	"torch_demo/assets/pbcli"
	"torch_demo/assets/trpc"
	"torch_demo/hall/entry"

	"github.com/wlf92/torch"
	"github.com/wlf92/torch/cluster/node"
	"github.com/wlf92/torch/database"
	"github.com/wlf92/torch/mq"
	mqredis "github.com/wlf92/torch/mq/redis"
	"github.com/wlf92/torch/registry/consul"
)

func main() {
	redisClient, err := database.NewRedis()
	if err != nil {
		panic("create redis fail")
	}
	mq.Default = mqredis.Create(context.Background(), redisClient)

	reg := consul.NewRegistry()

	nd := node.Create("hall")
	nd.SetRegistry(reg)
	nd.AddRouteHandler(uint32(pbcli.Msg_Id_LoginReq), entry.LoginReq)

	trpc.NodeProxy = nd

	go func() {
		for {
			bts, _ := json.Marshal(map[string]interface{}{"name": "www"})

			mq.Default.Publish(constant.TopicBroadcast, bts)
			time.Sleep(time.Second * 3)
		}
	}()

	container := torch.NewContainer(nd)
	container.Serve()
}
