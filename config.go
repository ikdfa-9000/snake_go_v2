package main

import "github.com/spf13/viper"

type Config struct {
	deskRows       int
	deskColumns    int
	deskFrameSpeed int
}

func InitConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	config.deskRows = viper.GetInt("desk.rows")
	config.deskColumns = viper.GetInt("desk.columns")
	config.deskFrameSpeed = viper.GetInt("desk.frameSpeed")

	return config, nil
}
