package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	url := flag.String("s", nats.DefaultURL, "nats server url")
	subject := flag.String("subj", "demo.hello", "subject to subscribe")
	flag.Parse()

	nc, err := nats.Connect(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	_, err = nc.Subscribe(*subject, func(m *nats.Msg) {
		fmt.Printf("received: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// サーバに購読を確実に送る
	if err := nc.Flush(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("subscribing on:", *subject)

	select {} // 終了しないように待つ
}
