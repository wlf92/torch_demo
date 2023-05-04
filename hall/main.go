package main

import (
	"torch_demo/assets/pbcli"
	"torch_demo/assets/trpc"
	"torch_demo/hall/entry"

	"github.com/wlf92/torch"
	"github.com/wlf92/torch/cluster/node"
	"github.com/wlf92/torch/registry/consul"
)

func main() {
	reg := consul.NewRegistry()

	nd := node.Create("hall")
	nd.SetRegistry(reg)
	nd.AddRouteHandler(uint32(pbcli.Msg_Id_LoginReq), entry.LoginReq)

	trpc.NodeProxy = nd

	container := torch.NewContainer(nd)
	container.Serve()
}
