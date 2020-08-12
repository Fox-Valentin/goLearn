package errAndHanding

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanicFn(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testUserError string

func (e testUserError) Error() string {
	return e.Message()
}
func (e testUserError) Message() string {
	return string(e)
}

func errUserErr(writer http.ResponseWriter, request *http.Request) error {
	return testUserError("user error")
}

func errorNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func errorNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}
func errorUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

var testss = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanicFn, 500, "Internal Server Error"},
	{errUserErr, 400, "user error"},
	{errorNotFound, 404, "Not Found"},
	{errorNoPermission, 403, "Forbidden"},
	{errorUnknown, 500, "Internal Server Error"},
}

func TestErrWrap(t *testing.T) {
	for _, tt := range testss {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		f(response, request)
		verifyResponseFn(response.Result(), tt.code, tt.message, t)
	}
}

func verifyResponseFn(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("got (%d, %s); expected (%d %s);", resp.StatusCode, body, expectedCode, expectedMsg)
	}
}
func TestErrByServer(t *testing.T) {
	for _, tt := range testss {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponseFn(resp, tt.code, tt.message, t)
	}
}
