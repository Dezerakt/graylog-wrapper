package graylog_wrapper

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/color"
	"log"
	"net/http"
)

type GraylogConfig struct {
	address string
	message string
	host    string
	body    interface{}
}

func Init(address string) *GraylogConfig {
	return &GraylogConfig{
		address: address,
		message: "empty",
		host:    "localhost",
	}
}

func (g *GraylogConfig) WriteLog(body interface{}, shortMessage string) {
	g.body = body
	if shortMessage != "" {
		g.message = shortMessage
	}

	marshalledMessage, err := json.Marshal(g)
	if err != nil {
		log.Println(color.RedString(err.Error()))
	}

	request, err := http.NewRequest(http.MethodPost, g.address, bytes.NewBuffer(marshalledMessage))
	if err != nil {
		log.Println(color.RedString(err.Error()))
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		log.Println(color.RedString(err.Error()))
	}

}
