/******
** @创建时间 : 2020/8/11 20:38
** @作者 : SongZhiBin
******/
package redis

import (
	"Happy/settings"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

// rdb:redis全局变量
var rdb *redis.Client

// Init:Redis初始化
func Init() error {
	rdb = CRedis(
		settings.GetString("REDIS.Host"),
		settings.GetInt("REDIS.Port"),
		settings.GetString("REDIS.Password"),
		settings.GetInt("REDIS.Block"))
	if rdb == nil {
		return errors.New("Redis init is nil ")
	}
	return nil

}

// CRedis:获取redis对象
func CRedis(host string, port int, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			host,
			port),
		Password: password,
		DB:       db,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		zap.L().Error(fmt.Sprintf("Redis No Connect Error:%s", err))
		return nil
	}
	return rdb
}

// Close:回收资源
func Close() {
	_ = rdb.Close()
}

// 封装Set操作的一些方法

// StrAdd:插入set集合
// flag:标志用于区分不同Str集合数据
// key,value:分别是存入的key key与flag结合生成独一无二的key
// timeOut:过期时间
func StrAdd(db *redis.Client, flag string, key string, value interface{}, timeOut time.Duration) {
	db.Set(flag+key, value, timeOut)
}

// IsStr:判断key是否在redis中
func IsStr(db *redis.Client, flag string, key string) bool {
	res, err := db.Get(flag + key).Result()
	if err != nil {
		return false
	}
	return res != ""
}

// GetStr:获取Key的值
func GetStr(db *redis.Client, flag string, key string) (string, bool) {
	res, err := db.Get(flag + key).Result()
	if err != nil {
		return "", false
	}
	return res, true
}
