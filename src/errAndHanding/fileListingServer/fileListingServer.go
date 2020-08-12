package fileListingServer

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start " + "with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if userErr, ok := err.(userError); ok {
		http.Error(writer, userErr.Message(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
func HttpHandleFn() {
	http.HandleFunc("/", errWrapper(HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

type appHandler func(writer http.ResponseWriter, requset *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, requset *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic %v", r)
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
