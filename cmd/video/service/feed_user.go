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
	"github.com/1359332949/douyin/kitex_gen/video"
	// "github.com/1359332949/douyin/pkg/errno"
)

const (
	LIMIT = 30 // 单次返回最大视频数
)

type GetUserFeedService struct {
	ctx context.Context
}

// NewGetUserByIdService new GetUserByIdService
func NewGetUserFeedService(ctx context.Context) *GetUserFeedService {
	return &GetUserFeedService{ctx: ctx}
}

// get video info.
func (s *GetUserFeedService) GetVideoFeed(req *video.FeedRequest) (vis []*video.Video, nextTime int64, err error) {
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
	for index, val := range videos{
		users, err := rpc.MgetUser(s.ctx, videos[index].AuthorID)
		u := users[0]
		if err != nil{
			return nil, 0, err
		}
		
		if temp := pack.Video(val, u); temp != nil{
			pack_videos = append(pack_videos, temp)

		}
		

	}

	return pack_videos, nextTime, nil
}