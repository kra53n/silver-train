package util

import (
	"bytes"
	"fmt"
	"time"
	"net/http"
	"encoding/json"

	"silver-train/vars"
)

func SendMsg(webhookURL string, msg string) error {
	payload := struct {
		Msg       string `json:"msg"`
		Timestamp int64  `json:"timestamp"`
	}{
		Msg:       msg,
		Timestamp: time.Now().Unix(),
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send post request: %w", err)
	}
	defer resp.Body.Close()
	return nil
}

func SendMsgAtWebHook(msg string) {
	if err := SendMsg(vars.MsgWebhook(), msg); err != nil {
		fmt.Println("[ERROR]  occured while trying to send msg. Err message:", err)
	}
}
