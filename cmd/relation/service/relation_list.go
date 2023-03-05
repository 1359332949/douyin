
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "log"
	"github.com/1359332949/douyin/cmd/relation/dal/db"
	"github.com/1359332949/douyin/cmd/relation/rpc"
	"github.com/1359332949/douyin/kitex_gen/relation"
	"github.com/1359332949/douyin/kitex_gen/user"
	// "github.com/1359332949/douyin/pkg/errno"
)

type RelationListService struct {
	ctx context.Context
}

// NewRelationListService new RelationListService
func NewRelationListService(ctx context.Context) *RelationListService {
	return &RelationListService{ctx: ctx}
}

// 查找关注列表
func (s *RelationListService) RelationFollowList(req *relation.RelationFollowListRequest) ([]*user.User, error) {
	relation_list, err := db.RelationFollowList(s.ctx, req.UserId)
	if err != nil{
		return nil,err
	}

	userIDs :=make([]int64,0)
	for _, u := range relation_list{
		userIDs= append(userIDs,int64(u.ToUserID))
	}
	users, err := rpc.MGetUsers(ctx,userIDs)
	if err != nil {
		return nil, err
	}
	// log.Println(users)
	// return BuildUsers(s.ctx,uid,users)
	return users, nil
}

// 查找粉丝列表 
func (s *RelationListService) RelationFollowerList(req *relation.RelationFollowerListRequest)  ([]*user.User, error) {
	users, err := db.RelationFollowerList(s.ctx, req.UserId)
	if err != nil{
		return nil,err
	}

	userIDs :=make([]int64,0)
	for _, u := range relation_list{
		userIDs= append(userIDs,int64(u.FromUserID))
	}
	users, err := rpc.MGetUsers(ctx,userIDs)
	if err != nil {
		return nil, err
	}
	// log.Println(users)
	// return BuildUsers(s.ctx,uid,users)
	return users, nil
}

// 查找好友列表  💦先用粉丝列表代替，返回为user包装得到的FriendUser
func (s *RelationListService) RelationFriendList(req *relation.RelationFriendListRequest)  ([]*relation.FriendUser, error) {
	LRelationList, RRelationList, err := db.RelationFriendList(s.ctx, req.UserId)
	if err != nil{
		return nil, err
	}
	LuserIDs :=make([]int64,0)
	for _,u := range LRelationList{
		LuserIDs= append(LuserIDs,int64(u.ToUserID))
		log.Println(LuserIDs)
	}

	RuserIDs :=make([]int64,0)
	for _,u := range RRelationList{
		RuserIDs= append(RuserIDs,int64(u.FromUserID))
		log.Println(RuserIDs)
	}
	userIDs :=make([]int64,0)

	m := make(map[int64]int)
	for _,v :=range LuserIDs{
		m[v]++
	}
	for _,v :=range RuserIDs{
		if m[v]==1{
			userIDs = append(userIDs,v)
		}
	}
	log.Println(userIDs)
	users, err := rpc.MGetUsers(ctx,userIDs)
	if err != nil {
		return nil, err
	}
	// return BuildFriendUsers(ctx,id,users)
	return users, nil
}
