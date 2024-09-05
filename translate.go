package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
)

func TranslateLanguage(sourceLanguage, targetLanguage, text string, args ...interface{}) string {
	if runtime.GOOS != "linux" {
		return text
	}
	if len(sourceLanguage) == 0 {
		sourceLanguage = "auto"
	}
	if len(targetLanguage) == 0 {
		targetLanguage = "pt"
	}
	version := "huoshan"
	if len(args) > 0 && reflect.TypeOf(args[0]).String() == "string" {
		version = args[0].(string)
	}

	uri := "http://translate:8080/api/exec-translate"
	param := make(map[string]string)
	param["content"] = text
	param["source_language"] = sourceLanguage
	param["target_language"] = targetLanguage
	param["version"] = version
	_, response, err := GetRequest(RequestMethodGet, uri, param, nil)
	if err != nil {
		fmt.Println("-----------TranslateLanguage", err)
		return text
	}
	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Content string `json:"content"`
		} `json:"data"`
	}

	var res Response
	json.Unmarshal([]byte(response), &res)
	if res.Code != 200 {
		fmt.Println("-----------TranslateLanguage", response)
		return text
	}

	return res.Data.Content
}
