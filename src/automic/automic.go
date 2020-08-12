package automic

import (
	"fmt"
	"sync"
	"time"
)

type automic struct {
	value int
	lock  sync.Mutex
}

func (a *automic) increment() {
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *automic) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func Demo() {
	var a automic
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
