
package db

import (
	"context"
	// "fmt"
	"time"
	"github.com/1359332949/douyin/pkg/consts"
	"gorm.io/gorm"
	"github.com/cloudwego/kitex/pkg/klog"
)

type Video struct {
	gorm.Model
	ID       int64   `gorm:"column:id;primary_key;AUTO_INCREMENT"`   
	AuthorID      int64     `gorm:"column:author_id;NOT NULL"`
	// PublishTime   time.Time `gorm:"column:publish_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	PlayUrl      string    `gorm:"column:play_url;NOT NULL"`
	CoverUrl     string    `gorm:"column:cover_url;NOT NULL"`
	FavoriteCount int64     `gorm:"column:favorite_count;default:0"`
	CommentCount  int64     `gorm:"column:comment_count;default:0"`
	Title         string    `gorm:"column:title;NOT NULL"`
	// IsFavorite bool  `gorm:"column:is_favorite;default:0"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;default:null " json:"updated_at"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}


// 返回某一用户的视频表
func MGetVideosOfUserIDList(ctx context.Context, userId int64) ([]*Video, error) {
	// 获取视频列表
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Model(&Video{}).Where("author_id = ?", userId).Order("id desc").Scan(&res).Error; err != nil{
		return nil, err
	}

	// 返回
	return res, nil
}

func CreateVideo(ctx context.Context, videos []*Video) error {
	if err := DB.WithContext(ctx).Create(videos).Error; err != nil {
		return err
	}
	return nil
}

// GetFeedVideos  get list of videos info
func MGetFeedVideos(ctx context.Context, limit int, latestTime int64) ([]*Video, error) {
	videos := make([]*Video, 0)

	if latestTime == 0 {
		cur_time := int64(time.Now().UnixMilli())
		latestTime = cur_time
	}
	conn := DB.WithContext(ctx)

	if err := conn.Limit(limit).Order("updated_at desc").Find(&videos, "updated_at < ?", time.UnixMilli(latestTime)).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
// QueryVideoByVideoIds query video info by video ids 返回对应ids的视频s
func QueryVideoByVideoIds(ctx context.Context, videoId []int64) ([]*Video, error) {
	var videos []*Video
	err := DB.WithContext(ctx).Where("id in (?)", videoId).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByVideoIds error " + err.Error())
		return nil, err
	}
	return videos, nil
}