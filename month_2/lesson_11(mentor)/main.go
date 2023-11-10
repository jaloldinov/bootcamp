package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// G1()

	// WaitGroup()

	// Channel()

	// ChannelClose()

	// ChannelNil()

	chanRW()

	// BufferedChan()
	// time.Sleep(time.Second)

}

// goroutine
func G1() {
	go fmt.Println("hello")

	fmt.Println("hi")
	fmt.Println("how are you?")
	time.Sleep(time.Second)
}

// WaitGroup
func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("hello")
	}()

	fmt.Println("hi")
	wg.Wait()
}

// Channel
func Channel() {
	ch := make(chan bool)
	go func() {
		fmt.Println("hello", time.Now())
		ch <- true

		ch <- false
		fmt.Println("send 2")
	}()
	fmt.Println("hi")
	time.Sleep(time.Second)
	fmt.Println(<-ch, time.Now())

	fmt.Println(<-ch, time.Now())
	time.Sleep(time.Second)

}
func ChannelClose() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello")
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }
	// for v := range ch {
	// 	fmt.Println(v)
	// }
	for {
		v, open := <-ch
		if !open {
			break
		}
		fmt.Println(v)
	}
	fmt.Println("hi")
}
func ChannelNil() {
	var ch chan bool
	go func() {
		close(ch)  // panic
		ch <- true //block
		fmt.Println("hello")
	}()
	// fmt.Println(<-ch) //block
	fmt.Println("hi")
	for {
	}
}

func chanRW() {
	ch := make(chan int)
	go readChan(ch)
	println("writing")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("num:", runtime.NumCPU())
	// go writeChan(ch)

	// println("reading")
	// for v := range ch {
	// 	fmt.Println(v)
	// }

}
func readChan(ch <-chan int) {
	println("reading")

	for v := range ch {
		fmt.Println(v)
	}
}
func writeChan(ch chan<- int) {
	println("writing")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func BufferedChan() {
	ch := make(chan int, 100)
	go func() {
		fmt.Println("writing time:", time.Now())
		ch <- 1
		fmt.Println("writing2 time:", time.Now())
		ch <- 2
		fmt.Println("writing end time:", time.Now())
		close(ch)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("reading time:", time.Now())
	fmt.Println(<-ch)

	time.Sleep(5 * time.Second)

	fmt.Println("hi")
}
