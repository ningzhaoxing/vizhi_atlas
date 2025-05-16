package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"vizhi_atlas/internal/pkg/globals"
)

func initConfig() error {
	viper.SetConfigFile(getConfigPath())
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config Load error %s \n", err.Error())
		return err
	}

	var config globals.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("config bind error %s \n", err.Error())
		return err
	}
	fmt.Println(config)

	globals.C = &config
	return nil
}

func getConfigPath() string {
	return "./configs/dev.yaml"
}
