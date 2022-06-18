package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

// go routineの走っている数を限定することができる

//何個同時に走らせられるか
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("could not get lock")
		return
	}

	////	contextのなかでaquireしたものがこのプロセスを走らせられる
	//if err := s.Acquire(ctx, 1); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	defer s.Release(1)
	time.Sleep(1 * time.Second)
	fmt.Println("Done")

}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	time.Sleep(5 * time.Second)

}
