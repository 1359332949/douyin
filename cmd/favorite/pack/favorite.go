package pack

import (
	"github.com/1359332949/douyin/cmd/interact/dal/db"
	// "github.com/1359332949/douyin/kitex_gen/interact"
	"github.com/1359332949/douyin/kitex_gen/user"
	"github.com/1359332949/douyin/kitex_gen/video"
)	

// VideoList pack video list info
func VideoList(currentId int64, videoData []*video.Video, userMap map[int64]*user.User, favoriteMap map[int64]*db.Favorite, relationMap map[int64]*db.Relation) []*video.Video {
	videoList := make([]*video.Video, 0)
	for _, video := range videoData {
		videoUser, ok := userMap[video.AuthorID]
		if !ok {
			videoUser = &db.User{
				Name:          "未知用户",
				FollowCount:   0,
				FollowerCount: 0,
			}
			videoUser.Id = 0
		}

		var isFavorite bool = false
		var isFollow bool = false

		if currentId != -1 {
			_, ok := favoriteMap[int64(video.Id)]
			if ok {
				isFavorite = true
			}
			_, ok = relationMap[video.Author.Id]
			if ok {
				isFollow = true
			}
		}
		videoList = append(videoList, &video.Video{
			Id: int64(video.Id),
			Author: &user.User{
				Id:            int64(videoUser.Id),
				Name:          videoUser.Name,
				FollowCount:   videoUser.FollowCount,
				FollowerCount: videoUser.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}

	return videoList
}