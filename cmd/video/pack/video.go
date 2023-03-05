
package pack

import (
	"github.com/1359332949/douyin/cmd/video/dal/db"
	"github.com/1359332949/douyin/kitex_gen/video"
	"github.com/1359332949/douyin/kitex_gen/user"
)

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
			videos[index].Author =  author[index]
		}
		
	}
	return videos
}


// Videos pack list of person
func VideosByOne(vs []*db.Video, author *user.User) []*video.Video {
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

