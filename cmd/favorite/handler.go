package main

import (
	"context"
	favorite "github.com/1359332949/douyin/kitex_gen/favorite"
	
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/favorite/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/favorite/service"
	// "time"

	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
	// "log"

)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.FavoriteActionResponse)

	if len(req.Token) == 0 || req.VideoId == 0 || req.ActionType == 0 {
		resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	return resp, nil
	
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(favorite.FavoriteListResponse)

	if req.UserId == 0 {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	videoList, err := service.NewFavoriteListService(ctx).FavoriteList(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	resp = pack.BuildFavoriteListBaseResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}
