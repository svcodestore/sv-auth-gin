package system

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/svcodestore/sv-auth-gin/global"
)

const (
	GrantedCodeRedisKey        = "grantedCode"
	IssuedAccessTokenRedisKey  = "issuedAccessToken"
	IssuedRefreshTokenRedisKey = "issuedRefreshToken"
)

type OauthService struct {
}

func (s OauthService) deleteTokenFromRedis(key string) (affected int64, err error) {
	ctx := context.Background()
	affected, err = global.REDIS.Del(ctx, key).Result()
	if err != redis.Nil {
		err = nil
	}
	return
}

func (s *OauthService) DeleteAccessTokenFromRedis(userId string) (affected int64, err error) {
	k := IssuedAccessTokenRedisKey + ":" + userId
	affected, err = s.deleteTokenFromRedis(k)
	return
}

func (s *OauthService) IsUserLogin(userId string) (isLogin bool) {
	ctx := context.Background()
	k := IssuedAccessTokenRedisKey + ":" + userId
	count, err := global.REDIS.Exists(ctx, k).Result()
	if err == nil {
		if count == 1 {
			isLogin = true
			return
		}
	}
	return
}
