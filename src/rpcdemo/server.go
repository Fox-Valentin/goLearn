package rpcdemo

import (
	"encoding/json"
	"hashdemo"
	"io"
	"net/http"
)

var bc *hashdemo.BlockChain

func Run() {
	bc = hashdemo.NewBlockChain()
	http.HandleFunc("/blockchain/get", blockChainGetHandler)
	http.HandleFunc("/blockchain/write", blockChainWriteHandler)
	http.ListenAndServe("localhost:8888", nil)
}

func blockChainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(bc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockChainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	bc.SendData(blockData)
	blockChainGetHandler(w, r)
}
