package entry

import (
	"torch_demo/assets/pbcli"

	"github.com/wlf92/torch/transport"
)

func LoginReq(head *transport.Header, req *pbcli.LoginReq) *pbcli.LoginRsp {
	rsp := new(pbcli.LoginRsp)

	// trpc.GateBindUser(context.Background(), head.GetGateId(), head.GetConnId(), 12345)

	rsp.UserId = 12345

	// fmt.Println("-----", head.GetMsgId(), head.GetConnId(), head.GetUserId(), head.GetGateId())

	return rsp
}
