
package db

import (
	"context"
	// "fmt"
	"time"
	"log"
	"github.com/1359332949/douyin/pkg/consts"
	// "github.com/1359332949/douyin/cmd/user/dal/db"
	"gorm.io/gorm"
)


type Message struct {
	gorm.Model
	ToUserId   int64  `gorm:"type:varchar(32);not null" json:"to_user_id"`
	FromUserId int64  `gorm:"type:varchar(32);not null" json:"from_user_id"`
	Content    string `gorm:"type:varchar(256);not null" json:"content"`
	CreateTime   time.Time   `gorm:"column:create_time;default:null " json:"create_time"`
	
	
}


func (u *Message) TableName() string {
	return consts.MessageTableName

}

// CreateMessage create message info
func CreateMessage(ctx context.Context, message *Message) error {
	log.Println(message)
	if err := DB.WithContext(ctx).Create(message).Error; err != nil {
		log.Println(err)
		return err
	}
	log.Println("++++++++++++++++++++++++++++++",message)
	return nil
}

// MGetMessages multiple get list of message info
func MGetMessages(ctx context.Context, uid int64, toUId int64) ([]*Message, error) {
	res := make([]*Message, 0)
	
	if err := DB.WithContext(ctx).Model(&Message{}).Where("from_user_id = ? AND to_user_id = ? Or from_user_id = ? AND to_user_id = ?", uid, toUId, toUId, uid).Order("id desc").Scan(&res).Error; err != nil{
		return nil, err
	}
	
	return res, nil
}


