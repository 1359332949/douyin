
package db

import (
	"context"
	"fmt"
	"github.com/1359332949/douyin/pkg/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	// Avatar string  `json:"avatar"`
	// BackgroundImage string  `json:"background_image"`
	// Signature string  `json:"signature"`
	// TotalFavorited string  `json:"total_favorited"`
	// WorkCount int64  `json:"work_count"`
	// FavoriteCount int64  `json:"favorite_count"`
}



func (u *User) TableName() string {
	return consts.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	fmt.Println("%s", users[0].Username)
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUserInfo query list of user info
func QueryUserInfo(ctx context.Context, uid int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", uid).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
// MgetUserInfo query list of user info
func MgetUserById(ctx context.Context, userIds []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIds) == 0 {
		return res, nil
	}
	
	
	if err := DB.WithContext(ctx).Where("id = ?", userIds).Find(&res).Error; err != nil {
			return nil, err
	}

	
	return res, nil
}
