package main

import (
	"context"
	video "github.com/1359332949/douyin/kitex_gen/video/video/video"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// PublishAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionRequest) (resp *video.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserFeed implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserFeed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
