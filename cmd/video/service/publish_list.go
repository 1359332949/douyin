
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "log"
	"github.com/1359332949/douyin/cmd/video/rpc"
	"github.com/1359332949/douyin/cmd/video/dal/db"
	"github.com/1359332949/douyin/kitex_gen/video"
	"github.com/1359332949/douyin/cmd/video/pack"
	// "github.com/1359332949/douyin/main/pkg/errno"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new CreateUserService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

/// PublishListService query video info
func (s *PublishListService) PublishList(req *video.PublishListRequest) ([]*video.Video, error){
	videoModels, err := db.MGetVideosOfUserIDList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	users, err := rpc.MgetUser(s.ctx, req.UserId)
	u := users[0]
	videos := pack.Videos(videoModels, u)
	// log.Println(videos[0].PlayUrl)
	return videos, nil
}