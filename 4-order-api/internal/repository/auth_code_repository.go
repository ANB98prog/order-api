package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const AuthCodeKeyPattern = "AUTH_CODE_%v"

func getKey(sessionId string) string {
	return fmt.Sprintf(AuthCodeKeyPattern, sessionId)
}

type AuthCodeRepository interface {
	Save(code AuthCode, ttl time.Duration) (bool, error)
	GetBySessionId(sessionId string) (AuthCode, bool)
	Delete(sessionId string)
}

var ctx = context.Background()

type redisAuthCodeRepository struct {
	rdb *redis.Client
}

func NewRedisAuthCodeRepository(rdb *redis.Client) AuthCodeRepository {
	return &redisAuthCodeRepository{rdb: rdb}
}

func (r *redisAuthCodeRepository) Save(code AuthCode, ttl time.Duration) (bool, error) {
	data, err := json.Marshal(code)
	if err != nil {
		return false, err
	}

	ok, err := r.rdb.SetNX(ctx, getKey(code.SessionId), data, ttl).Result()
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (r *redisAuthCodeRepository) GetBySessionId(sessionId string) (AuthCode, bool) {
	var authCode AuthCode
	result, err := r.rdb.Get(ctx, getKey(sessionId)).Result()
	if err != nil {
		return AuthCode{}, false
	}

	err = json.Unmarshal([]byte(result), &authCode)
	if err != nil {
		return AuthCode{}, false
	}

	return authCode, true
}

func (r *redisAuthCodeRepository) Delete(sessionId string) {
	r.rdb.Del(ctx, getKey(sessionId))
}
