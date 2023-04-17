package main

import (
	"github.com/wlf92/torch"
	"github.com/wlf92/torch/pkg/log"
	"github.com/wlf92/torch/registry/consul"
)

func main() {
	reg := consul.NewRegistry()

	nd := torch.NewNode()
	nd.SetRegistry(reg)
	nd.AddRouteHandler(1, Hi)
	nd.AddRouteHandler(2, Hi)

	container := torch.NewContainer(nd)
	container.Serve()
}

func Hi(p []byte) []byte {
	log.Infow("????")
	return []byte("12345")
}
