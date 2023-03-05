
package pack

import (
	"errors"
	// "time"

	"github.com/1359332949/douyin/kitex_gen/comment"
	"github.com/1359332949/douyin/pkg/errno"
)
// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *comment.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *comment.BaseResp {
	return &comment.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
// BuildFavoriteBaseResp build comment baseResp from error
func BuildFavoriteBaseResp(err error) *comment.FavoriteActionResponse {
	if err == nil {
		return commentbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return commentbaseResp(s)
}

func commentbaseResp(err errno.ErrNo) *comment.FavoriteActionResponse {
	return &comment.FavoriteActionResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

func BuildFavoriteListBaseResp(err error) *comment.FavoriteListResponse {
	if err == nil {
		return commentListbaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return commentListbaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return commentListbaseResp(s)
}

func commentListbaseResp(err errno.ErrNo) *comment.FavoriteListResponse {
	return &comment.FavoriteListResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}

func Err1(err error) *comment.CommentActionResponse {
	msg := err.Error()
	return &comment.CommentActionResponse{StatusCode: errno.CommentError, StatusMsg: msg}
}

func Err2(err error) *comment.CommentListResponse {
	msg := err.Error()
	return &comment.CommentListResponse{StatusCode: errno.SuccessCode, StatusMsg: msg,
		CommentList: []*comment.Comment{}}
}