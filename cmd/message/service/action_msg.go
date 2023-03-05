package service

import (
	"context"
	"log"
	// "time"
	"github.com/1359332949/douyin/cmd/message/dal/db"
	// "github.com/1359332949/douyin/cmd/message/pack"
	"github.com/1359332949/douyin/kitex_gen/message"
)
type ActionMsgService struct {
	ctx context.Context
}

func NewActionMsgService(ctx context.Context) *ActionMsgService {
	return &ActionMsgService{ctx: ctx}
}

// Create Message
func (s *ActionMsgService) ActionMsg(req *message.MessageActionRequest) error {
	MessageModel := &db.Message{
		ToUserId:   req.ToUserId,
		FromUserId:  req.FromUserId,
		Content: req.Content,

	}
	
	
	log.Println(req.FromUserId, "------------------------", req.ToUserId)
	return db.CreateMessage(s.ctx, MessageModel)
}