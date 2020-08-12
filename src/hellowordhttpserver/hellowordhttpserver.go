package hellowordhttpserver

import (
	"fmt"
	"net/http"
	"sort"
)

func HelloServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1> hello world! %s</h1>", request.FormValue("name"))
	})
	http.ListenAndServe(":8888", nil)
}

func HelloChan() {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		go func(i int) {
			ch <- fmt.Sprintf("hello world %d\n", i)
		}(i)
	}
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func Sort() {
	a := []int{1, 5, 3, 7, 9, 4, 1, 2, 3, 1}
	sort.Ints(a)
	for _, v := range a {
		fmt.Println(v)
	}
}
