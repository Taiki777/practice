package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
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

func sendEmails() error {
	time.Sleep(200 * time.Millisecond)

	return nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	// Stream取得（Producerが先に作ってても良いし、ここでCreateStreamしても良い）
	stream, err := js.Stream(ctx, "Tasks")
	if err != nil {
		log.Fatal(err)
	}

	// 1) Consumer作成 or 更新（durable pull）
	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:      "worker",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("waiting for messages...")

	// メッセージを取得するループ
	go func() {
		iter, _ := cons.Messages()
		defer iter.Stop()

		for {
			msg, err := iter.Next()
			if err != nil {
				if errors.Is(err, jetstream.ErrMsgIteratorClosed) {
					break
				}
				log.Printf("Iterator error: %v", err)
				continue
			}

			var emailTask EmailTask
			if err := json.Unmarshal(msg.Data(), &emailTask); err != nil {
				log.Printf("Unmarshal error: %v", err)
				msg.Term()
				continue
			}

			// ここで実際の処理を行う（メール送信など）
			log.Printf("NATS: Sending to %s (%s)...", emailTask.User.ID, emailTask.User.Email)
			sendEmails()

			// resultをpublish
			res := EmailResult{
				UserID: emailTask.User.ID,
			}
			data, err := json.Marshal(res)
			if err != nil {
				log.Fatal(err)
			}
			nc.Publish(emailTask.ResultSubject, data)

			// 処理完了をAck
			if err := msg.Ack(); err != nil {
				log.Printf("Ack error: %v", err)
			}
		}
	}()

	// シグナルを待ってクリーンに終了
	<-ctx.Done()
	log.Printf("Shutting down...")
}
