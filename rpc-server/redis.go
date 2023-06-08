package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	cli *redis.Client
}

func (c *RedisClient) InitClient(ctx context.Context, addr, pw string) error {
	r := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pw,
		DB:       0,
	})

	if err := r.Ping(ctx).Err(); err != nil {
		return err
	}

	c.cli = r
	return nil
}

func (c *RedisClient) SaveMessage(ctx context.Context, key string, message *Message) error {
	text, err := json.Marshal(message)
	if err != nil {
		return err
	}

	member := &redis.Z{
		Score:  float64(message.Timestamp),
		Member: text,
	}

	_, err = c.cli.ZAdd(ctx, key, *member).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *RedisClient) GetMessages(ctx context.Context, key string, start, end int64, reverse bool) ([]*Message, error) {
	var (
		raw      []string
		messages []*Message
		err      error
	)

	if reverse {
		raw, err = c.cli.ZRevRange(ctx, key, start, end).Result()
	} else {
		raw, err = c.cli.ZRange(ctx, key, start, end).Result()
	}
	
	if err != nil {
		return nil, err
	}

	for _, msg := range raw {
		newMsg := &Message{}
		err := json.Unmarshal([]byte(msg), newMsg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, newMsg)
	}

	return messages, nil
}
