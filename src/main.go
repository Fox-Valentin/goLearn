package main

import (
	"bufio"
	"bytes"
	"learngo/src/crawler"

	// "crawler"
	"fmt"
	"io"
	"learngo/src/interface_demo"
	"learngo/src/retriever/mock"
	"learngo/src/retriever/real"
	"learngo/src/retrievers"
	"learngo/src/retrievers/mocks"
	"learngo/src/retrievers/reals"
	"learngo/src/tree"
	"math"
	"math/cmplx"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

const url = "http://www.imooc.com"

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "cczk",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}
func inspect(r Retriever) {
	fmt.Printf("%T  %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(v.Contents)
	case *real.Retriever:
		fmt.Println(v.TimeOut)
	}
}
func log(s interface{}) {
	fmt.Println(s)
}

type Person struct {
	Name    string `xml:"NAME,attr"`
	Gender  string
	Age     uint8
	Address string
}

func (p *Person) Move(addr string) string {
	p.Address, addr = addr, p.Address
	return addr
}
func readFile(path string) ([]byte, error) {
	parentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fullpath := filepath.Join(parentPath, path)
	file, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return read(file)
}
func read(r io.Reader) ([]byte, error) {
	br := bufio.NewReader(r)
	var buf bytes.Buffer
	for {
		line, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		buf.Write(line)
		if !isPrefix {
			buf.WriteByte('\n')
		}
	}
	return buf.Bytes(), nil
}
func makeSlice() {
	mSlice := make([]string, 3)
	mSlice[0] = "a"
	mSlice[1] = "b"
	mSlice[2] = "c"
	fmt.Println(mSlice)
}
func makeMap() {
	mMap := make(map[int]string)
	mMap[1] = "a"
	mMap[2] = "b"
	mMap[3] = "c"
	fmt.Println(mMap)
}
func makeChan() {
	mChan := make(chan int)
	close(mChan)
}
func newMap() {
	newMap := new(map[int]string)
	makeMap := make(map[int]string)
	fmt.Println("newMap:", reflect.TypeOf(newMap))
	fmt.Println("makeMap:", reflect.TypeOf(makeMap))
}
func appendSlice() {
	mSlice := make([]string, 2)
	mSlice[0] = "a"
	mSlice[1] = "b"
	fmt.Println("len of mSlice is:", len(mSlice))
	fmt.Println("cap of mSlice is:", cap(mSlice))
	mSlice = append(mSlice, "c")
	fmt.Println("len of mSlice is:", len(mSlice))
	fmt.Println("cap of mSlice is:", cap(mSlice))
}
func copySlice() {
	srcSlice := make([]string, 2)
	distSlice := make([]string, 3)
	srcSlice[0] = "a"
	srcSlice[1] = "b"
	distSlice[0] = "a1"
	distSlice[1] = "a2"
	distSlice[2] = "a3"
	copy(distSlice, srcSlice)
	fmt.Println("distSlice", distSlice)
	fmt.Println("len of distSlcie", len(distSlice))
}
func deleteMap() {
	mMap := make(map[int]string)
	mMap[1] = "a"
	mMap[2] = "b"
	delete(mMap, 1)
	fmt.Println(mMap)
}
func recoverFn() {
	p := recover()
	if p != nil {
		switch p.(type) {
		case string:
			fmt.Printf("this is a string type panic: %s", p)
		case error:
			fmt.Printf("this is a error type panic: %v", p)
		default:
			fmt.Println("unknow panic")
		}
	}
}
func makePanic() {
	defer recoverFn()
	panic(1)
}
func closeChan() {
	mchan := make(chan int, 1)
	close(mchan)
	mchan <- 1
}
func action(b interface_demo.Behavior) string {
	return b.Run()
}

func euler() {
	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
}
func enums() {
	const (
		app = iota
		app1
	)
	fmt.Println(app, app1)
}
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func apply(fn func(int, int) int, a, b int) int {
	p := reflect.ValueOf(fn).Pointer()
	fnName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with params %d, %d", fnName, a, b)
	return fn(a, b)
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
func pringAry(ary *[1]int) {
	ary[0] = 100
	for i, v := range ary {
		fmt.Println(i, v)
	}
}
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, v := range []rune(s) {
		if lastI, ok := lastOccurred[v]; ok && lastI >= start {
			start = lastI + i
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[v] = i
	}
	return maxLength
}
func inspects(r retrievers.Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mocks.MockRetrievers:
		fmt.Println("Contents: ", v.Contents)
	case reals.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
func main() {
	crawler.MainRun()
	//regexpdemo.Demo()
}
