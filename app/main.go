package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type LogEntry struct {
	Job     string `json:"job"`
	Env     string `json:"env"`
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
	log.Println("Starting Log Generator... Writing to /var/log/app/app.log")

	// Create directory if it doesn't exist
	os.MkdirAll("/var/log/app", os.ModePerm)

	// Open file for writing
	f, err := os.OpenFile("/var/log/app/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer f.Close()

	for {
		// Wait 1 to 3 seconds
		sleepDuration := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(sleepDuration)

		entry := LogEntry{
			Job:     "golang-payment-service",
			Env:     "dev",
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

		// Write to file instead of sending via HTTP
		f.Write(payload)
		f.WriteString("\n")
		f.Sync() // Ensure it's written to disk immediately

		log.Printf("Wrote log: %s", string(payload))
	}
}
