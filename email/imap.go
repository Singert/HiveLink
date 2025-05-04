package email

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Message 封装了从 Mailpit 接收的邮件简略信息
type Message struct {
	ID      string `json:"ID"`
	Subject string `json:"Subject"`
	From    struct {
		Address string `json:"Address"`
	} `json:"From"`
}

// MailpitClient 通过 REST API 与 Mailpit 交互
type MailpitClient struct {
	apiBaseURL string
}

// NewMailpitClient 创建新的 Mailpit REST 客户端
func NewMailpitClient(apiBaseURL string) *MailpitClient {
	return &MailpitClient{apiBaseURL: apiBaseURL}
}

// FetchLatestSubject 获取最新邮件的主题
func (c *MailpitClient) FetchLatestSubject() (string, error) {
	url := fmt.Sprintf("%s/api/v1/messages", c.apiBaseURL)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("无法获取邮件列表: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应体失败: %v", err)
	}

	var data struct {
		Messages []Message `json:"Messages"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("解析 JSON 失败: %v", err)
	}

	if len(data.Messages) == 0 {
		return "", nil
	}

	latest := data.Messages[0]
	return latest.Subject, nil
}
