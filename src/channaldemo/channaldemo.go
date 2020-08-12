package channaldemo

import (
	"fmt"
	"sync"
)

func DoWorker(i int, w *worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %d;\n", i, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func CreateWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go DoWorker(id, &w)
	return w
}

func ChanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)
	for i := range workers {
		workers[i] = CreateWorker(i, &wg)
	}
	for i, worker := range workers {
		worker.in <- i
	}
	for i, worker := range workers {
		worker.in <- i + 1
	}
	wg.Wait()
}
