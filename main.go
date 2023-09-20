package graylog_wrapper

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/color"
	"log"
	"net/http"
)

var graylogObject graylogConfig

type graylogConfig struct {
	address    string
	Message    string      `json:"message"`
	Host       string      `json:"host"`
	MethodName string      `json:"method_name"`
	Body       interface{} `json:"body"`
}

func Init(address string) {
	graylogObject = graylogConfig{
		address: address,
		Host:    "localhost",
	}
}

func WriteLog(methodName string, body interface{}, stage string) {
	graylogObject.Message = stage
	graylogObject.Body = body
	graylogObject.MethodName = methodName

	marshalledMessage, err := json.Marshal(graylogObject)
	if err != nil {
		log.Println(color.RedString(err.Error()))
	}

	request, err := http.NewRequest(http.MethodPost, graylogObject.address, bytes.NewBuffer(marshalledMessage))
	if err != nil {
		log.Println(color.RedString(err.Error()))
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		log.Println(color.RedString(err.Error()))
	}
}
