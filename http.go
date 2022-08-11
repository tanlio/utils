package utils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"net/url"
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
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
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
	body, err := ioutil.ReadAll(resp.Body)
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
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
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
	body, err := ioutil.ReadAll(resp.Body)
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
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
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
	body, err := ioutil.ReadAll(resp.Body)
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
		case "int64":
			data.Set(k, strconv.FormatInt(v.(int64), 10))
		case "string":
			data.Set(k, v.(string))
		case "float64":
			data.Set(k, decimal.NewFromFloat(v.(float64)).String())
		case "float32":
			data.Set(k, decimal.NewFromFloat(float64(v.(float32))).String())
		default:
			return "", errors.New("Parameter format error")
		}
	}

	client := &http.Client{}
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "bool" && args[0].(bool) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
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
	body, err := ioutil.ReadAll(resp.Body)
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
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client.Transport = tr
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
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}
