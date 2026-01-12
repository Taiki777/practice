package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

type User struct {
	ID    string
	Email string
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

func sendEmails() error {
	time.Sleep(200 * time.Millisecond)

	return nil
}

func sendEmailsSync(users []User) {
	for _, user := range users {
		log.Printf("同期: Sending to %s (%s)...", user.ID, user.Email)
		sendEmails()
		log.Printf("同期: Done: %s", user.ID)
	}
}

func sendEmailsGoroutine(users []User, workers int) {
	jobs := make(chan User, len(users))
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range jobs {
				log.Printf("非同期: Sending to %s (%s)...", user.ID, user.Email)
				sendEmails()
				log.Printf("非同期: Done: %s", user.ID)
			}
		}()
	}

	for _, user := range users {
		jobs <- user
	}
	close(jobs)

	wg.Wait()
}

func main() {
	// テスト用ユーザーデータ
	users := makeUsers(100)

	// RPSとは
	// RPS = Requests Per Second（1秒あたりの処理件数）です。
	// 今回なら「1秒あたり何通処理できたか」。
	// 計算はこれだけ：
	// RPS = 件数 / 経過秒

	startSync := time.Now()
	sendEmailsSync(users)
	totalTimeSync := time.Since(startSync)
	rpsSync := float64(len(users)) / totalTimeSync.Seconds()
	fmt.Printf("\n同期処理完了: %v, rps: %v\n", totalTimeSync, rpsSync)

	const numWorkers = 10

	startAsyncGoroutine := time.Now()
	sendEmailsGoroutine(users, numWorkers)
	totalTimeAsyncGoroutine := time.Since(startAsyncGoroutine)
	rpsAsyncGoroutine := float64(len(users)) / totalTimeAsyncGoroutine.Seconds()
	fmt.Printf("\n非同期処理(Goroutine)完了: %v, rps: %v\n", totalTimeAsyncGoroutine, rpsAsyncGoroutine)
}
