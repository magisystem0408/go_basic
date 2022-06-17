package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}
func main() {
	ch := make(chan string)
	//タイムアウトの設定などに使える
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	//空のcontextを使用したい時、TODO()
	//ctx := context.TODO()

	defer cancel()
	go longProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("###########")
}
