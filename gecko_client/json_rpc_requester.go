package gecko_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// ============= RPC Requester ===================
const (
	JSON_RPC_VERSION = "2.0"
)

// This needs to be public so the JSON package can serialize it
type JsonRpcRequest struct {
	JsonRpc string	`json:"jsonrpc"`
	Method string 	`json:"method"`
	Params map[string]interface{} `json:"params"`
	Id int `json:"id"`
}

type JsonRpcError struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json: "data"`
}

type JsonRpcResponse struct {
	JsonRpcVersion string             `json:"jsonrpc"`
	Error JsonRpcError `json: "error"`
	Result map[string]interface{} `json: "result"`
	Id             int                `json:"id"`
}

type jsonRpcRequester interface {
	makeRpcRequest(endpoint string, method string, params map[string]interface{}) ([]byte, error)
}

type geckoJsonRpcRequester struct {
	ipAddr string
	port nat.Port
	client http.Client
}

func newGeckoJsonRpcRequester(ipAddr string, port nat.Port, requestTimeout time.Duration) *geckoJsonRpcRequester {
	return &geckoJsonRpcRequester{
		ipAddr: ipAddr,
		port:   port,
		client: http.Client{
			Timeout: requestTimeout,
		},
	}
}


func (requester geckoJsonRpcRequester) makeRpcRequest(endpoint string, method string, params map[string]interface{}) ([]byte, error) {
	// Either Golang or Ava have a very nasty & subtle behaviour where duplicated '//' in the URL is treated as GET, even if it's POST
	// https://stackoverflow.com/questions/23463601/why-golang-treats-my-post-request-as-a-get-one
	endpoint = strings.TrimLeft(endpoint, "/")
	request := JsonRpcRequest{
		JsonRpc: JSON_RPC_VERSION,
		Method: method,
		Params:  params,
		Id: 1,
	}

	requestBodyBytes, err := json.Marshal(request)
	if err != nil {
		return nil, stacktrace.Propagate(
			err,
			"Could not marshall request to endpoint '%v' with method '%v' and params '%v' to JSON",
			endpoint,
			method,
			params)
	}

	url := fmt.Sprintf("http://%v:%v/%v", requester.ipAddr, requester.port.Int(), endpoint)

	logrus.Tracef("Making request to url: %v", url)
	logrus.Tracef("Request body: %v", string(requestBodyBytes))
	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Error occurred when making JSON RPC POST request to %v", url)
	}
	defer resp.Body.Close()
	statusCode := resp.StatusCode
	logrus.Tracef("Got response with status code: %v", statusCode)

	responseBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, stacktrace.Propagate(err, "Error occurred when reading response body")
	}
	logrus.Tracef("Response body: %v", string(responseBodyBytes))

	if statusCode != 200 {
		return nil, stacktrace.NewError(
			"Received response with non-200 code '%v' and response body '%v'",
			statusCode,
			string(responseBodyBytes))
	}

	var response JsonRpcResponse
	if err := json.Unmarshal(responseBodyBytes, &response); err != nil {
		return nil, stacktrace.Propagate(err, "Error unmarshalling JSON response")
	}
	if response.Error.Code != 0 {
		return nil, stacktrace.NewError("RPC call failed: %+v", response.Error)
	}
	return responseBodyBytes, nil
}
