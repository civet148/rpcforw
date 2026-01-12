package rpcforw

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
)

type FromRequest struct {
	Name string
}

type FromReply struct {
	Id uint64
}

type RpcRequest struct {
	Name       string
	UpdateUser string
}

type RpcReply struct {
	Id uint64
}

func MockWebFunc(ctx context.Context, req *FromRequest) (reply *FromReply, err error) {
	reply, _, err = Call[
		FromRequest,
		FromReply,
		RpcRequest,
		RpcReply,
	](ctx, req, MockRpcFunc)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func MockRpcFunc(ctx context.Context, req *RpcRequest, opts ...grpc.CallOption) (reply *RpcReply, err error) {
	return &RpcReply{Id: 10000}, nil
}

func TestMockWebFunc(t *testing.T) {
	ctx := context.Background()
	reply, err := MockWebFunc(ctx, &FromRequest{Name: "test"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("reply [%+v] \n", reply)
}
