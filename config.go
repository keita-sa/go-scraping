package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Db struct {
	Host     string
	DbName   string
	Port     string
	User     string
	Password string
}

type Config struct {
	Db
	BaseURL          string
	DownloadBasePath string
}

func loadConfig() (*Config, error) {
	var conf Config

	_, err := os.Stat(filepath.Join(".", "conf", "config.yml"))
	if err == nil {
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(".", "conf"))

	// viperでファイルから設定を読みこむ
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file error: %w", err)
	}

	// viperで読みこんだ設定をConfig構造体に変換して、conf変数にセットする
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, fmt.Errorf("unmarshal config file error: %w", err)
	}

	return &conf, nil
}
