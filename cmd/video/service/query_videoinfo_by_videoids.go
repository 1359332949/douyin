
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
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/cmd/video/pack"
	// "github.com/1359332949/douyin/main/pkg/errno"
)

type QueryVideoByVideoIdsService struct {
	ctx context.Context
}

// NewQueryVideoByVideoIdsService new CreateUserService
func NewQueryVideoByVideoIdsService(ctx context.Context) *QueryVideoByVideoIdsService {
	return &QueryVideoByVideoIdsService{ctx: ctx}
}

/// QueryVideoByVideoIdsService query video info
func (s *QueryVideoByVideoIdsService) QueryVideoByVideoIds(req *video.QueryVideoByVideoIdsRequest) ([]*video.Video, error){
	videoModels, err := db.QueryVideoByVideoIds(s.ctx, req.VideoIds)
	if err != nil {
		return nil, err
	}
	pack_videos := make([]*video.Video, 0) 
	user_ids := make([]int64, 0) 
	for _, video := range videoModels{
		user_ids = append(user_ids, video.AuthorID)
		
	}
	
	users, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserIds: user_ids})
	if err != nil{
		return nil, err
	}

	pack_videos = pack.Videos(videoModels, users)
	return pack_videos, nil
}