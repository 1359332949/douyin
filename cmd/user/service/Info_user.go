
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	"log"
	"github.com/1359332949/douyin/cmd/user/pack"
	"github.com/1359332949/douyin/cmd/user/dal/db"
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/pkg/errno"
)

type UserInfoService struct {
	ctx context.Context
}

// NewUserInfoService new CreateUserService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

/// UserInfoService query user info
func (s *UserInfoService) UserInfo(req *user.UserInfoRequest) (*user.User, error) {
	
	
	user_id := req.UserId
	
	users, err := db.QueryUserInfo(s.ctx, user_id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.AuthorizationFailedErr
	}
	u := users[0]
	user_info := pack.User(u)
	log.Println("***------------------------------------kitex-test---------------------------------------***")
	// log.Println(user_info)
	return user_info, nil
}