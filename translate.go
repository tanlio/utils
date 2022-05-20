package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func TranslateLanguage(sourceLanguage, targetLanguage, text string) string {
	if len(sourceLanguage) == 0 {
		sourceLanguage = "auto"
	}
	if len(targetLanguage) == 0 {
		targetLanguage = "pt"
	}

	uri := "http://translate.sampsong.com/api/exec-translate?content=%s&source_language=%s&target_language=%s"
	uri = fmt.Sprintf(uri, url.QueryEscape(text), sourceLanguage, targetLanguage)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	resp, err := client.Do(req)
	if err != nil {
		return text
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("-----------TranslateLanguage", string(body))
		return text
	}
	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Content string `json:"content"`
		} `json:"data"`
	}

	var response Response
	json.Unmarshal(body, &response)
	if response.Code != 200 {
		fmt.Println("-----------TranslateLanguage", string(body))
		return text
	}

	return response.Data.Content
}
