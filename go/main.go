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

type EmailResult struct {
	UserID  string
	Success bool
	Error   error
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

func sendEmails(user User) error {
	time.Sleep(200 * time.Millisecond)

	return nil
}

func sendEmailsSync(users []User) []EmailResult {
	results := make([]EmailResult, 0, len(users))

	for _, user := range users {
		log.Printf("同期: Sending to %s (%s)...", user.ID, user.Email)
		err := sendEmails(user)
		if err != nil {
			log.Printf("同期: Failed: %s - %v", user.ID, err)
		} else {
			log.Printf("同期: Success: %s", user.ID)
		}

		result := EmailResult{
			UserID:  user.ID,
			Success: err == nil,
			Error:   err,
		}
		results = append(results, result)
	}
	return results
}

func sendEmailsGoroutine(users []User, workers int) []EmailResult {
	jobs := make(chan User, len(users))
	resultCh := make(chan EmailResult, len(users))
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range jobs {
				log.Printf("非同期: Sending to %s (%s)...", user.ID, user.Email)
				err := sendEmails(user)
				if err != nil {
					log.Printf("非同期: Failed: %s - %v", user.ID, err)
				} else {
					log.Printf("非同期: Success: %s", user.ID)
				}
				resultCh <- EmailResult{
					UserID:  user.ID,
					Success: err == nil,
					Error:   err,
				}
			}
		}()
	}

	for _, user := range users {
		jobs <- user
	}
	close(jobs)

	wg.Wait()
	close(resultCh)

	results := make([]EmailResult, 0, len(users))
	for result := range resultCh {
		results = append(results, result)
	}

	return results

	// for _, user := range users {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		log.Printf("非同期: Sending to %s (%s)...", user.ID, user.Email)
	// 		err := sendEmails(user)
	// 		if err != nil {
	// 			log.Printf("非同期: Failed: %s - %v", user.ID, err)
	// 		} else {
	// 			log.Printf("非同期: Success: %s", user.ID)
	// 		}

	// 		result := EmailResult{
	// 			UserID:  user.ID,
	// 			Success: err == nil,
	// 			Error:   err,
	// 		}
	// 		resultCh <- result
	// 	}()
	// }
	// wg.Wait()
	// close(resultCh)

	// results := make([]EmailResult, 0, len(users))
	// for result := range resultCh {
	// 	results = append(results, result)
	// }

	// return results
}

func resultSummary(results []EmailResult) {
	successCount := 0
	failCount := 0

	for _, result := range results {
		if result.Success {
			successCount++
		} else {
			failCount++
		}
	}

	fmt.Printf("成功: %d件, 失敗: %d件\n", successCount, failCount)
}

func main() {
	// テスト用ユーザーデータ
	// 選択できるユーザー数と送信時間は変えれてもいい？（優先度: 低）
	users := makeUsers(10)

	startSync := time.Now()
	resultsSync := sendEmailsSync(users)
	durationSync := time.Since(startSync)
	rpsSync := len(users) / int(durationSync.Seconds())

	fmt.Printf("\n同期処理完了: %v, rps: %v\n", durationSync, rpsSync)
	resultSummary(resultsSync)

	const numWorkers = 5
	startAsyncGoroutine := time.Now()
	resultsAsyncGoroutine := sendEmailsGoroutine(users, numWorkers)
	durationAsyncGoroutine := time.Since(startAsyncGoroutine)
	//rpsAsyncGoroutine := len(users) / int(durationAsyncGoroutine.Seconds())

	fmt.Printf("\n非同期処理(Goroutine)完了: %v\n", durationAsyncGoroutine)
	resultSummary(resultsAsyncGoroutine)
}
