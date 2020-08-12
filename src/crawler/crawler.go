package crawler

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func determinEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	handleErr(err)
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func MainRun() {
	const targetUrl = "http://www.zhenai.com/zhenghun"
	resp, err := http.Get(targetUrl)
	handleErr(err)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code:", resp.StatusCode)
		return
	}
	e := determinEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	handleErr(err)
	fmt.Printf("%s\n", all)

}
