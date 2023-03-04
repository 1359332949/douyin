
package pack

import (
	"errors"
	// "time"
	// "log"
	"github.com/1359332949/douyin/kitex_gen/video"
	"github.com/1359332949/douyin/pkg/errno"
)
// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *video.BaseResp {
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

func baseResp(err errno.ErrNo) *video.BaseResp {
	
	return &user.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
