package main

import (
	"github.com/wlf92/torch"
	"github.com/wlf92/torch/pkg/log"
	"github.com/wlf92/torch/registry/consul"
)

func main() {
	reg := consul.NewRegistry()

	nd := torch.NewNode("hall")
	nd.SetRegistry(reg)
	nd.AddRouteHandler(1, Hi)
	nd.AddRouteHandler(2, Hi)

	container := torch.NewContainer(nd)
	container.Serve()
}

func Hi(channel, area int32, uid int64, bts []byte) []byte {
	log.Infow("????")

	return []byte("12345")
}
