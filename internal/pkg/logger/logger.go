package logger

import (
	"fmt"
	"os"
	"time"
	"vizhi_atlas/internal/pkg/globals"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Init 初始化日志
func Init(cfg *globals.Config) error {
	// 设置日志级别
	var level zapcore.Level
	err := level.UnmarshalText([]byte(cfg.Logger.Level))
	if err != nil {
		return err
	}

	// 创建日志目录
	logDir := "./logs"
	if err = os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 配置 encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 配置输出
	var cores []zapcore.Core

	// 文件输出 - 按天切割
	currentDate := time.Now().Format("2006-01-02")
	logFilename := fmt.Sprintf("%s/%s.log", logDir, currentDate)

	// 使用 lumberjack 进行日志切割
	fileWriter := &lumberjack.Logger{
		Filename:   logFilename,
		MaxSize:    cfg.Logger.MaxSize,    // 每个日志文件的最大大小，单位MB
		MaxBackups: cfg.Logger.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     cfg.Logger.MaxAge,     // 保留旧文件的最大天数
		Compress:   cfg.Logger.Compress,   // 是否压缩
		LocalTime:  true,                  // 使用本地时间
	}

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	fileCore := zapcore.NewCore(
		fileEncoder,
		zapcore.AddSync(fileWriter),
		level,
	)
	cores = append(cores, fileCore)

	// 控制台输出
	if cfg.Logger.Console {
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		core := zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			level,
		)
		cores = append(cores, core)
	}

	// 创建 logger
	core := zapcore.NewTee(cores...)
	globals.Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// 记录初始化成功日志
	Info("日志系统初始化成功",
		String("level", cfg.Logger.Level),
		String("filename", logFilename))

	return nil
}

// timeEncoder 自定义时间编码器
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// Debug 输出 Debug 级别日志
func Debug(msg string, fields ...zap.Field) {
	globals.Log.Debug(msg, fields...)
}

// Info 输出 Info 级别日志
func Info(msg string, fields ...zap.Field) {
	globals.Log.Info(msg, fields...)
}

// Warn 输出 Warn 级别日志
func Warn(msg string, fields ...zap.Field) {
	globals.Log.Warn(msg, fields...)
}

// Error 输出 Error 级别日志
func Error(msg string, fields ...zap.Field) {
	globals.Log.Error(msg, fields...)
}

// Fatal 输出 Fatal 级别日志
func Fatal(msg string, fields ...zap.Field) {
	globals.Log.Fatal(msg, fields...)
}

// With 创建带有额外字段的 globals.Logger
func With(fields ...zap.Field) *zap.Logger {
	return globals.Log.With(fields...)
}

// Sync 同步日志缓冲
func Sync() error {
	return globals.Log.Sync()
}

// String 创建字符串类型的日志字段
func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

// Int 创建整数类型的日志字段
func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

// Bool 创建布尔类型的日志字段
func Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

// Any 创建任意类型的日志字段
func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// ErrorZ 创建错误类型的日志字段
func ErrorZ(err error) zap.Field {
	return zap.Error(err)
}
