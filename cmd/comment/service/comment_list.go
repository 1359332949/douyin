package service

import (
	"context"
	// "errors"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/comment/dal/db"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/comment/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/comment"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/pkg/consts"
	// "github.com/YANGJUNYAN0715/douyin/tree/main/pkg/jwt"
	// "sync"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/comment/rpc"
	"log"
	"fmt"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentListService new CommentListService
func NewCommentListService(ctx context.Context) *CommentListService {
	log.Println("=======================CommentListService======service====================================")
	return &CommentListService{ctx: ctx}
}

// CommentList get video information that users mainke
func (s *CommentListService) CommentList(req *comment.CommentListRequest) ([]*comment.Comment, error) {
	// log.QueryUserInfo("get comment list req", *req)
	// TODO: Your code here...
	log.Println("----------------videoid--",req.VideoId,"-----------------------")
	cmts, err := db.QueryComments(s.ctx, int64(req.VideoId))
	if err != nil {
		return nil, err
	}
	log.Println("------------------",cmts[0],"-----------------------")
	res := []*comment.Comment{}
	for _, c := range cmts {
		fmt.Println("get user , id = ", c.UserId)
		user, err := rpc.QueryUserInfo(s.ctx, c.UserId)
		if err != nil {
			return nil, err
		}
		u := user[0]
		// tmp := &comment.Comment{Id: int64(c.ID), Content: c.Content,
		// 	CreateDate: c.CreateDate, User: user}
		tmp :=	pack.Comment(c, u)
		res = append(res, tmp)
	}
	return res, nil
	// return &comment.CommentListResponse{StatusCode: errno.SuccessCode, CommentList: res}, nil



}