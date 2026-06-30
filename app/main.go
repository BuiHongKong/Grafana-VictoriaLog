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
	Stream  string `json:"_stream"`
	Level   string `json:"level"`
	Message string `json:"message"`
	UserID  string `json:"user_id"`
	Time    string `json:"time"`
}

var (
	levels   = []string{"info", "warn", "error"}
	messages = []string{"Transaction failed", "User logged in", "Item added to cart", "Checkout initiated", "Payment processed"}
)

func main() {
	url := "http://victorialogs:9428/insert/jsonline"
	
	log.Println("Starting Log Generator...")
	
	client := &http.Client{Timeout: 10 * time.Second}

	for {
		// Wait 1 to 3 seconds
		sleepDuration := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(sleepDuration)

		entry := LogEntry{
			Stream:  `{job="golang-payment-service", env="dev"}`,
			Level:   levels[rand.Intn(len(levels))],
			Message: messages[rand.Intn(len(messages))],
			UserID:  fmt.Sprintf("%d", rand.Intn(99999)+10000),
			Time:    time.Now().UTC().Format(time.RFC3339),
		}

		payload, err := json.Marshal(entry)
		if err != nil {
			log.Printf("Error marshaling log entry: %v", err)
			continue
		}

		// Push to VictoriaLogs
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			log.Printf("Error creating request: %v", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error sending log: %v", err)
			continue
		}
		
		if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
			log.Printf("Unexpected status code: %d", resp.StatusCode)
		}
		
		resp.Body.Close()
		log.Printf("Sent log: %s", string(payload))
	}
}
