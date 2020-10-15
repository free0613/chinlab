package main

import (
	"fmt"
	"time"
)

type ChanPool struct {
	work chan func()
	sem  chan struct{}
}

func NewPool(size int) *ChanPool {
	return &ChanPool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *ChanPool) NewTask(task func()) {

	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)

	}
}

func (p *ChanPool) worker(task func()) {
	defer func() { <-p.sem }()
	for {
		task()
		task = <-p.work
	}
}

func main() {
	input := make(chan int)
	pool := NewPool(5)

	go func() {
		for i := 1; i < 1000; i++ {
			input <- i
		}
	}()
	p := func() {
		for {
			select {
			case n := <-input:
				fmt.Println(n)
			}
		}
	}
	for m := 1; m < 5; m++ {
		pool.NewTask(p)
	}

	time.Sleep(10 * time.Second)

}
