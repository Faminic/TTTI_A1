package main

import (
	"context"
	"strings"
	"time"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	timestamp := time.Now().Unix()

	//adjusting order of chat based on sender -> sender will always be first, receiver will always be second
	//s1, s2 := strings.Split(req.Message.GetChat(), ":")[0], strings.Split(req.Message.GetChat(), ":")[1]
	//if req.Message.GetSender() == s2 {
	//	req.Message.Chat = s2 + ":" + s1
	//}

	key := strings.ToLower(req.Message.GetChat())

	message := &Message{
		Message:   req.Message.GetText(),
		Sender:    req.Message.GetSender(),
		Timestamp: timestamp,
	}

	err := redisClient.SaveMessage(ctx, key, message)
	if err != nil {
		return nil, err
	}

	resp := rpc.NewSendResponse()
	resp.Code, resp.Msg = 0, "success"
	return resp, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	key := strings.ToLower(req.GetChat())

	limit := int64(req.GetLimit())
	if limit <= 0 {
		limit = 10
	}
	start := req.GetCursor()
	end := start + limit
	reverse := req.GetReverse()

	messages, err := redisClient.GetMessages(ctx, key, start, end, reverse)
	if err != nil {
		return nil, err
	}

	respMessages := make([]*rpc.Message, 0)
	var counter int64 = 0
	var nextCursor int64 = 0
	hasMore := false
	for _, msg := range messages {
		if counter+1 > limit {
			hasMore = true
			nextCursor = end
			break
		}
		newMsg := &rpc.Message{
			Chat:     req.GetChat(),
			Text:     msg.Message,
			Sender:   msg.Sender,
			SendTime: msg.Timestamp,
		}
		respMessages = append(respMessages, newMsg)
		counter += 1
	}

	resp := rpc.NewPullResponse()
	resp.Messages, resp.Code, resp.Msg = respMessages, 0, "success"
	resp.HasMore, resp.NextCursor = &hasMore, &nextCursor

	return resp, nil
}
