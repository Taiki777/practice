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

type Task struct {
	To       string `json:"to"`
	Template string `json:"template"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
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
		Name: "worker",
		//AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	// consumerを取得？ 上の作成との違いはまだ不明
	// consumer, err := stream.Consumer(ctx, "worker")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("waiting for messages...")

	//log.Printf("worker started sleep=%dms", *sleepMs)

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

			// var task Task
			// if err := json.Unmarshal(msg.Data(), &task); err != nil {
			// 	log.Printf("Unmarshal error: %v", err)
			// 	msg.Term()
			// 	continue
			// }

			// log.Printf("Processing: to=%s, template=%s", task.To, task.Template)

			var user User
			if err := json.Unmarshal(msg.Data(), &user); err != nil {
				log.Printf("Unmarshal error: %v", err)
				msg.Term()
				continue
			}

			//log.Printf("Processing: to=%s, template=%s", user.ID, user.Email)

			// ここで実際の処理を行う（メール送信など）
			log.Printf("NATS: Sending to %s (%s)...", user.ID, user.Email)
			sendEmails()

			// 処理完了をAck
			if err := msg.Ack(); err != nil {
				log.Printf("Ack error: %v", err)
			}
		}
	}()

	// シグナルを待ってクリーンに終了
	<-ctx.Done()
	log.Printf("Shutting down...")

	// 2) pullで取り出して処理（バッチ10件）
	// for {
	// 	msgs, err := c.Fetch(10, jetstream.FetchMaxWait(2*time.Second))
	// 	if err != nil {
	// 		continue
	// 	}

	// 	for msg := range msgs.Messages() {
	// 		var u User
	// 		if err := json.Unmarshal(msg.Data(), &u); err != nil {
	// 			_ = msg.Ack() // 壊れたpayloadは捨てる（学習用）
	// 			continue
	// 		}

	// 		time.Sleep(time.Duration(*sleepMs) * time.Millisecond)
	// 		log.Printf("sent to %s (%s)", u.ID, u.Email)

	// 		// ★最低限のACK 1行：これがあるから「落としても未完了ジョブが残る」が成立する
	// 		_ = msg.Ack()
	// 	}
	// }
}
