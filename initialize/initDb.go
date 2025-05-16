package initialize

import (
	"vizhi_atlas/internal/pkg/database"
	"vizhi_atlas/internal/pkg/globals"
	"vizhi_atlas/internal/pkg/logger"
)

// InitDb 初始化数据库
func initDb(c *globals.Config) error {
	// 初始化MySQL连接
	if err := database.Init(c); err != nil {
		logger.Error("初始化MySQL失败", logger.String("error", err.Error()))
		return err
	}

	// 这里可以添加数据库迁移代码
	if err := autoMigrate(); err != nil {
		logger.Error("数据库迁移失败", logger.String("error", err.Error()))
		return err
	}

	return nil
}

// autoMigrate 自动迁移数据库表结构
func autoMigrate() error {
	// 在这里添加需要自动迁移的模型
	// 例如：
	// if err := database.DB.AutoMigrate(&model.User{}); err != nil {
	//     return err
	// }
	return nil
}
