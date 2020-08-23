/******
** @创建时间 : 2020/8/11 17:44
** @作者 : SongZhiBin
******/
package logger

import (
	"Happy/settings"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var t *time.Timer

// Init:logger初始化
func Init() error {
	// 定制logger
	// 指定日志将写到哪里去
	writeSyncer := getLogWriter()
	// 日志格式
	encoder := getEncoder()
	// 创建自定义logger
	var core zapcore.Core
	if settings.GetString("APP.Mode") != "release" {
		// 非发布版本
		devCore := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		// 双输出
		core = zapcore.NewTee(
			// 添加日志输入
			zapcore.NewCore(encoder, writeSyncer, getLever()),
			// 添加到终端输出
			zapcore.NewCore(devCore, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, getLever())
	}
	lg := zap.New(core, zap.AddCaller())
	// 替换zap库全局的logger对象
	// 使用zap.L() 调用全局对象
	zap.ReplaceGlobals(lg)
	if settings.GetBool("LOGGER.IsAddData") {
		go guard()
	}
	// 启动哨兵
	return nil
}

// guard:哨兵,根据日志切分日志记录
func guard() {
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		<-t.C
		// 先进行日志的退出
		_ = zap.L().Sync()
		// 重新调用Init创建新的句柄
		_ = Init()
		time.Sleep(time.Minute)
	}
}

// Close:关于回收资源
func Close() {
	_ = zap.L().Sync()
}

// getEncoder:返回编码器(日志格式)
func getEncoder() zapcore.Encoder {
	// 设置生产配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改日期格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 修改日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter:返回日志设置
func getLogWriter() zapcore.WriteSyncer {
	// 调用settings配置管理
	// 使用第三方库剪切日志
	var filename string
	if settings.GetBool("LOGGER.IsAddData") {
		// filename: LogPath+LogPrefix+日期+LogSuffix
		filename = fmt.Sprintf("%s/%s%s.%s",
			settings.GetString("LOGGER.LogPath"),
			settings.GetString("LOGGER.LogPrefix"),
			time.Now().Format("2006:01:02"),
			settings.GetString("LOGGER.LogSuffix"))
	} else {
		filename = fmt.Sprintf("%s/%s.%s",
			settings.GetString("LOGGER.LogPath"),
			settings.GetString("LOGGER.LogPrefix"),
			settings.GetString("LOGGER.LogSuffix"))
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    settings.GetInt("LOGGER.MaxSize"),
		MaxBackups: settings.GetInt("LOGGER.MaxBackups"),
		MaxAge:     settings.GetInt("LOGGER.MaxAge"),
		Compress:   settings.GetBool("LOGGER.Compress"),
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getLever:解析配置文件的log lever
func getLever() *zapcore.Level {
	var l = new(zapcore.Level)
	// 提供的内置方法 通过字符串获取到lever
	err := l.UnmarshalText([]byte(settings.GetString("LOGGER.Lever")))
	// 如果出错 默认为debug模式
	if err != nil {
		fmt.Println("lever UnmarshalText error,Use Default Debug!:error", err)
		*l = -1
		return l
	}
	return l
}

// Middleware

// GinLogger:接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
