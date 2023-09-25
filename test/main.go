package main

import graylog_wrapper "github.com/Dezerakt/graylog-wrapper"

func main() {
	graylog_wrapper.Init("http://0.0.0.0:12201/gelf")
	graylog_wrapper.SetPartnerId(1231312)
	graylog_wrapper.SetSessionUUID("757d84b0")
	graylog_wrapper.WriteLog("methodNameTest", map[string]interface{}{
		"result": "hello",
	}, "TEST")
}
