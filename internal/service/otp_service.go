package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/passworless-auth-server/internal/config"
	"github.com/AkifhanIlgaz/passworless-auth-server/pkg/helper"
	"github.com/AkifhanIlgaz/passworless-auth-server/pkg/otp"
	"github.com/redis/go-redis/v9"
)

type OTPService struct {
	ctx         context.Context
	redisClient *redis.Client
	config      *config.Config
}

func NewOTPService(ctx context.Context, redisClient *redis.Client, config *config.Config) *OTPService {
	return &OTPService{
		ctx:         ctx,
		redisClient: redisClient,
		config:      config,
	}
}

func (s *OTPService) Store(uid string) (string, error) {
	code := otp.CreateCode(s.config.OTP.Length)
	ttl := time.Duration(s.config.OTP.Ttl) * time.Minute
	key := helper.CreateRedisKey("otp", uid)

	exists, err := s.redisClient.Exists(s.ctx, key).Result()
	if err != nil {
		return "", err
	}
	if exists > 0 {
		return "", fmt.Errorf("OTP for user %s already exists", uid)
	}

	if err := s.redisClient.Set(s.ctx, key, code, ttl).Err(); err != nil {
		return "", err
	}

	return code, nil
}
