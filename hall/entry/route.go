package entry

import (
	"context"
	"fmt"
	"torch_demo/assets/pbcli"
	"torch_demo/assets/trpc"

	"github.com/wlf92/torch/transport"
)

func LoginReq(head transport.Header, req *pbcli.LoginReq) *pbcli.LoginRsp {
	rsp := new(pbcli.LoginRsp)

	trpc.GateBindUser(context.Background(), 12345, req.GetConnId())

	rsp.UserId = 12345

	fmt.Println("-----", head.GetMsgId())

	return rsp
}
