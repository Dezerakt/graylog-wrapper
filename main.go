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
	address     string
	Message     string      `json:"message"`
	Host        string      `json:"host"`
	MethodName  string      `json:"method_name"`
	Body        interface{} `json:"body"`
	SessionUUID string      `json:"session_uuid"`
	PartnerId   uint        `json:"partner_id"`
}

func Init(address string) {
	graylogObject = graylogConfig{
		address: address,
		Host:    "localhost",
	}
}

func SetPartnerId(partnerId uint) {
	graylogObject.PartnerId = partnerId
}

func SetSessionUUID(sessionUUID string) {
	graylogObject.SessionUUID = sessionUUID
}

func WriteLog(methodName string, body interface{}, stage string) int {
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
	response, err := client.Do(request)
	if err != nil {
		log.Println(color.RedString(err.Error()))
	}

	if response != nil {
		return response.StatusCode
	} else {
		return 500
	}
}
