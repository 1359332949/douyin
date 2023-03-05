package service

import (
	"context"
	// "errors"
	"github.com/1359332949/douyin/cmd/comment/dal/db"
	"github.com/1359332949/douyin/kitex_gen/comment"
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/cmd/comment/pack"
	"github.com/1359332949/douyin/cmd/comment/rpc"
	// "github.com/1359332949/douyin/pkg/consts"
	// "github.com/1359332949/douyin/pkg/errno"
	// "github.com/1359332949/douyin/pkg/jwt"
	// "log"
	"errors"
	"time"
)


type CommentActionService struct {
	ctx context.Context
}

// NewCommentActionService new CommentActionService
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// CommentAction implement the mainke and unmainke operations
func (s *CommentActionService) CommentAction(req *comment.CommentActionRequest) (*comment.Comment, error) {

	// log.Info("get comment action req", *req)
	// resp = new(comment.CommentActionResponse)

	//TODO check video id
	if req.ActionType == 1 {
		cmt := &db.Comment{UserId: int64(req.UserId), Content: req.CommentText,
			VideoId: int64(req.VideoId), IsValid: true, CreateDate: time.Now().Format("06-01")}
		if err := db.CreateComment(s.ctx, cmt); err != nil {
			return nil, err
		}
		u, err := rpc.QueryUserInfo(s.ctx, &user.UserInfoRequest{UserId: cmt.UserId})
		
		if err != nil {
			return nil, err
		}
		// u := user[0]
		
		comment := pack.Comment(cmt, u)
		return comment, nil
		// return &comment.Comment{Id: int64(cmt.ID), User: user, Content: cmt.Content, CreateDate: cmt.CreateDate}, nil
	} else if req.ActionType == 2 {
		cmt := &db.Comment{ID: int64(req.CommentId)}
		tmp, err := db.SelectComment(s.ctx, int64(cmt.ID))
		if err != nil {
			return nil, err
		}
		if tmp == nil {
			return nil, err
		}
		if err := db.DeleteComment(s.ctx, cmt); err != nil {
			return nil, err
		}
		u, err := rpc.QueryUserInfo(s.ctx, &user.UserInfoRequest{UserId: tmp.UserId})
		if err != nil {
			return nil, err
		}
		// u := user[0]
		
		
		comment := pack.Comment(tmp, u)
		return comment, nil
		// return &comment.CommentActionResponse{StatusCode: errno.SuccessCode,
		// 	Comment: &comment.Comment{Id: int64(tmp.ID), User: user, Content: tmp.Content, CreateDate: tmp.CreateDate}}, nil
	} else {
		// msg := "err"
		err := errors.New("ActionTypeErrCode")
		return nil, err
		// return &comment.CommentActionResponse{StatusCode: errno.ActionTypeErrCode, StatusMsg: msg}, nil
	}
}
