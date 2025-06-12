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
	"time"
)

func Values2Map(body []byte) (map[string]string, []byte) {
	values, _ := url.ParseQuery(string(body))
	parsedData := make(map[string]string)
	for key, v := range values {
		if len(v) > 0 {
			parsedData[key] = v[0]
		}
	}
	data, _ := json.Marshal(parsedData)

	return parsedData, data
}

type RequestMethod string

const (
	RequestMethodGet    RequestMethod = "GET"    //get
	RequestMethodPost   RequestMethod = "POST"   //post
	RequestMethodPut    RequestMethod = "PUT"    //put
	RequestMethodDelete RequestMethod = "DELETE" //delete
)

func PostRequest(method RequestMethod, uri string, param map[string]interface{}, header map[string]string, args ...interface{}) (int, string, error) {
	paramJson, err := json.Marshal(param)
	if err != nil {
		return 0, "", err
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			IdleConnTimeout: 30 * time.Second,
		}
	}

	request, err := http.NewRequest(string(method), uri, strings.NewReader(string(paramJson)))
	if request == nil {
		return 0, "", err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(body), nil
}

func PostRequest2(method RequestMethod, uri string, param map[string]string, header map[string]string, args ...interface{}) (int, string, error) {
	paramJson, err := json.Marshal(param)
	if err != nil {
		return 0, "", err
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			IdleConnTimeout: 30 * time.Second,
		}
	}

	request, err := http.NewRequest(string(method), uri, strings.NewReader(string(paramJson)))
	if request == nil {
		return 0, "", err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(body), nil
}

func PostRequest3(method RequestMethod, uri string, param map[string]string, header map[string]string, args ...interface{}) (int, string, error) {
	data := url.Values{}
	for k, v := range param {
		data.Set(k, v)
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			IdleConnTimeout: 30 * time.Second,
		}
	}

	request, err := http.NewRequest(string(method), uri, strings.NewReader(data.Encode()))
	if request == nil {
		return 0, "", err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(body), nil
}

func PostRequest4(method RequestMethod, uri string, param map[string]interface{}, header map[string]string, args ...interface{}) (int, string, error) {
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
			return 0, "", errors.New("Parameter format error")
		}
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			IdleConnTimeout: 30 * time.Second,
		}
	}

	request, err := http.NewRequest(string(method), uri, strings.NewReader(data.Encode()))
	if request == nil {
		return 0, "", err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(body), nil
}

func PostRequest5(method RequestMethod, uri string, param map[string]string, file *os.File, header map[string]string, args ...interface{}) (int, string, error) {
	data := url.Values{}
	for k, v := range param {
		data.Set(k, v)
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			IdleConnTimeout: 30 * time.Second,
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

	request, err := http.NewRequest(string(method), uri, buf)
	if request == nil {
		return 0, "", err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "multipart/form-data")
	resp, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(body), nil
}

func PostRequest6(method RequestMethod, uri string, paramData []byte, header map[string]string, args ...interface{}) (int, string, error) {
	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			IdleConnTimeout: 30 * time.Second,
		}
	}

	request, err := http.NewRequest(string(method), uri, strings.NewReader(string(paramData)))
	if request == nil {
		return 0, "", err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(body), nil
}

func GetRequest(method RequestMethod, uri string, param map[string]string, header map[string]string, args ...interface{}) (int, string, error) {
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
			IdleConnTimeout: 30 * time.Second,
		}
	}

	req, _ := http.NewRequest(string(method), uri, nil)
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return resp.StatusCode, string(body), err
}
