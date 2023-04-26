package trpc

import (
	"context"
	"fmt"
	"torch_demo/assets/pbrpc"
)

func getGateClient() (pbrpc.GatewayClient, error) {
	cc := NodeInstance.GetServiceClient("gateway")
	if cc != nil {
		return nil, fmt.Errorf("find service %s fail", "gateway")
	}
	return pbrpc.NewGatewayClient(cc), nil
}

func GateBroadcast(ctx context.Context, userIds []int64, msgId uint32, datas []byte) error {
	ct, err := getGateClient()
	if err != nil {
		return err
	}

	_, err = ct.Broadcast(ctx, &pbrpc.BroadcastReq{
		UserIds: userIds,
		MsgId:   msgId,
		Datas:   datas,
	})

	return err
}
