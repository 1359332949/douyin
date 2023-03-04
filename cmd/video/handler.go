package main

import (
	"context"
	video "github.com/1359332949/douyin/kitex_gen/video"


	// "fmt"
	// "log"
	"github.com/1359332949/douyin/cmd/video/pack"
	"github.com/1359332949/douyin/cmd/video/service"
	"github.com/1359332949/douyin/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.


// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(user.PublishActionResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	err = service.NewPublishActionService(ctx).PublishAction(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}

	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg

	return resp, nil
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(user.PublishListResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	videos_list, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}

	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	resp.VideoList = videos_list

	return resp, nil
}

// GetUserFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetUserFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(user.FeedResponse)

	vis, nextTime, err := service.NewGetUserFeedService(ctx).GetUserFeed(req)

	log.Println(vis[0])
	log.Println(nextTime)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}

	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg

	resp.VideoList = vis
	resp.NextTime = nextTime
	return resp, nil
}
