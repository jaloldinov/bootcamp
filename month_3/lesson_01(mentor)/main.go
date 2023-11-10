package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	withCancel()
}

func withCancel() {
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// cancel()
	// go func() {
	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	fmt.Println("men bozorga kettim!")

	timeConsumingFunc(ctx, 2*time.Second, "Qaytib kelaver, xech narsa kerakmas")

	defer cancel()
	fmt.Println("keldim")
}

func timeConsumingFunc(ctx context.Context, duration time.Duration, s string) {
	for {
		select {
		case <-time.After(duration):
			fmt.Println(s)
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			return
		}
	}
}
