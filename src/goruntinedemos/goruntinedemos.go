package goruntinedemos

import (
	"fmt"
	"time"
)

func worker(i int, c <-chan int) {
	for {
		if n, ok := <-c; ok {
			fmt.Printf("Worker %d received %d;\n", i, n)
		} else {
			break
		}
	}
}

func CreateWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func ChanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}
	time.Sleep(time.Millisecond)
}

func BufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 3
	close(c)
	time.Sleep(time.Millisecond)
}
