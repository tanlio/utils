package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	defaultHTTPClient = &http.Client{
		Timeout: 30 * time.Second,
	}

	insecureHTTPClient = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			IdleConnTimeout: 30 * time.Second,
		},
	}
)

const defaultUserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 18_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.5 Mobile/15E148 Safari/604.1"

func pickHTTPClient(args ...any) *http.Client {
	if len(args) > 0 {
		if v, ok := args[0].(bool); ok && v {
			return insecureHTTPClient
		}
	}

	return defaultHTTPClient
}

func RequestForm(method, uri string, paramJson []byte, header map[string]string, args ...any) (int, []byte, error) {
	client := pickHTTPClient(args...)

	var m map[string]interface{}
	json.Unmarshal(paramJson, &m)

	values := url.Values{}
	for k, v := range m {
		values.Set(k, fmt.Sprint(v))
	}

	req, err := http.NewRequest(method, uri, strings.NewReader(values.Encode()))
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("User-Agent", defaultUserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return resp.StatusCode, body, err
}

func RequestJson(method, uri string, paramJson []byte, header map[string]string, args ...any) (int, []byte, error) {
	client := pickHTTPClient(args...)

	req, err := http.NewRequest(method, uri, bytes.NewReader(paramJson))
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("User-Agent", defaultUserAgent)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return resp.StatusCode, body, err
}

func RequestFile(method, uri string, param map[string]string, fileFieldName string, file *os.File, header map[string]string, args ...any) (int, []byte, error) {
	client := pickHTTPClient(args...)

	buf := &bytes.Buffer{}
	bw := multipart.NewWriter(buf)

	for k, v := range param {
		if err := bw.WriteField(k, v); err != nil {
			return 0, nil, err
		}
	}

	if file != nil && fileFieldName != "" {
		fw, err := bw.CreateFormFile(fileFieldName, filepath.Base(file.Name()))
		if err != nil {
			return 0, nil, err
		}
		if _, err = io.Copy(fw, file); err != nil {
			return 0, nil, err
		}
	}

	if err := bw.Close(); err != nil {
		return 0, nil, err
	}

	req, err := http.NewRequest(method, uri, buf)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("User-Agent", defaultUserAgent)
	req.Header.Set("Content-Type", bw.FormDataContentType())
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return resp.StatusCode, body, err
}

func RequestGet(method, rawURL string, param, header map[string]string, args ...any) (int, []byte, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return 0, nil, err
	}

	if len(param) > 0 {
		q := u.Query()
		for k, v := range param {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	client := pickHTTPClient(args...)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return 0, nil, err
	}

	req.Header.Set("User-Agent", defaultUserAgent)
	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return resp.StatusCode, body, err
}
