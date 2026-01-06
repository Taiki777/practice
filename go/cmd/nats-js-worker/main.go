package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func sendEmails() error {
	time.Sleep(200 * time.Millisecond)

	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
	stream, err := js.Stream(ctx, "mail.job")
	if err != nil {
		log.Fatal(err)
	}

	// 1) Consumer作成 or 更新（durable pull）
	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:      "CONSUMER",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("worker started sleep=%dms", *sleepMs)

	// 2) pullで取り出して処理（バッチ10件）
	for {
		msgs, err := c.Fetch(10, jetstream.FetchMaxWait(2*time.Second))
		if err != nil {
			continue
		}

		for msg := range msgs.Messages() {
			var u User
			if err := json.Unmarshal(msg.Data(), &u); err != nil {
				_ = msg.Ack() // 壊れたpayloadは捨てる（学習用）
				continue
			}

			time.Sleep(time.Duration(*sleepMs) * time.Millisecond)
			log.Printf("sent to %s (%s)", u.ID, u.Email)

			// ★最低限のACK 1行：これがあるから「落としても未完了ジョブが残る」が成立する
			_ = msg.Ack()
		}
	}
}
