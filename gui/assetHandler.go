package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AssetHandler struct {
	http.Handler
	config *AppConfig
}

func NewAssetHandler(conf *AppConfig) *AssetHandler {
	return &AssetHandler{config: conf}
}

func (h *AssetHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	target, _ := url.JoinPath(h.config.ApiServer, req.RequestURI)
	fmt.Println("new request: " + target)

	uri := req.RequestURI
	if strings.HasPrefix(uri, "/api") {
		uri = uri[4:]
	}

	target, err := url.JoinPath(h.config.ApiServer, uri)
	if err != nil {
		fmt.Printf("failed to joinpath, %v", err)
		return
	}

	fmt.Println("direct to target: " + target)

	// generate new request
	request, _ := http.NewRequest(req.Method, target, req.Body)
	request.Header = req.Header

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("failed to do request, %v", err)
		return
	}
	defer response.Body.Close()

	w.WriteHeader(response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("failed to read response body, %v", err)
		return
	}

	_, _ = w.Write(body)
}
