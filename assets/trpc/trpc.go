package trpc

import "google.golang.org/grpc"

type ISelector interface {
	GetServiceClient(alias string) *grpc.ClientConn
}

var NodeInstance ISelector
