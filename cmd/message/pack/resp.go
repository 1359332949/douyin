
package pack

import (
	"errors"
	// "time"

	"github.com/1359332949/douyin/kitex_gen/message"
	"github.com/1359332949/douyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *message.BaseResp {
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

func baseResp(err errno.ErrNo) *message.BaseResp {
	return &message.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
