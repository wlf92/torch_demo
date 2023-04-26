package main

import (
	"context"
	"fmt"
	"torch_demo/assets/pbrpc"

	"github.com/wlf92/torch"
	"github.com/wlf92/torch/network"
	"github.com/wlf92/torch/network/ws"
	"github.com/wlf92/torch/pkg/log"
	"github.com/wlf92/torch/registry/consul"
)

func main() {
	reg := consul.NewRegistry()

	server := ws.NewServer()
	server.OnConnect(func(conn network.Conn) {
		log.Infow(fmt.Sprintf("%d %d", conn.ID(), conn.UID()))
	})

	gw := torch.NewGateway()
	gw.SetServer(server)
	gw.SetRegistry(reg)
	gw.SetRpcService(&pbrpc.Gateway_ServiceDesc, &Service{})

	container := torch.NewContainer(gw)
	container.Serve()
}

type Service struct {
	pbrpc.UnimplementedGatewayServer
}

func (svc *Service) Broadcast(context.Context, *pbrpc.BroadcastReq) (*pbrpc.BroadcastRsp, error) {
	return nil, nil
}
