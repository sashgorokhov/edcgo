package main

import (
	"net/http"
	"os"
	"bytes"
	"strings"
	"io/ioutil"
	"io"
)


const RELAY_URL = "https://httpbin.org/post"


func getRelayUrl() string {
	url, found := os.LookupEnv("RELAY_URL")
	if found && url != "" {
		return url
	}
	return RELAY_URL
}

func makePostRequest(url string, body io.Reader) string {
	logger := Logger.WithField("url", url)
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		logger.Errorln(err)
		return ""
	}
	defer resp.Body.Close()
	response_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorln(err)
		return ""
	}
	return string(response_body)
}

func pushLines(lines *[]string, token string) {
	url := getRelayUrl()
	buff := bytes.NewBufferString("")
	buff.WriteString("[")
	buff.WriteString(strings.Join(*lines, ","))
	buff.WriteString("]")
	relay_response := makePostRequest(url, buff)
	Logger.Debugln("Relay response:", relay_response)
}
