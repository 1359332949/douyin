
package pack

import (
	"github.com/1359332949/douyin/cmd/user/dal/db"
	"github.com/1359332949/douyin/kitex_gen/user"
	
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{
		Id: int64(u.ID), 
		Name: u.Username,
		FollowCount: int64(u.FollowCount),
		FollowerCount: int64(u.FollowerCount),
		IsFollow: bool(u.IsFollow),
		// Avatar: string(u.Avatar),
		// BackgroundImage: string(u.BackgroundImage),
		// Signature: string(u.Signature),
		// TotalFavorited: string(u.TotalFavorited),
		// WorkCount: int64(u.WorkCount),
		// FavoriteCount: int64(u.FavoriteCount),
		}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}


