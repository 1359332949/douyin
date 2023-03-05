package pack

import (
	// "time"

	"github.com/1359332949/douyin/cmd/relation/dal/db"
	"github.com/1359332949/douyin/kitex_gen/relation"
)

// Message pack message info
func Message(u *db.Message) *relation.Message {
	if u == nil {
		return nil
	}
	return &relation.Message{
		Id:         int64(u.ID),
		FromUserId: u.FromUserId,
		ToUserId:   u.ToUserId,
		Content:    u.Content,
		CreateTime: u.CreateTime.Unix(),

	}
}

// Messages pack list of message info
func Messages(msgs []*db.Message) []*relation.Message {
	messages := make([]*relation.Message, 0)
	for _, msg := range msgs {
		if temp := Message(msg); temp != nil {
			messages = append(messages, temp)
		}
	}
	return messages
}

