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

// 封装Redis操作String操作的一些方法

// StrAdd:插入set集合
// flag:标志用于区分不同Str集合数据
// key,value:分别是存入的key key与flag结合生成独一无二的key
// timeOut:过期时间
func StrAdd(db *redis.Client, flag string, key string, value interface{}, timeOut time.Duration) {
	db.Set(PartSplice(Partial(flag), key), value, timeOut)
}

// IsStr:判断key是否在redis中
func IsStr(db *redis.Client, flag string, key string) bool {
	res, err := db.Get(PartSplice(Partial(flag), key)).Result()
	if err != nil {
		return false
	}
	return res != ""
}

// GetStr:获取Key的值
func GetStr(db *redis.Client, flag string, key string) (string, bool) {
	res, err := db.Get(PartSplice(Partial(flag), key)).Result()
	if err != nil {
		return "", false
	}
	return res, true
}

// 封装Redis Hash一些方法

// HashAdd:map加入
func HashAdd(db *redis.Client, key string, field string, value interface{}) {
	db.HSet(Partial(key), field, value)
}

// HashAddSplice:map加入 带拼接
func HashAddSplice(db *redis.Client, key string, splice string, field string, value interface{}) {
	db.HSet(PartSplice(Partial(key), splice), field, value)
}

// HashMAddSplice:map批量加入
func HashMAddSplice(db *redis.Client, key string, splice string, objs map[string]interface{}) {
	db.HMSet(PartSplice(Partial(key), splice), objs)
}

// HashIsExists:判断是否存在 存在为true
func HashIsExists(db *redis.Client, key string, splice string, field string) bool {
	v, err := db.HExists(PartSplice(Partial(key), splice), field).Result()
	if err != nil {
		zap.L().Error("HashIsExists", zap.String("key", PartSplice(Partial(key), splice)), zap.Error(err))
		return false
	}
	return v
}

// HashContrast:查找判断是否有效
func HashContrast(db *redis.Client, key string, field string, value string) bool {
	// 获取uid对应的token 跟accessToken进行对比
	old, err := db.HGet(Partial(key), field).Result()
	if err != nil {
		return false
	}
	return old == value
}

// HashChange:更改
func HashChange(db *redis.Client, flag string, key string, field string, inc int64) {
	db.HIncrBy(PartSplice(Partial(flag), key), field, inc)
}

// HashDelete:删除
func HashDelete(db *redis.Client, key string) {
	db.Del(key)
}

// 封装Redis操作ZSET的一些方法

// ZSetAdd:新增
func ZSetAdd(db *redis.Client, key string, objs ...redis.Z) {
	db.ZAdd(Partial(key), objs...)
}

// ZSetAddSplice:新增 带参数
func ZSetAddSplice(db *redis.Client, flag string, key string, objs ...redis.Z) {
	db.ZAdd(PartSplice(Partial(flag), key), objs...)
}

// IsStr:判断key是否在redis中
func IsZSet(db *redis.Client, flag string, key string, member string) bool {
	res, err := db.ZRank(PartSplice(Partial(flag), key), member).Result()
	if err != nil {
		return false
	}
	return res >= 0
}

// ZSetChangeV:修改member值
func ZSetChangeV(db *redis.Client, flag string, key string, increment float64, member string) float64 {
	res, err := db.ZIncrBy(PartSplice(Partial(flag), key), increment, member).Result()
	if err != nil {
		zap.L().Error("ZSetChangeV:", zap.String("key", key), zap.Error(err))
		return -1
	}
	return res
}

// ZSetRemove:删除member值
func ZSetRemove(db *redis.Client, flag string, key string, member string) {
	db.ZRem(PartSplice(Partial(flag), key), member)
}
