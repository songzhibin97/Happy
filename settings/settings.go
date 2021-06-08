package settings

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init:Viper配置文件初始化
func Init() error {
	// 指定配置文件名称(无扩展名)
	viper.SetConfigName("config")
	// 配置文件类型
	viper.SetConfigType("ini")
	// 添加配置文件路径
	viper.AddConfigPath(".")
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	// 开启监控 实现热更新
	viper.WatchConfig()
	// 设置更新回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config Change Success:", in.String())
	})
	return nil
}

// 2020年08月15日新增
// 解决多个包引入 viper 获取配置参数的情况

// GetString:
func GetString(key string) string { return viper.GetString(key) }

// GetBool:
func GetBool(key string) bool { return viper.GetBool(key) }

// GetInt
func GetInt(key string) int { return viper.GetInt(key) }

// GetInt32
func GetInt32(key string) int32 { return viper.GetInt32(key) }

// GetInt64
func GetInt64(key string) int64 { return viper.GetInt64(key) }

// GetUint
func GetUint(key string) uint { return viper.GetUint(key) }

// GetUint32
func GetUint32(key string) uint32 { return viper.GetUint32(key) }

// GetUint64
func GetUint64(key string) uint64 { return viper.GetUint64(key) }

// GetFloat64
func GetFloat64(key string) float64 { return viper.GetFloat64(key) }

// GetTime
func GetTime(key string) time.Time { return viper.GetTime(key) }

// GetDuration
func GetDuration(key string) time.Duration { return viper.GetDuration(key) }

// GetIntSlice
func GetIntSlice(key string) []int { return viper.GetIntSlice(key) }

// GetStringSlice
func GetStringSlice(key string) []string { return viper.GetStringSlice(key) }

// GetStringMap
func GetStringMap(key string) map[string]interface{} { return viper.GetStringMap(key) }

// GetStringMapString
func GetStringMapString(key string) map[string]string { return viper.GetStringMapString(key) }

// GetStringMapStringSlice
func GetStringMapStringSlice(key string) map[string][]string {
	return viper.GetStringMapStringSlice(key)
}

// GetSizeInBytes
func GetSizeInBytes(key string) uint { return viper.GetSizeInBytes(key) }

// UnmarshalKey
func UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	return viper.UnmarshalKey(key, rawVal, opts...)
}
