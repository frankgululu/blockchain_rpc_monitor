package config

import (
	"fmt"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
	"os"
)

type RedisConfigurations struct {
	Addr     string
	Password string
	DB       int
}

type TelegramConfigurations struct {
	BootId string
	ChatId string
}

type Configurations struct {
	Redis      RedisConfigurations
	DomainList string
	Telegram   TelegramConfigurations
	Cron       string
}

// Cfg 配置文件
var Cfg Configurations

func NewConfig() {
	fmt.Println("init config...")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// info / debug / trace
	viper.BindEnv("redis.addr", "REDIS_ADDR")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("redis.db", "REDIS_DB")
	viper.BindEnv("domainList", "DOMAINLIST")
	viper.BindEnv("telegram.bootid", "TELEGRAM_BOOTID")
	viper.BindEnv("telegram.chatid", "TELEGRAM_CHATID")
	viper.BindEnv("cron", "CRON")
	if err := viper.ReadInConfig(); err != nil {
		klog.Error("Error load config file, %s\n", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		klog.Error("Error unmarshal config file, %s\n", err)
		os.Exit(1)
	}

}
