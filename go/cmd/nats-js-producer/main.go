package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func ensureStream(js nats.JetStreamContext) error {
	cfg := jetstream.StreamConfig{
		Name:      "Email",
		Subjects:  []string{"mail.job"},
		Retention: jetstream.WorkQueuePolicy,
		Storage:   jetstream.FileStorage,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, _ := js.CreateStream(ctx, cfg)
	fmt.Println("created the stream")
	return err
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
}
