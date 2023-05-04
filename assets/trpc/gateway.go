package trpc

import (
	"context"
	"fmt"
	"torch_demo/assets/pbrpc"
)

func getGateClient(insId string) (pbrpc.GatewayClient, error) {
	cc := NodeProxy.GetGateClient(insId)
	if cc == nil {
		return nil, fmt.Errorf("find service %s fail", insId)
	}
	return pbrpc.NewGatewayClient(cc), nil
}

func GateMultiSend(ctx context.Context, userIds []int64, msgId uint32, datas []byte) error {
	// ct, err := getGateClient()
	// if err != nil {
	// 	return err
	// }

	// _, err = ct.MultiSend(ctx, &pbrpc.MultiSendReq{
	// 	UserIds: userIds,
	// 	MsgId:   msgId,
	// 	Datas:   datas,
	// })

	// return err
	return nil
}

func GateBindUser(ctx context.Context, insId string, connId, userId int64) error {
	ct, err := getGateClient(insId)
	if err != nil {
		return err
	}

	_, err = ct.BindUser(ctx, &pbrpc.BindUserReq{
		ConnId: connId,
		UserId: userId,
	})

	return err
}
