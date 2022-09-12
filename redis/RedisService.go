package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

/**
 * @title redis操作
 * @author xiongshao
 * @date 2022-06-27 08:30:08
 */

var ctx = context.Background()
var clinet = &redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
}

// Get 获取缓存数据
func GetKey(key string) (string, error) {
	result, err := redis.NewClient(clinet).Get(ctx, key).Result()
	return result, err
}

// Sey 设置数据 永久有效
func SetKey(key string, value interface{}) {
	redis.NewClient(clinet).Set(ctx, key, value, 0).Err()
}

// // Set 设置数据 包含过期时间
func SetKeyTime(key string, value interface{}, expiration time.Duration) error {
	err := redis.NewClient(clinet).Set(ctx, key, value, expiration).Err()
	return err
}

func redisGetKeyByValues(str string) interface{} {
	key, _ := GetKey(str)
	return key
}

// 判断一个key在redis数据库中是否存在 返回查询个数
func KeyExists(key string) int64 {
	result, err := redis.NewClient(clinet).Exists(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	return result

}

// 获取当前redis库所有key值
func GetKeyList() []string {
	var cursor uint64
	keys, cursor, err := redis.NewClient(clinet).Scan(ctx, cursor, "*", 1000000).Result()
	for err != nil {
		log.Panicln("scan keys failed err:", err)
	}
	return keys
}
