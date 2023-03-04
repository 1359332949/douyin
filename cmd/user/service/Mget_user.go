
package service

import (
	"context"
	// "crypto/md5"
	// "fmt"
	// "io"
	// "log"
	"github.com/1359332949/douyin/cmd/user/pack"
	"github.com/1359332949/douyin/cmd/user/dal/db"
	"github.com/1359332949/douyin/kitex_gen/user"
	// "github.com/1359332949/douyin/pkg/errno"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	modelUsers, err := db.MgetUserById(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}