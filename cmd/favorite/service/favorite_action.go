package service

import (
	"context"
	"errors"
	"github.com/1359332949/douyin/cmd/favorite/dal/db"
	"github.com/1359332949/douyin/cmd/favorite/rpc"
	"github.com/1359332949/douyin/kitex_gen/favorite"
	"github.com/1359332949/douyin/pkg/consts"
	"github.com/1359332949/douyin/kitex_gen/video"
	// "github.com/1359332949/douyin/pkg/jwt"
	"log"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction implement the mainke and unmainke operations
func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {

	log.Println("------------------------",req.VideoId,"-----------------------")
	// videos, err := rpc.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	videos, err := rpc.QueryVideoByVideoIds(s.ctx, &video.QueryVideoByVideoIdsRequest{VideoIds: []int64{req.VideoId}})
	log.Println("------------------------",videos[0],"-----------------------")
	if err != nil {
		return err
	}
	if len(videos) == 0 {
		return errors.New("video not exist")
	}

	//若ActionType（操作类型）等于1，则向favorite表创建一条记录，同时向video表的目标video增加点赞数
	//若ActionType等于2，则向favorite表删除一条记录，同时向video表的目标video减少点赞数
	//若ActionType不等于1和2，则返回错误
	if req.ActionType == consts.Like {
		favorite := &db.Favorite{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		}

		err := db.CreateFavorite(s.ctx, favorite, req.VideoId)
		if err != nil {
			return err
		}
	}
	if req.ActionType == consts.Unlike {
		err := db.DeleteFavorite(s.ctx, req.UserId, req.VideoId)
		if err != nil {
			return err
		}

	}
	if req.ActionType != consts.Like && req.ActionType != consts.Unlike {
		return errors.New("action type no equal 1 and 2")
	}
	return nil
}