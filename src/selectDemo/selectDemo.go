package selectDemo

import (
	"fmt"
	"math/rand"
	"time"
)

func Demo() {
	c1, c2 := generator(), generator()
	worker := CreateWorker(0)
	n := 0
	var values []int
	tm := time.After(10 * time.Second)
	tik := time.Tick(time.Second)
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tik:
			fmt.Printf("values length is %d\n", len(values))
		case <-tm:
			fmt.Println("bye")
		}
	}
}

func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}
func worker(i int, c <-chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d;\n", i, n)
	}
}
func CreateWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}
