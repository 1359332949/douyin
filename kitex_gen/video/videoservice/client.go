// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	video "github.com/1359332949/douyin/kitex_gen/video"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PublishAction(ctx context.Context, req *video.PublishActionRequest, callOptions ...callopt.Option) (r *video.PublishActionResponse, err error)
	PublishList(ctx context.Context, req *video.PublishListRequest, callOptions ...callopt.Option) (r *video.PublishListResponse, err error)
	GetVideoFeed(ctx context.Context, req *video.FeedRequest, callOptions ...callopt.Option) (r *video.FeedResponse, err error)
	QueryVideoByVideoIds(ctx context.Context, req *video.QueryVideoByVideoIdsRequest, callOptions ...callopt.Option) (r *video.QueryVideoByVideoIdsResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) PublishAction(ctx context.Context, req *video.PublishActionRequest, callOptions ...callopt.Option) (r *video.PublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, req)
}

func (p *kVideoServiceClient) PublishList(ctx context.Context, req *video.PublishListRequest, callOptions ...callopt.Option) (r *video.PublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, req)
}

func (p *kVideoServiceClient) GetVideoFeed(ctx context.Context, req *video.FeedRequest, callOptions ...callopt.Option) (r *video.FeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoFeed(ctx, req)
}

func (p *kVideoServiceClient) QueryVideoByVideoIds(ctx context.Context, req *video.QueryVideoByVideoIdsRequest, callOptions ...callopt.Option) (r *video.QueryVideoByVideoIdsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryVideoByVideoIds(ctx, req)
}
