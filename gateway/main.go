package main

import (
	"context"
	"fmt"
	"torch_demo/assets/constant"
	"torch_demo/assets/pbcli"
	"torch_demo/assets/pbrpc"
	"torch_demo/gateway/entry"

	"github.com/wlf92/torch"
	"github.com/wlf92/torch/cluster/gate"
	"github.com/wlf92/torch/database"
	"github.com/wlf92/torch/mq"
	mqredis "github.com/wlf92/torch/mq/redis"
	"github.com/wlf92/torch/network"
	"github.com/wlf92/torch/network/ws"
	"github.com/wlf92/torch/packet"
	"github.com/wlf92/torch/pkg/log"
	"github.com/wlf92/torch/registry/consul"
	"google.golang.org/protobuf/proto"
)

func main() {
	redisClient, err := database.NewRedis()
	if err != nil {
		panic("create redis fail")
	}
	mq.Default = mqredis.Create(context.Background(), redisClient)
	mq.Default.Subscribe(constant.TopicBroadcast, "", func(message []byte) error {
		fmt.Println(string(message))
		return nil
	})

	server := ws.NewServer()
	server.OnConnect(func(conn network.Conn) {
		log.Infow(fmt.Sprintf("%d %d", conn.ID(), conn.UID()))
	})

	registry := consul.NewRegistry()

	gw := gate.Create()
	gw.SetServer(server)
	gw.SetRegistry(registry)
	gw.SetRpcService(&pbrpc.Gateway_ServiceDesc, &entry.Service{})
	gw.SetErrorHandler(func(conn network.Conn, err error) {
		send := &pbcli.ErrorNtf{Code: 1, Msg: err.Error()}
		bts, _ := proto.Marshal(send)
		bts = packet.Pack(&packet.Message{Route: uint32(pbcli.Msg_Id_ErrorNtf), Buffer: bts})
		conn.Send(bts)
	})

	container := torch.NewContainer(gw)
	container.Serve()
}
