package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

/**
 * @title redis操作
 * @author xiongshao
 * @date 2022-06-27 08:30:08
 */

// 声明一个全局的rdb变量
var rdb *redis.Client
var ctx context.Context
var cancel context.CancelFunc

// 初始化连接redis
func InitClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// Redis哨兵模式
func InitClient2() (err error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"localhost:26379", "localhost:26379", "localhost:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// Redis集群模式
func InitClient3() (err error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// Get 获取缓存数据
func GetKey(key string) (string, error) {
	result, err := rdb.Get(key).Result()
	return result, err
}

// Set 设置数据 包含过期时间
func SetKeyTime(key string, value interface{}, expiration time.Duration) error {
	err := rdb.Set(key, value, expiration).Err()
	return err
}

// Sey 设置数据 默认时间为24h
func SetKey(key string, value interface{}) error {
	err := rdb.Set(key, value, time.Hour*24).Err()
	return err
}

//
func redisGetKeyByValues(str string) interface{} {
	key, _ := GetKey(str)
	return key
}

// 判断一个key在redis数据库中是否存在 返回查询个数
func KeyExists(key string) int64 {
	result, err := rdb.Exists(key).Result()
	if err != nil {
		fmt.Println(key, "检索失败")
	}
	return result
}

// 获取当前redis库所有key值
func GetKeyList() []string {
	InitClient()
	var cursor uint64
	keys, cursor, err := rdb.Scan(cursor, "*", 1000000).Result()
	for err != nil {
		fmt.Println("scan keys failed err:", err)
	}
	return keys
}
