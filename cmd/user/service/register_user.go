package service
import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/1359332949/douyin/cmd/user/dal/db"
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/pkg/errno"
)

type RegisterUserService struct {
	ctx context.Context
}

// NewRegisterUserService new RegisterUserService
func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *RegisterUserService) RegisterUser(req *user.RegisterUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: password,
		IsFollow: true,
	}})
}


