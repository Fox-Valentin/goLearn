package errAndHanding

import (
	"bufio"
	"errors"
	"fibo"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func log(s interface{}) {
	fmt.Println(s)
}
func TryDefer() {
	for i := 0; i < 100; i++ {
		defer log(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func WriteFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibo.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func errPanic(_ http.ResponseWriter, _ *http.Request) error {
	panic(123)
}

type testingUserError string

type appHandler func(writer http.ResponseWriter, requset *http.Request) error
type userError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, requset *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("panic %v", r)
				http.Error(
					writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}()
		err := handler(writer, requset)
		if err != nil {
			if userError, ok := err.(userError); ok {
				http.Error(
					writer,
					userError.Message(),
					http.StatusBadRequest,
				)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(_ http.ResponseWriter, _ *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrPermission
}
func errNoPermission(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrPermission
}
func errUnknown(_ http.ResponseWriter, _ *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, _ *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expected (%d, %s); "+"got (%d %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
