package main

import (
	"fmt"

	"github.com/wlf92/torch"
)

func main() {
	gw := torch.NewGateway()
	fmt.Println(gw.Name())

	node := torch.NewNode()
	fmt.Println(node.Name())
}
