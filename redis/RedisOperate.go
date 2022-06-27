package redis

import (
	"context"
	"github.com/go-redis/redis"
	"time"
)

/**
 * @title redis操作
 * @author xiongshao
 * @date 2022-06-27 08:30:08
 */

// 声明全局变量
var rdb *redis.Client
var ctx context.Context
var cancel context.CancelFunc

// 初始化连接redis
func initClient1() (err error) {
	redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// 设置过期时间
	context.WithTimeout(context.Background(), 5*time.Second)
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// Redis哨兵模式
func initClient2() (err error) {
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
func initClient3() (err error) {
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
func Get(key string) (string, error) {
	result, err := rdb.Get(key).Result()
	return result, err
}

// Set 设置数据 过期时间默认24H
func Set(key, value string) error {
	err := rdb.Set(key, value, time.Hour*24).Err()
	return err
}
