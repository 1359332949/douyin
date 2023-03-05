package rpc

import (
	"context"
	"log"
	"github.com/1359332949/douyin/kitex_gen/video"
	"github.com/1359332949/douyin/kitex_gen/video/videoservice"
	"github.com/1359332949/douyin/pkg/consts"
	"github.com/1359332949/douyin/pkg/errno"
	"github.com/1359332949/douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func Init() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.FavoriteServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := videoservice.NewClient(
		consts.VideoServiceName, // DestService
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
	videoClient = c
}


func Info(ctx context.Context, req *video.UserInfoRequest) (*video.User, error) {
	resp, err := videoClient.UserInfo(ctx, req)
	if err != nil {
		return resp.User, err
	}
	if resp.StatusCode != 0 {
		return resp.User, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	log.Println(resp.User)
	return resp.User, nil
}
