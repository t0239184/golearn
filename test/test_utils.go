package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/t0239184/golearn/internal/database"
	"github.com/t0239184/golearn/internal/router"
)

const (
	GET  = "GET"
	POST = "POST"
)

func Get(url *string, data *[]byte) *httptest.ResponseRecorder {
	request := createTestRequest(GET, url, data)
	return sendRequest(request)
}

func Post(url *string, data *[]byte) *httptest.ResponseRecorder {
	request := createTestRequest(POST, url, data)
	return sendRequest(request)
}

func ConvertResponseToMap(response *httptest.ResponseRecorder) (map[string]interface{}, error) {
	var jsonObj map[string]interface{}
	if err := json.Unmarshal(response.Body.Bytes(), &jsonObj); err != nil {
		return nil, err	
	}
	return jsonObj, nil
}

func createTestRequest(method string, url *string, body *[]byte) *http.Request {
	var req *http.Request
	if body != nil {
		req, _ = http.NewRequest(method, *url, bytes.NewReader(*body))
	} else {
		req, _ = http.NewRequest(method, *url, nil)
	}
	return req
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	// 建立一個ResponseRecorder其實作http.ResponseWriter，用來記錄response狀態
	responseRecorder := httptest.NewRecorder()

	db := database.InitDatabase() //之後可以更換成Mock的
	server := router.New(db)

	// gin.Engine.ServerHttp實作http.Handler介面，用來處理HTTP請求及回應。
	server.ServeHTTP(responseRecorder, request)
	return responseRecorder
}
