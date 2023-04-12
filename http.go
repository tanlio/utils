package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func PostRequest(uri string, param map[string]interface{}, header map[string]string, args ...interface{}) (string, error) {
	paramJson, err := json.Marshal(param)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	request, err := http.NewRequest("POST", uri, strings.NewReader(string(paramJson)))
	if request == nil {
		return "", errors.New("build http request error")
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func PostRequest2(uri string, param map[string]string, header map[string]string, args ...interface{}) (string, error) {
	paramJson, err := json.Marshal(param)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	request, err := http.NewRequest("POST", uri, strings.NewReader(string(paramJson)))
	if request == nil {
		return "", errors.New("build http request error")
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func PostRequest3(uri string, param map[string]string, header map[string]string, args ...interface{}) (string, error) {
	data := url.Values{}
	for k, v := range param {
		data.Set(k, v)
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	request, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	if request == nil {
		return "", errors.New("build http request error")
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func PostRequest4(uri string, param map[string]interface{}, header map[string]string, args ...interface{}) (string, error) {
	data := url.Values{}
	for k, v := range param {
		switch reflect.TypeOf(v).String() {
		case "int":
			data.Set(k, strconv.Itoa(v.(int)))
		case "int32":
			data.Set(k, strconv.Itoa(int(v.(int32))))
		case "int64":
			data.Set(k, strconv.FormatInt(v.(int64), 10))
		case "string":
			data.Set(k, v.(string))
		case "float64":
			data.Set(k, decimal.NewFromFloat(v.(float64)).String())
		case "float32":
			data.Set(k, decimal.NewFromFloat(float64(v.(float32))).String())
		case "bool":
			data.Set(k, IfString(v.(bool), "true", "false"))
		default:
			return "", errors.New("Parameter format error")
		}
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	request, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	if request == nil {
		return "", errors.New("build http request error")
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func PostRequest5(uri string, param map[string]string, file *os.File, header map[string]string, args ...interface{}) (string, error) {
	data := url.Values{}
	for k, v := range param {
		data.Set(k, v)
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf)

	for k, v := range param {
		if len(v) == 0 {
			continue
		}
		pw, _ := bw.CreateFormField(k)
		pw.Write([]byte(v))
	}

	//file
	if file != nil {
		var fileName string
		for k, v := range param {
			if len(v) != 0 {
				continue
			}
			fileName = k
			break
		}
		if len(fileName) != 0 {
			fw, _ := bw.CreateFormFile(fileName, file.Name())
			io.Copy(fw, file)
		}
	}

	request, err := http.NewRequest("POST", uri, buf)
	if request == nil {
		return "", errors.New("build http request error")
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "multipart/form-data")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetRequest(uri string, param map[string]string, header map[string]string, args ...interface{}) (string, error) {
	data := url.Values{}
	for k, v := range param {
		data.Set(k, v)
	}
	if len(param) > 0 {
		uri += "?"
		uri += data.Encode()
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	req, _ := http.NewRequest("GET", uri, nil)
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body), err
}
