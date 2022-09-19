package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func doThingThatShouldBeLimited() string {
	time.Sleep(1 * time.Second)
	return "done"
}

func main() {
	pg := New(2)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
func main3() {
	//main0()
	//main2()
	//ch, cancel := countTo(10)
	//for i := range ch {
	//	if i > 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//defer cancel()

	gauge := New(1)
	err := gauge.Process(func() {})
	if err != nil {
		fmt.Println("??", err)
	}
}

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:

			}
		}
		close(ch)
	}()
	return ch, cancel
}
func main0() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "additional goroutine"
		v2 := <-ch2
		fmt.Println("go func", v2)
	}()
	v := "main goroutine"
	var v2 string
	for i := 0; i < 2; i++ {
		select {
		case ch2 <- v:
		case v2 = <-ch1:
		}
	}
	fmt.Println("exit", v, v2)
}
func main2() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func(v int) {
			ch <- v * 2
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
