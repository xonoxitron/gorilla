package handlers

import (
	"strings"

	"github.com/xonoxitron/gorilla/config"
	"github.com/xonoxitron/gorilla/storage"
)

func Start() string {
	return "[gorilla] bot v." + config.Version()
}

func SubscribeUser(id string) string {
	content := storage.Get("subscribers")

	if strings.Contains(content, id) {
		return "already subscribed"
	}

	storage.Update("subscribers", id+"\r", false)

	return "subscribed"
}

func UnsubscribeUser(id string) string {
	content := storage.Get("subscribers")

	if !strings.Contains(content, id) {
		return "already unsubscribed"
	}

	updatedContent := strings.TrimSpace(strings.ReplaceAll(content, id, ""))

	storage.Update("subscribers", updatedContent, true)

	return "unsubscribed"
}
