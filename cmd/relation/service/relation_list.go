
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "log"
	"github.com/1359332949/douyin/cmd/relation/dal/db"
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
	users, err := db.RelationFollowList(s.ctx, req.UserId)
	if err != nil{
		return nil,err
	}
	// log.Println("relation-service")
	// log.Println(users)
	return users,nil
}

// 查找粉丝列表 
func (s *RelationListService) RelationFollowerList(req *relation.RelationFollowerListRequest)  ([]*user.User, error) {
	users, err := db.RelationFollowerList(s.ctx, req.UserId)
	if err != nil{
		return nil,err
	}
	return users,nil
}

// 查找好友列表  💦先用粉丝列表代替，返回为user包装得到的FriendUser
func (s *RelationListService) RelationFriendList(req *relation.RelationFriendListRequest)  ([]*relation.FriendUser, error) {
	users, err := db.RelationFriendList(s.ctx, req.UserId)
	if err != nil{
		return nil,err
	}
	return users,nil
}
