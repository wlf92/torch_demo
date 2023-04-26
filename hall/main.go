package main

import (
	"torch_demo/assets/pbcli"
	"torch_demo/assets/trpc"

	"github.com/wlf92/torch"
	"github.com/wlf92/torch/pkg/log"
	"github.com/wlf92/torch/registry/consul"
)

func main() {
	reg := consul.NewRegistry()

	nd := torch.NewNode("hall")
	nd.SetRegistry(reg)
	nd.AddRouteHandler(uint32(pbcli.Msg_Id_LoginReq), Hi)

	trpc.NodeInstance = nd

	container := torch.NewContainer(nd)
	container.Serve()
}

func Hi(channel, area int32, uid int64, bts []byte) ([]byte, error) {

	log.Infow("????")

	return []byte("12345"), nil
}
