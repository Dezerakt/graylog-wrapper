package graylog_wrapper

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GraylogConfig struct {
	address string
}

func Init(address string) *GraylogConfig {
	return &GraylogConfig{address: address}
}

func (g *GraylogConfig) WriteLog(message interface{}) error {
	marshalledMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, g.address, bytes.NewBuffer(marshalledMessage))
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
