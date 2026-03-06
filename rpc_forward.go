package rpcforw

import (
	"context"
	"fmt"

	"github.com/civet148/go-copier"
	"google.golang.org/grpc"
)

// Call 通用转发方法
func Call[TFromRequest, TFromReply, TRpcReq, TRpcReply any](
	ctx context.Context,
	req *TFromRequest,
	rpcHandler func(ctx context.Context, rpcReq *TRpcReq, opts ...grpc.CallOption) (*TRpcReply, error),
	callOptions ...grpc.CallOption,
) (*TFromReply, *TRpcReply, error) {

	// 创建RPC请求对象
	var rpcReq TRpcReq
	if err := copier.Copy(&rpcReq, req); err != nil {
		return nil, nil, fmt.Errorf("copy rpc request failed: %v", err)
	}

	// 调用RPC方法
	rpcReply, err := rpcHandler(ctx, &rpcReq, callOptions...)
	if err != nil {
		return nil, nil, err
	}

	// 创建并填充回复对象
	var reply TFromReply
	if err = copier.Copy(&reply, rpcReply); err != nil {
		return nil, nil, fmt.Errorf("copy rpc reply failed: %v", err)
	}
	return &reply, rpcReply, nil
}

// CallSame 通用转发方法(双方请求和响应结构是均为相同类型)
func CallSame[TRequest, TReply any](
	ctx context.Context,
	req *TRequest,
	rpcHandler func(ctx context.Context, rpcReq *TRequest, opts ...grpc.CallOption) (*TReply, error),
	callOptions ...grpc.CallOption,
) (*TReply, error) {
	// 调用RPC方法
	rpcReply, err := rpcHandler(ctx, req, callOptions...)
	if err != nil {
		return nil, err
	}
	return rpcReply, nil
}
