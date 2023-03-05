package service

import (
	"context"
	"time"
	"log"
	// "crypto/md5"
	// "fmt"
	// "io"

	"github.com/1359332949/douyin/cmd/video/pack"
	"github.com/1359332949/douyin/cmd/video/dal/db"
	"github.com/1359332949/douyin/cmd/video/rpc"
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/kitex_gen/video"
	// "github.com/1359332949/douyin/pkg/errno"
)

const (
	LIMIT = 30 // 单次返回最大视频数
)

type GetVideoFeedService struct {
	ctx context.Context
}

// NewGetUserByIdService new GetUserByIdService
func NewGetVideoFeedService(ctx context.Context) *GetVideoFeedService {
	return &GetVideoFeedService{ctx: ctx}
}

// get video info.
func (s *GetVideoFeedService) GetVideoFeed(req *video.FeedRequest) (vis []*video.Video, nextTime int64, err error) {
	log.Println("----------------------kitex feed--------------------------------------")
	videos, err := db.MGetVideos(s.ctx, LIMIT, req.LatestTime)
	log.Println("-------------req.LatestTime----------")
	log.Println(req.LatestTime)
	log.Println(videos[0])
	if err != nil {
		return vis, nextTime, err
	}

	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
		return vis, nextTime, nil
	} else {
		nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	}
	log.Println("-------------req.nextTime----------")
	log.Println(nextTime)

	//查询视频作者信息
	nextTime = time.Now().UnixMilli()
	pack_videos := make([]*video.Video, 0) 
	user_ids := make([]int64, 0) 
	for _, video := range videos{
		user_ids = append(user_ids, video.AuthorID)
		
	}
	
	users, err := rpc.MgetUser(s.ctx, &user.MGetUserRequest{UserIds: user_ids})
	if err != nil{
		return nil, 0, err
	}

	pack_videos = pack.Videos(videos, users)
	return pack_videos, nextTime, nil
}