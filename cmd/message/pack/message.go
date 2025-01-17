package pack

import (
	// "time"

	"github.com/1359332949/douyin/cmd/message/dal/db"
	"github.com/1359332949/douyin/kitex_gen/message"
)

// Message pack message info
func Message(u *db.Message) *message.Message {
	if u == nil {
		return nil
	}
	return &message.Message{
		Id:         int64(u.ID),
		FromUserId: u.FromUserId,
		ToUserId:   u.ToUserId,
		Content:    u.Content,
		CreateTime: u.CreateTime.Unix(),

	}
}

// Messages pack list of message info
func Messages(msgs []*db.Message) []*message.Message {
	messages := make([]*message.Message, 0)
	for _, msg := range msgs {
		if temp := Message(msg); temp != nil {
			messages = append(messages, temp)
		}
	}
	return messages
}

// func FormatTime(t time.Time) *string {
// 	s := t.Format("2006-01-02 15:04:05")
// 	return &s
// }