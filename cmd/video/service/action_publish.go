
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "github.com/1359332949/douyin/cmd/video/pack"
	"github.com/1359332949/douyin/cmd/video/dal/db"
	"github.com/1359332949/douyin/kitex_gen/video"
	// "github.com/1359332949/douyin/pkg/errno"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new CreateUserService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

/// PublishActionService query video info
func (s *PublishActionService) PublishAction(req *video.PublishActionRequest) error{
	

	VideoModel := &db.Video{
		AuthorID:   req.UserId,
		PlayUrl:  req.FileUrl,
		CoverUrl: req.CoverUrl,
		Title: req.Title,
	}
	return db.CreateVideo(s.ctx, []*db.Video{VideoModel})
}