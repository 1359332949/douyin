package pack

import (
	"github.com/1359332949/douyin/cmd/favorite/dal/db"
	// "github.com/1359332949/douyin/kitex_gen/interact"
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/kitex_gen/video"
	// "github.com/1359332949/douyin/kitex_gen/relation"
)	

// VideoList pack video list info
func VideoList(currentId int64, videoData []*video.Video, userMap map[int64]*user.User, favoriteMap map[int64]*db.Favorite, follow_users []*user.User) []*video.Video {
	videoList := make([]*video.Video, 0)
	for _, v := range videoData {
		videoUser, ok := userMap[v.Author.Id]
		if !ok {
			videoUser = &user.User{
				Name:          "未知用户",
				FollowCount:   0,
				FollowerCount: 0,
			}
			videoUser.Id = 0
		}

		var isFavorite bool = false
		var isFollow bool = false

		if currentId != -1 {
			_, ok := favoriteMap[int64(v.Id)]
			if ok {
				isFavorite = true
			}
			for _, u := range follow_users{
				if v.Author.Id == u.Id{
					isFollow = true
				}

			}
			// _, ok = follow_users[video.Author.Id]
			// if ok {
			// 	isFollow = true
			// }
		}
		videoList = append(videoList, &video.Video{
			Id: int64(v.Id),
			Author: &user.User{
				Id:            int64(videoUser.Id),
				Name:          videoUser.Name,
				FollowCount:   videoUser.FollowCount,
				FollowerCount: videoUser.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isFavorite,
			Title:         v.Title,
		})
	}

	return videoList
}