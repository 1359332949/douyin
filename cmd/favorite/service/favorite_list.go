package service

import (
	"context"
	"errors"
	"github.com/1359332949/douyin/cmd/favorite/dal/db"
	"github.com/1359332949/douyin/cmd/favorite/pack"
	"github.com/1359332949/douyin/cmd/favorite/rpc"
	"github.com/1359332949/douyin/kitex_gen/favorite"
	"github.com/1359332949/douyin/kitex_gen/relation"
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/kitex_gen/video"
	// "github.com/1359332949/douyin/pkg/consts"
	// "github.com/1359332949/douyin/pkg/jwt"
	"sync"
	"log"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

// FavoriteList get video information that users mainke
func (s *FavoriteListService) FavoriteList(req *favorite.FavoriteListRequest) ([]*video.Video, error) {

	log.Println("1===============================",req.UserId,"==================================")

	
	u, err := rpc.QueryUserInfo(s.ctx, &user.UserInfoRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("user not exist")
	}

	//获取目标用户的点赞视频id号
	videoIds, err := db.QueryFavoriteById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	log.Println("2===============================",videoIds[0],"==================================")
	//获取点赞视频的信息
	videoData, err := rpc.QueryVideoByVideoIds(s.ctx, &video.QueryVideoByVideoIdsRequest{VideoIds: videoIds})
	if err != nil {
		return nil, err
	}
	log.Println("3===============================",videoData,"==================================")
	//获取点赞视频的用户id号
	user_ids := make([]int64, 0)
	for _, video := range videoData {
		user_ids = append(user_ids, video.Author.Id)
	}

	//获取点赞视频的用户信息
	users, err := rpc.MgetUser(s.ctx, &user.MGetUserRequest{UserIds: user_ids})
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*user.User)
	for _, user := range users {
		userMap[int64(user.Id)] = user
	}

	var favoriteMap map[int64]*db.Favorite
	var follow_users []*user.User
	// var relationMap map[int64]*relation.Relation
	//if user not logged in
	if req.UserId == -1 {
		favoriteMap = nil
		follow_users = nil
	} else {
		var wg sync.WaitGroup
		wg.Add(2)
		var favoriteErr, relationErr error
		//获取点赞信息
		go func() {
			defer wg.Done()
			favoriteMap, err = db.QueryFavoriteByIds(s.ctx, req.UserId, videoIds)
			if err != nil {
				favoriteErr = err
				return
			}
		}()
		//获取关注信息
		go func() {
			defer wg.Done()
			follow_users, err = rpc.RelationFollowList(s.ctx, &relation.RelationFollowListRequest{UserId: req.UserId})
			if err != nil {
				relationErr = err
				return
			}
		}()
		wg.Wait()
		if favoriteErr != nil {
			return nil, favoriteErr
		}
		if relationErr != nil {
			return nil, relationErr
		}

	}
	// log.Println("4===============================",userIds,"==================================")
	videoList := pack.VideoList(req.UserId, videoData, userMap, favoriteMap, follow_users)
	return videoList, nil

}