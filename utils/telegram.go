package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func SendToTelegram(message string) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", message)

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send message to Telegram: %v\n", err)
		return
	}
	defer resp.Body.Close()
}
func getServerName() string {
	serverName := os.Getenv("RENDER_SERVICE_NAME")
	if serverName == "" {
		serverName = os.Getenv("SERVER_NAME")
	}
	if serverName == "" {
		serverName = "unknown"
	}
	return serverName
}

func LogError(tag string, err error) {
	if err != nil {
		serverName := getServerName()
		SendToTelegram(fmt.Sprintf("❌ [%s][%s] %v", tag, serverName, err))
	}
}

func LogInfo(tag string, message string) {
	serverName := getServerName()
	SendToTelegram(fmt.Sprintf("✅ [%s][%s] %s", tag, serverName, message))
}
