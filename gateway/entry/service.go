package entry

import (
	"context"
	"torch_demo/assets/pbrpc"

	"github.com/wlf92/torch/cluster/gate"
)

type Service struct {
	pbrpc.UnimplementedGatewayServer
}

func (svc *Service) MultiSend(ctx context.Context, req *pbrpc.MultiSendReq) (*pbrpc.Empty, error) {
	gate.BroadCast(req.GetUserIds(), req.GetMsgId(), req.GetDatas())

	rsp := new(pbrpc.Empty)
	return rsp, nil
}

func (svc *Service) BindUser(ctx context.Context, req *pbrpc.BindUserReq) (*pbrpc.Empty, error) {
	err := gate.BindUser(req.GetConnId(), req.GetUserId())
	if err != nil {
		return nil, err
	}

	rsp := new(pbrpc.Empty)
	return rsp, nil
}
