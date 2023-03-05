package main

import (
	"context"
	comment "github.com/1359332949/douyin/kitex_gen"
	"github.com/1359332949/douyin/cmd/comment/pack"
	"github.com/1359332949/douyin/cmd/comment/service"
	// "time"

	"github.com/1359332949/douyin/pkg/errno"
	"log"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.CommentActionResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}


	comment, err := service.NewCommentActionService(ctx).CommentAction(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	resp.Comment = comment
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	resp = new(comment.CommentListResponse)

	if req.UserId == 0 {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	log.Println("------------------------------kitex--commentlist-----------------------------------------------------")
	comment_list, err := service.NewCommentListService(ctx).CommentList(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg

	resp.CommentList = comment_list
	return resp, nil
}
