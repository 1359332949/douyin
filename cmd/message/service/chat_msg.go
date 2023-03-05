package service

import (
	"context"

	"github.com/1359332949/douyin/cmd/message/dal/db"
	"github.com/1359332949/douyin/cmd/message/pack"
	"github.com/1359332949/douyin/kitex_gen/message"
)

type ChatMsgService struct {
	ctx context.Context
}

func NewChatMsgService(ctx context.Context) *ChatMsgService {
	return &ChatMsgService{ctx: ctx}
}

func (s *ChatMsgService) MGetChatMsg(req *message.MessageChatRequest) ([]*message.Message, error) {
	messageModels, err := db.MGetMessages(s.ctx, req.FromUserId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	
	messages := pack.Messages(messageModels)
	
	return messages, nil
	
}