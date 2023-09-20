package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"playground/cpp-bootcamp/config"
	"playground/cpp-bootcamp/handler"
	"playground/cpp-bootcamp/storage/db"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	cfg := config.Load()
	strg, err := db.NewStorage(context.Background(), cfg)
	fmt.Println("strg:", strg)
	h := handler.NewHandler(strg)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/person/", h.PersonHandler)

	fmt.Printf("server is running on port %s", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
	if err != nil {
		panic(err)
	}
}
func getAllPerson(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {

		values := req.URL.Query()
		search := values.Get("search")
		limit := values.Get("limit")
		limitN, _ := strconv.Atoi(limit)
		fmt.Println(search)
		fmt.Println(limitN)
		data, err := json.Marshal("persons")
		if err != nil {
			fmt.Println("error:", err.Error())
			return
		}
		rw.Header().Add("content-type", "application/json")
		rw.Write(data)
		rw.WriteHeader(http.StatusOK)
	}
}

//

//

//

//

//

//

//

//

//
//select
func Select() {
	var ch1, ch2 = make(chan interface{}), make(chan interface{})
	// go sendMsg(ch1, false)
	go sendMsg(ch2, true)

	for {
		select {
		case msg, ok := <-ch1:
			if !ok {
				break
			}
			fmt.Println("from ch1:", msg)
		case msg, ok := <-ch2:
			if !ok {
				break
			}
			fmt.Println("from ch2:", msg)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("waiting")
		}
	}
}
func sendMsg(ch chan interface{}, char bool) {
	if char {
		for i := 'a'; i < 'a'+30; i++ {
			ch <- string(i)
			time.Sleep(499 * time.Millisecond)
		}
		close(ch)
	} else {
		for i := 0; i < 30; i++ {
			ch <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(ch)
	}
}

//goroutine
func G1() {
	go fmt.Println("hello")

	fmt.Println("hi")
	fmt.Println("how are you?")
	time.Sleep(time.Second)
}

//WaitGroup
func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("hello")
	}()

	fmt.Println("hi")
	wg.Wait()
}

//Channel
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
func readChan(ch chan int) {
	println("reading")

	for v := range ch {
		fmt.Println(v)
	}
}
func writeChan(ch chan int) {
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
