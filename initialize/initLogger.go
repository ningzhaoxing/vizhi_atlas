package initialize

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"vizhi_atlas/internal/pkg/globals"
	"vizhi_atlas/internal/pkg/logger"
)

func initLogger(c *globals.Config) error {
	// 获取当前日期作为日志文件名
	currentDate := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("logs/%s.log", currentDate)

	// 确保日志目录存在
	logDir := filepath.Dir(logFileName)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 初始化日志
	if err := logger.Init(c); err != nil {
		return fmt.Errorf("初始化日志失败: %w", err)
	}

	logger.Info("日志系统初始化成功", logger.String("logFile", logFileName))
	return nil
}
