package json_demo

import (
	"encoding/json"
	"fmt"
)

type Service struct {
	ServiceName  string `json:"name"`
	ServiceIP    string `json:"ip"`
	ServicePoint int    `json:"point"`
}

func TestJsonStruct() {
	service := new(Service)
	service.ServiceName = "json-for-struct"
	service.ServiceIP = "192.167.12.1"
	service.ServicePoint = 8888
	j, err := json.Marshal(service)
	if err != nil {
		fmt.Println("json marshal error:", err.Error())
		return
	}
	fmt.Println("struct json marshal :", string(j))
}

func TestJsonMap() {
	service := make(map[string]interface{})
	service["ServiceName"] = "json-for-map"
	service["ServiceIP"] = "128.12.1.1"
	service["ServicePoint"] = 7777
	j, err := json.Marshal(service)
	if err != nil {
		fmt.Println("json for map err:", err.Error())
		return
	}
	fmt.Println("json for map:", string(j))
}

func TestDeserializeStruct() {
	jsonString := `{"name":"json-for-struct","ip":"192.167.12.1","point":8888}`
	server := new(Service)
	err := json.Unmarshal([]byte(jsonString), &server)
	if err != nil {
		fmt.Println("Deserialize struct err:", err.Error())
		return
	}
	fmt.Println("Deserialize strunt:", server)
}

func TestDeserializeMap() {
	jsonString := `{"name":"json-for-struct","ip":"192.167.12.1","point":8888}`
	service := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &service)
	if err != nil {
		fmt.Println("Deserialize map err: ", err.Error())
		return
	}
	fmt.Println("Deserialize map:", service)
}

func TestSugar(values ...string) {
	for _, v := range values {
		fmt.Printf("value : %s\n", v)
	}
}
