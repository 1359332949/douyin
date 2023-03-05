
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "log"
	"github.com/1359332949/douyin/cmd/relation/dal/db"
	"github.com/1359332949/douyin/kitex_gen/relation"
	// "github.com/1359332949/douyin/kitex_gen/user"
	// "github.com/1359332949/douyin/pkg/errno"
)

type IsFollowService struct {
	ctx context.Context
}

// NewIsFollowService new IsFollowService
func NewIsFollowService(ctx context.Context) *IsFollowService {
	return &IsFollowService{ctx: ctx}
}

// 查找关注列表
func (s *IsFollowService) IsFollow(req *relation.IsFollowRequest) (bool, error) {
	relation, err := db.GetRelation(s.ctx, req.UserId, req.ToUserId)
	is_follow := false
	if err != nil{
		return is_follow, err
	}
	if relation != nil{
		is_follow = true
	}
	// log.Println("relation-service")
	// log.Println(users)
	return is_follow,nil
}

