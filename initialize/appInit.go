package initialize

import "vizhi_atlas/internal/pkg/globals"

func InitApp() error {
	// 初始化配置
	err := initConfig()
	if err != nil {
		return err
	}

	// 初始化日志
	err = initLogger(globals.C)
	if err != nil {
		return err
	}

	// 初始化数据库
	err = initDb(globals.C)
	if err != nil {
		return err
	}

	// 初始化gin引擎
	initEngine()

	return nil
}
