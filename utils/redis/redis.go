package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lfyr/go-api/utils"
	"github.com/sirupsen/logrus"
	"time"
)

var client *redis.Client

func init() {
	redisCfg := utils.GVA_CONFIG.Redis
	if redisCfg.Addr == "" {
		logrus.Error(redisCfg)
		return
	}
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
		Username: redisCfg.Username, // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Error("redis connect ping failed, err:", err.Error())
	} else {
		logrus.Infof("redis connect ping response:%+v", pong)
	}
}

func NewDefaultRedisStore(expiration time.Duration) *RedisStore {
	return &RedisStore{
		Expiration: time.Second * expiration,
		PreKey:     "APPSTORE_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

//func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
//	rs.Context = ctx
//	return rs
//}

func (rs *RedisStore) Set(key string, value string) error {
	if rs.Context == nil {
		rs.Context = context.Background()
	}
	err := client.Set(rs.Context, rs.PreKey+key, value, rs.Expiration).Err()
	if err != nil {
		logrus.Error("RedisStoreSetError!", err.Error())
	}
	return err
}

func (rs *RedisStore) Get(key string, clear bool) string {
	if rs.Context == nil {
		rs.Context = context.Background()
	}
	val, err := client.Get(rs.Context, rs.PreKey+key).Result()
	if err != nil {
		logrus.Error("RedisStoreGetError!", err.Error())
		return ""
	}
	if clear {
		err = client.Del(rs.Context, key).Err()
		if err != nil {
			logrus.Error("RedisStoreClearError!", err.Error())
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

func (rs *RedisStore) SetNX(key string, value string) error {
	if rs.Context == nil {
		rs.Context = context.Background()
	}
	err := client.SetNX(rs.Context, rs.PreKey+key, value, rs.Expiration).Err()
	if err != nil {
		logrus.Error("RedisStoreSetNXError!", err.Error())
	}
	return err

}

func (rs *RedisStore) Del(key string) error {
	if rs.Context == nil {
		rs.Context = context.Background()
	}
	err := client.Del(rs.Context, rs.PreKey+key).Err()
	if err != nil {
		logrus.Error("RedisStoreDelError!", err.Error())
	}
	return err
}
