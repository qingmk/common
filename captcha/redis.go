package captcha

import (
	"context"
	"time"

	"github.com/mojocn/base64Captcha"
	redis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
		Context:    context.TODO(),
	}
}

func NewDefaultRedisStoreV2(Redis *redis.ClusterClient) *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
		Context:    context.TODO(),
		Redis:      Redis,
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
	Redis      *redis.ClusterClient
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}
func (rs *RedisStore) Set(id string, value string) error {
	err := rs.Redis.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		logx.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := rs.Redis.Get(rs.Context, key).Result()
	if err != nil {
		logx.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := rs.Redis.Del(rs.Context, key).Err()
		if err != nil {
			logx.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
