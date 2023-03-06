
package rpc

import (
	"context"
	// "log"

	"github.com/1359332949/douyin/kitex_gen/message"
	"github.com/1359332949/douyin/kitex_gen/message/messageservice"
	"github.com/1359332949/douyin/pkg/consts"
	"github.com/1359332949/douyin/pkg/errno"
	"github.com/1359332949/douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func initMessage() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := messageservice.NewClient(
		consts.MessageServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

// 传递 关注操作 的上下文, 并获取 RPC Server 端的响应.
func MessageAction(ctx context.Context, req *message.MessageActionRequest) (error) {
	resp, err := messageClient.MessageAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

// 传递 获取正在关注列表操作 的上下文, 并获取 RPC Server 端的响应.
func MessageFollowList(ctx context.Context, req *message.MessageFollowListRequest) ([]*message.User,error) {
	resp, err := messageClient.MessageFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	// log.Println("***api-rpc-message.go***")
	// log.Println(resp.UserList)
	return resp.UserList, nil
}

// 传递 获取粉丝列表操作 的上下文, 并获取 RPC Server 端的响应.
func MessageFollowerList(ctx context.Context, req *message.MessageFollowerListRequest) ([]*message.User,error) {
	resp, err := messageClient.MessageFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserList, nil
}

// MessageFriendList
func MessageFriendList(ctx context.Context, req *message.MessageFriendListRequest) ([]*message.FriendUser, error) {
	resp, err := messageClient.MessageFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return 0, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserList, nil
}



func MessageChat(ctx context.Context, req *message.MessageChatRequest) ([]*message.Message,error) {
	resp, err := messageClient.MessageChat(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return nil,errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.Messages, nil
}


func MessageAction(ctx context.Context, req *message.MessageActionRequest) error {
	resp, err := messageClient.MessageAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		// return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}