package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type Task struct {
	To       string `json:"to"`
	Template string `json:"template"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type EmailJob struct {
	JobID string `json:"job_id"` // 一意（例: runID + ":" + userID）
	RunID string `json:"run_id"` // 1回の実験単位（結果の紐付け用）
	User  User   `json:"user"`   // 今回はDBなしで完結させたいので丸ごと入れる
}

type EmailResult struct {
	RunID        string `json:"run_id"`
	JobID        string `json:"job_id"`
	UserID       string `json:"user_id"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"error_message,omitempty"`
}

//func ensureStream(js nats.JetStreamContext) error {
// func ensureStream() error {
// 	js, _ := jetstream.New(nc)
// 	cfg := jetstream.StreamConfig{
// 		Name:      "Email",
// 		Subjects:  []string{"mail.job"},
// 		Retention: jetstream.WorkQueuePolicy,
// 		Storage:   jetstream.FileStorage,
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	stream, _ := js.CreateStream(ctx, cfg)
// 	fmt.Println("created the stream")
// 	return err
// }

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
	// 常駐用　ctrl + cなどで止める
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	// 上のやつはよく分からん

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	cfg := jetstream.StreamConfig{
		Name:      "Email",
		Subjects:  []string{"mail.job"},
		Retention: jetstream.WorkQueuePolicy,
		Storage:   jetstream.FileStorage,
	}

	// 単発用のctx　時間で止まる
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = js.CreateStream(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created the stream")

	//users := makeUsers(10)

	task := Task{
		To:       "user1@example.com",
		Template: "welcome",
	}
	data, err := json.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}

	ack, err := js.Publish(ctx, "tasks.email", data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("published stream=%s, seq=%d", ack.Stream, ack.Sequence)

	// for _, user := range users {
	// 	ack, err := js.Publish(ctx, "mail.job", nil)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}
