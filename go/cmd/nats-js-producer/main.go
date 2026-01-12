package main

import (
	"context"
	"encoding/json"
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

type EmailTask struct {
	User          User   `json:"user"`           // 今回はDBなしで完結させたいので丸ごと入れる
	ResultSubject string `json:"result_subject"` // 完了通知の返却先（tasks.results.<runID>）
}

type EmailResult struct {
	UserID string `json:"user_id"`
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
	users := makeUsers(20)
	// 常駐用　ctrl + cなどで止める
	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// defer stop()

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
		Name:      "Tasks",
		Subjects:  []string{"tasks.email"},
		Retention: jetstream.WorkQueuePolicy,
		Storage:   jetstream.FileStorage,
	}

	// 単発用のctx　時間で止まる
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = js.CreateOrUpdateStream(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created the stream")

	runID := fmt.Sprintf("%d", time.Now().UnixNano())
	resultSubject := "tasks.results." + runID

	// resultsSubjectをsubscribe
	sub, err := nc.SubscribeSync(resultSubject)
	if err != nil {
		log.Fatal(err)
	}
	// SubscribeSync の後に Flush しておく
	// 購読がサーバに反映される前に結果が飛ぶと取り逃がすことがあるので、基本入れます。
	// if err := nc.FlushTimeout(2 * time.Second); err != nil {
	// 	log.Fatal(err)
	// }

	startTime := time.Now()
	for _, user := range users {
		emailTask := EmailTask{
			User:          user,
			ResultSubject: resultSubject,
		}
		data, err := json.Marshal(emailTask)
		if err != nil {
			log.Fatal(err)
		}
		ack, err := js.Publish(ctx, "tasks.email", data)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("published stream=%s, seq=%d", ack.Stream, ack.Sequence)
	}

	recieved := 0
	for recieved < len(users) {
		msg, err := sub.NextMsg(60 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
		log.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
		recieved++
	}
	totalTime := time.Since(startTime)
	rps := float64(len(users)) / totalTime.Seconds()
	fmt.Printf("total = %v rps = %v", totalTime, rps)
}
