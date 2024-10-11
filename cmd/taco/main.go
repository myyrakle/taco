package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func handleProxy(responseWriter http.ResponseWriter, request *http.Request) {
	// 1. 프록시로 호출할 최종 목적지 Host 획득
	proxyHost := request.Header.Get("Proxy-Host")
	request.Header.Del("Proxy-Host")

	if proxyHost == "" {
		http.Error(responseWriter, "Proxy-Host Header is required", http.StatusBadRequest)
		return
	}

	url, err := url.Parse(proxyHost)
	if err != nil {
		http.Error(responseWriter, "Invalid Proxy-Host Header", http.StatusBadRequest)
		return
	}

	// 2. 프록시 요청 값 생성
	// - path 및 query는 request를 그대로 사용
	// - 프로토콜 스킴과 host만을 proxy host에서 획득한 값으로 변경
	request.URL.Scheme = url.Scheme
	request.URL.Host = url.Host
	request.RequestURI = ""

	proxyRequest, err := http.NewRequest(request.Method, request.URL.String(), request.Body)
	if err != nil {
		http.Error(responseWriter, "Server Error", http.StatusInternalServerError)
		return
	}
	proxyRequest.Header = request.Header

	// 3. 최종 목적지 서버로 요청 전달
	client := &http.Client{}
	response, err := client.Do(proxyRequest)
	if err != nil {
		http.Error(responseWriter, "Server Error", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// 4. 받아온 값을 클라이언트에게 전달
	responseWriter.WriteHeader(response.StatusCode)

	for key, values := range response.Header {
		for _, value := range values {
			responseWriter.Header().Add(key, value)
		}
	}

	_, err = io.Copy(responseWriter, response.Body)
	if err != nil {
		fmt.Println("io.Copy error", err)
	}
}

func main() {
	http.HandleFunc("/", handleProxy)

	fmt.Println("Proxy Server Start")

	// 웹 서버 실행
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
