package rpc

import (
	"context"
	"log"
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/kitex_gen/user/userservice"
	"github.com/1359332949/douyin/pkg/consts"
	"github.com/1359332949/douyin/pkg/errno"
	"github.com/1359332949/douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func InitUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.FavoriteServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		consts.UserServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FavoriteServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// MgetUser query list of user info by name
func MgetUser(ctx context.Context, req *user.MGetUserRequest) ([]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	return resp.Users, nil

}
func QueryUserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.User, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return resp.User, err
	}
	if resp.StatusCode != 0 {
		return resp.User, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	log.Println(resp.User)
	return resp.User, nil
}
