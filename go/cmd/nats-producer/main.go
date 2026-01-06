package main

import (
	"fmt"
	"log"
	"strconv"
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

func makeUsers(userNum int) []User {
	users := make([]User, 0, userNum)
	for i := 1; i <= userNum; i++ {
		userID := "user" + strconv.Itoa(i)
		email := userID + "@example.com"

		user := User{
			ID:    userID,
			Email: email,
		}

		users = append(users, user)
	}

	return users
}

func main() {
	url := nats.DefaultURL
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	users := makeUsers(10)
	runID := fmt.Sprintf("%d", time.Now())
}
