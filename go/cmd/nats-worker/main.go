package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Job struct {
	RunID string `json:"run_id"`
	User  User   `json:"user"`
}

type EmailResult struct {
	RunID   string `json:"run_id"`
	UserID  string `json:"user_id"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func sendEmails(user User) error {
	time.Sleep(200 * time.Millisecond)

	return nil
}

func main() {
	url := nats.DefaultURL
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}

}
