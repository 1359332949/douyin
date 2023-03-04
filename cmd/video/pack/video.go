
package pack

import (
	"github.com/1359332949/douyin/cmd/video/dal/db"
	"github.com/1359332949/douyin/kitex_gen/video"
	"github.com/1359332949/douyin/kitex_gen/user"
)

// User pack video info
func User(u *db.User) *video.User {
	if u == nil {
		return nil
	}

	return &video.User{
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

// Users pack list of video info
// func Users(us []*db.User) []*video.User {
// 	users := make([]*video.User, 0)
// 	for _, u := range us {
// 		if temp := User(u); temp != nil {
// 			users = append(users, temp)
// 		}
// 	}
// 	return users
// }
// Video pack of One Video info
func QueryVideo(v *db.Video) *video.Video {
	if v == nil {
		return nil
	}

	return &video.Video{
		Id: int64(v.ID), 
		// Author: int64(v.AuthorID), 
		PlayUrl: string(v.PlayUrl),
		CoverUrl: string(v.CoverUrl),
		FavoriteCount: int64(v.FavoriteCount),
		CommentCount: int64(v.CommentCount),
		// IsFavorite: bool(v.IsFavorite),
		Title: string(v.Title),
		}
}


// Videos pack list of person
func Videos(vs []*db.Video, author []*user.User) []*video.Video {
	videos := make([]*video.Video, 0)
	for index, v := range vs {
		temp := QueryVideo(v)
		if temp != nil {
			videos = append(videos, temp)
			videos[index].Author =  author
		}
		
	}
	return videos
}
// // Video pack of One Video info
// func Video(v *db.Video, author *db.User) *video.Video {
// 	if v == nil {
// 		return nil
// 	}

// 	return &video.Video{
// 		Id: int64(v.ID), 
// 		Author: User(author),
// 		PlayUrl: string(v.PlayUrl),
// 		CoverUrl: string(v.CoverUrl),
// 		FavoriteCount: int64(v.FavoriteCount),
// 		CommentCount: int64(v.CommentCount),
// 		// IsFavorite: bool(v.IsFavorite),
// 		Title: string(v.Title),
// 		}
// }

// // Videos pack list of person
// func Videos(vs []*db.Video, author *db.User) []*video.Video {
// 	videos := make([]*video.Video, 0)
// 	for _, v := range vs {
// 		if temp := Video(v, author); temp != nil {
// 			videos = append(videos, temp)
// 		}
// 	}
// 	return videos
// }

