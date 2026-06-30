package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type LogEntry struct {
	Project string `json:"project"`
	Env     string `json:"env"`
	Level   string `json:"level"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
	Time    string `json:"time"`
}

var levels = []string{"INFO", "WARN", "ERROR"}
var envs = []string{"dev", "staging", "prod"}
var messages = []string{
	"User logged in",
	"Payment processed",
	"Item added to cart",
	"Transaction failed",
	"Password changed",
}

func main() {
	log.Println("Starting Log Generator... Bắn log tới Endpoint Fluent Bit tại http://fluent-bit:8888/app.logs")

	client := &http.Client{Timeout: 5 * time.Second}

	for {
		// Ngủ từ 2 đến 5 giây để sinh log nhanh hơn cho 3 môi trường
		sleepDuration := time.Duration(rand.Intn(4)+2) * time.Second
		time.Sleep(sleepDuration)

		entry := LogEntry{
			Project: "Project A",
			Env:     envs[rand.Intn(len(envs))],
			Level:   levels[rand.Intn(len(levels))],
			Message: messages[rand.Intn(len(messages))],
			UserID:  fmt.Sprintf("%d", rand.Intn(90000)+10000),
			Time:    time.Now().UTC().Format(time.RFC3339),
		}

		payload, err := json.Marshal(entry)
		if err != nil {
			log.Printf("Lỗi tạo JSON: %v", err)
			continue
		}

		// Gửi POST request tới Fluent Bit Endpoint
		req, err := http.NewRequest("POST", "http://fluent-bit:8888/app.logs", bytes.NewBuffer(payload))
		if err != nil {
			log.Printf("Lỗi tạo request: %v", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Lỗi khi bắn log tới Fluent Bit: %v", err)
			continue
		}
		
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			log.Printf("Đã bắn log thành công: %s", string(payload))
		} else {
			log.Printf("Bắn log thất bại. Status code: %d", resp.StatusCode)
		}
		resp.Body.Close()
	}
}
