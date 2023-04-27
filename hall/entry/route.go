package entry

import (
	"context"
	"torch_demo/assets/pbcli"
	"torch_demo/assets/trpc"
)

func LoginReq(uid int64, req *pbcli.LoginReq) *pbcli.LoginRsp {
	rsp := new(pbcli.LoginRsp)

	trpc.GateBindUser(context.Background(), 12345, req.GetConnId())

	rsp.UserId = 12345

	return rsp
}
