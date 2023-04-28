package trpc

import (
	"context"
	"fmt"
	"torch_demo/assets/pbrpc"
)

func getGateClient() (pbrpc.GatewayClient, error) {
	cc := NodeInstance.GetServiceClient("gate")
	if cc == nil {
		return nil, fmt.Errorf("find service %s fail", "gate")
	}
	return pbrpc.NewGatewayClient(cc), nil
}

func GateMultiSend(ctx context.Context, userIds []int64, msgId uint32, datas []byte) error {
	ct, err := getGateClient()
	if err != nil {
		return err
	}

	_, err = ct.MultiSend(ctx, &pbrpc.MultiSendReq{
		UserIds: userIds,
		MsgId:   msgId,
		Datas:   datas,
	})

	return err
}

func GateBindUser(ctx context.Context, connId, userId int64) error {
	ct, err := getGateClient()
	if err != nil {
		return err
	}

	_, err = ct.BindUser(ctx, &pbrpc.BindUserReq{
		ConnId: connId,
		UserId: userId,
	})

	return err
}
