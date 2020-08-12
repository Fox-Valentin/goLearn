package httpDemo

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/12.0 Mobile/15A372 Safari/604.1

func handlePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func Demo() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/12.0 Mobile/15A372 Safari/604.1")
	// resp, err := http.Get("http://www.imooc.com")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	handlePanic(err)
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	handlePanic(err)
	fmt.Printf("%s\n", s)
}
