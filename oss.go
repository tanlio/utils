package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"runtime"
)

type AllSign struct {
}

var AllSignObject AllSign

type OssSignResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Key         string `json:"key"`           //上传到Bucket内的Object的完整路径
		Policy      string `json:"policy"`        //上传的策略（Policy），Policy为经过Base64编码过的字符串
		AccessKeyId string `json:"access_key_id"` //用户请求的AccessKey ID
		Signature   string `json:"signature"`     //对Policy签名后的字符串
		FileURL     string `json:"file_url"`      //文件地址
		UploadURL   string `json:"upload_url"`    //上传地址
	} `json:"data"`
}

func (AllSign) OssSign(folderName, suffix string, args ...any) (OssSignResponse, error) {
	uri := IfString(runtime.GOOS == "linux", "http://oss:8080", "http://127.0.0.1:28080")
	uri += "/api/oss/oss-sign"

	isAccelerate := false
	if len(args) > 0 {
		if val, ok := args[0].(bool); ok && val {
			isAccelerate = true
		}
	}

	param := make(map[string]any)
	param["folder_name"] = folderName
	param["suffix"] = suffix
	param["is_accelerate"] = isAccelerate
	_, response, err := DefaultClient().RequestGet(context.Background(), http.MethodGet, uri, param, nil)
	if err != nil {
		return OssSignResponse{}, err
	}

	var ossSignResponse OssSignResponse
	json.Unmarshal(response, &ossSignResponse)

	return ossSignResponse, err
}

type OssSignNewResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Key         string `json:"key"`           //上传到Bucket内的Object的完整路径
		Policy      string `json:"policy"`        //上传的策略（Policy），Policy为经过Base64编码过的字符串
		AccessKeyId string `json:"access_key_id"` //用户请求的AccessKey ID
		Signature   string `json:"signature"`     //对Policy签名后的字符串
		FileURL     string `json:"file_url"`      //文件地址
		UploadURL   string `json:"upload_url"`    //上传地址
	} `json:"data"`
}

func (AllSign) OssSignNew(folderName, fileName, suffix string, args ...any) (OssSignNewResponse, error) {
	uri := IfString(runtime.GOOS == "linux", "http://oss:8080", "http://127.0.0.1:28080")
	uri += "/api/oss/oss-sign-new"

	isAccelerate := false
	if len(args) > 0 {
		if val, ok := args[0].(bool); ok && val {
			isAccelerate = true
		}
	}

	param := make(map[string]any)
	param["folder_name"] = folderName
	param["file_name"] = fileName
	param["suffix"] = suffix
	param["is_accelerate"] = isAccelerate
	_, response, err := DefaultClient().RequestGet(context.Background(), http.MethodGet, uri, param, nil)
	if err != nil {
		return OssSignNewResponse{}, err
	}

	var ossSignNewResponse OssSignNewResponse
	json.Unmarshal(response, &ossSignNewResponse)

	return ossSignNewResponse, err
}

type ossSignHeader struct {
	Key   string `json:"key"`   //header key
	Value string `json:"value"` //header value
}
type OssSignUrlResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		FileURL   string          `json:"file_url"`   //文件地址
		UploadURL string          `json:"upload_url"` //上传地址
		Header    []ossSignHeader `json:"header"`     //header
	} `json:"data"`
}

func (AllSign) OssSignUrl(folderName, fileName, suffix string, args ...any) (OssSignUrlResponse, error) {
	uri := IfString(runtime.GOOS == "linux", "http://oss:8080", "http://127.0.0.1:28080")
	uri += "/api/oss/oss-sign-url"

	isAccelerate := false
	if len(args) > 0 {
		if val, ok := args[0].(bool); ok && val {
			isAccelerate = true
		}
	}

	param := make(map[string]any)
	param["folder_name"] = folderName
	param["file_name"] = fileName
	param["suffix"] = suffix
	param["is_accelerate"] = isAccelerate
	_, response, err := DefaultClient().RequestGet(context.Background(), http.MethodGet, uri, param, nil)
	if err != nil {
		return OssSignUrlResponse{}, err
	}

	var ossSignUrlResponse OssSignUrlResponse
	json.Unmarshal(response, &ossSignUrlResponse)

	return ossSignUrlResponse, err
}

type signHeader struct {
	Key   string `json:"key"`   //header key
	Value string `json:"value"` //header value
}
type AwsSignResponse struct {
	FileUrl   string       `json:"file_url"`   //文件地址
	UploadUrl string       `json:"upload_url"` //上传地址
	Header    []signHeader `json:"header"`     //header
}

func (AllSign) AwsSign(folderName, fileName, suffix, fileMd5 string) (AwsSignResponse, error) {
	uri := IfString(runtime.GOOS == "linux", "http://oss:8080", "http://127.0.0.1:28080")
	uri += "/api/oss/aws-sign"

	param := make(map[string]any)
	param["folder_name"] = folderName
	param["file_name"] = fileName
	param["suffix"] = suffix
	param["file_md5"] = fileMd5
	_, response, err := DefaultClient().RequestGet(context.Background(), http.MethodGet, uri, param, nil)
	if err != nil {
		return AwsSignResponse{}, err
	}

	var awsSignResponse AwsSignResponse
	json.Unmarshal(response, &awsSignResponse)

	return awsSignResponse, err
}
