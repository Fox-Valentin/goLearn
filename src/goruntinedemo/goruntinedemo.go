package goruntinedemo

import (
	"fmt"
	"sync"
	"time"
)

var chanInt = make(chan int, 10)
var timeout = make(chan bool)
var GW sync.WaitGroup

func LoopFn() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(i)
	}
}
func Sender() {
	time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
	time.Sleep(time.Second * 2)
	timeout <- true
}
func Receiver() {
	// num := <-chanInt
	// fmt.Printf("num = %d\n", num)
	// num = <-chanInt
	// fmt.Printf("num = %d\n", num)
	// num = <-chanInt
	// fmt.Printf("num = %d\n", num)
	// <-timeout
	for {
		select {
		case num := <-chanInt:
			fmt.Printf("num = %d", num)
		case <-timeout:
			fmt.Println("timeout...")
		}
	}
}

func Read() {
	for i := 0; i < 10; i++ {
		GW.Add(1)
	}
}

func Write() {
	for i := 0; i < 3; i++ {
		fmt.Printf("Done => %d", i)
		time.Sleep(time.Second * 2)
		GW.Done()
	}
}
