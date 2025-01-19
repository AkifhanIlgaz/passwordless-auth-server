package service

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type UserService struct {
	ctx         context.Context
	redisClient *redis.Client
}

func NewUserService(ctx context.Context, redisClient *redis.Client) *UserService {
	return &UserService{
		ctx:         ctx,
		redisClient: redisClient,
	}
}


func (s *UserService) Login(uid string) (string, error) {
	key := "user:" + uid
	email, err := s.redisClient.HGet(s.ctx, key, "email").Result()
	if err != nil {
		return "", err
	}
	return email, nil
}
