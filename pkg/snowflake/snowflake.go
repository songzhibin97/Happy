package snowflake

import (
	"Happy/settings"
	"errors"
	"time"

	"github.com/sony/sonyflake"
	"go.uber.org/zap"
)

// 第三方包:分布式全局ID生成

// LSonyFlake:全局对象
var LSonyFlake *sonyflake.Sonyflake

// GetMachineID:获取MachineID
func GetMachineID() (uint16, error) {
	return uint16(settings.GetUint("SNOWFLAKE.MachineID")), nil
}

// GetStartTime:获取起始时间
func GetStartTime() time.Time {
	// 1.将字符串转化为时间戳
	startTimeString := settings.GetString("SNOWFLAKE.StartTime")
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		zap.L().Error("Load Time Zone Error", zap.Error(err))
		// 如果失败的情况下默认都返回当前时间
		return time.Now()
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006.01.02", startTimeString, loc)
	if err != nil {
		zap.L().Error("ParseInLocation Error", zap.Error(err))
		// 如果失败的情况下默认都返回当前时间
		return time.Now()
	}
	zap.L().Info("GetStartTime Success", zap.String("startTime", timeObj.String()))
	return timeObj
}

// Init:初始化
// startTime:起始时间,作为偏移量使用
// machineID:获取machineID
func Init() {
	st := sonyflake.Settings{}
	st.MachineID = GetMachineID
	st.StartTime = GetStartTime()
	// st.CheckMachineID = nil 表示不对MachineID进行校验
	LSonyFlake = sonyflake.NewSonyflake(st)
	return
}

// GetID
func GetID() (uint64, error) {
	if LSonyFlake == nil {
		Init()
	}
	if LSonyFlake == nil {
		zap.L().Error("LSonyFlake is Nil Not Exec GetID")
		return 0, errors.New("LSonyFlake is Nil Not Exec GetID")
	}
	return LSonyFlake.NextID()
}
