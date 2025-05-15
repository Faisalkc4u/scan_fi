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

	http.PostForm(endpoint, data)
}
func LogError(tag string, err error) {
	if err != nil {
		SendToTelegram(fmt.Sprintf("❌ [%s] %v", tag, err))
	}
}
func LogInfo(tag string, message string) {
	SendToTelegram(fmt.Sprintf("✅ [%s] %s", tag, message))
}
