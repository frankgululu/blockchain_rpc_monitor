package initialize

import (
	"context"
	"github.com/b2network/tools/config"
	"github.com/go-redis/redis/v8"
	"k8s.io/klog/v2"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	//fmt.Println(config.Cfg.Redis.Addr)
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     config.Cfg.Redis.Addr,
			Password: config.Cfg.Redis.Password,
			DB:       config.Cfg.Redis.DB,
		}),
	}

}

func (r *RedisClient) SetKey(ctx context.Context, keyName string, Value string, expiration time.Duration) error {
	return r.client.Set(ctx, keyName, Value, expiration).Err()
}

func (r *RedisClient) GetKey(ctx context.Context, keyName string) (string, error) {
	cmd := r.client.Get(ctx, keyName)
	return cmd.Val(), cmd.Err()
}
func (r *RedisClient) Close() {
	err := r.client.Close()
	if err != nil {
		klog.Error(err)

	}
}
