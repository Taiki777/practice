package main

import (
	"flag"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	url := flag.String("s", nats.DefaultURL, "nats server url")
	subject := flag.String("subj", "demo.hello", "subject to publish")
	msg := flag.String("msg", "hello from go", "message")
	flag.Parse()

	nc, err := nats.Connect(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	if err := nc.Publish(*subject, []byte(*msg)); err != nil {
		log.Fatal(err)
	}
	// 送信を確実にフラッシュ
	if err := nc.Flush(); err != nil {
		log.Fatal(err)
	}
}
