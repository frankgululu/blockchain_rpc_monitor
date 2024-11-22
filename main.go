package main

import (
	"context"
	"fmt"
	"github.com/b2network/tools/config"
	"github.com/b2network/tools/initialize"
	"github.com/b2network/tools/server"
	"github.com/robfig/cron/v3"
	"k8s.io/klog/v2"
	"strings"
)

func main() {
	//初始化配置
	config.NewConfig()
	//初始化redis
	rc := initialize.NewRedisClient()
	defer rc.Close()
	domainStr := config.Cfg.DomainList
	domainList := strings.Split(domainStr, ",")
	fmt.Println(domainList)
	//启动定时任务
	//测试redis
	ctx := context.Background()
	//server.SendTgMessage("bot7034301781:AAEuF76uwpisf1t58gNCdAYwazMTdcgimXo", "-1002028001234", "test")
	//rc.SetKey(ctx, "dyx", "123", 0)
	//value, _ := rc.GetKey(ctx, "dyx")
	//fmt.Println(value)
	crontab := cron.New()
	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	crontab.AddFunc("*/5 * * * *", func() {
		for _, domain := range domainList {
			fmt.Println(domain)
			_, newBlockNum, result := server.ComPareBlockNum(ctx, rc, domain)
			if result {
				rc.SetKey(ctx, domain, string(newBlockNum), 0)
				klog.Infof("%v 出块正常 number %v \n", domain, newBlockNum)
				//fmt.Printf("%v 出块正常 number %v \n", domain, newBlockNum)

			} else {
				rc.SetKey(ctx, domain, newBlockNum, 0)
				msg := domain + "出块异常 number:" + string(newBlockNum)
				//fmt.Printf("%v 出块异常 number %v \n", domain, newBlockNum)
				klog.Errorf("%v 出块异常 number %v \n", domain, newBlockNum)
				server.SendTgMessage(config.Cfg.Telegram.BootId, config.Cfg.Telegram.ChatId, msg)
			}
		}
	})

	// 启动定时器
	crontab.Start()

	select {}
}
