package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/farmerx/gorsa"
)

//md5

func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

//sha1

func SHA1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//sha256

func SHA256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//sha512

func SHA512(origin string) string {
	h := sha512.New()
	h.Write([]byte(origin))
	return hex.EncodeToString(h.Sum(nil))
}

//hmac sha1

func HmacSHA1(key, data string) []byte {
	keys := []byte(key)
	h := hmac.New(sha1.New, keys)
	h.Write([]byte(data))
	return h.Sum(nil)
}

//hmac sha256

func HmacSHA256(key, data string) []byte {
	keys := []byte(key)
	h := hmac.New(sha256.New, keys)
	h.Write([]byte(data))
	return h.Sum(nil)
}

//hmac sha512

func HmacSHA512(key, data string) []byte {
	keys := []byte(key)
	h := hmac.New(sha512.New, keys)
	h.Write([]byte(data))
	return h.Sum(nil)
}

//Rsa私钥加密

func RsaPriKeyEncrypt(plainText, key []byte) ([]byte, error) {
	gorsa.RSA.SetPrivateKey(string(key))
	var buffer bytes.Buffer
	cryptTextTemp, err := gorsa.RSA.PriKeyENCTYPT(plainText)
	if err != nil {
		return nil, err
	}
	buffer.Write(cryptTextTemp)

	return buffer.Bytes(), nil
}

//Rsa公钥解密

func RsaPubKeyDecrypt(cryptText, key []byte) ([]byte, error) {
	gorsa.RSA.SetPublicKey(string(key))
	var buffer bytes.Buffer
	plainTextTemp, err := gorsa.RSA.PubKeyDECRYPT(cryptText)
	if err != nil {
		return nil, err
	}
	buffer.Write(plainTextTemp)

	return buffer.Bytes(), nil
}

//Rsa公钥加密

func RsaPubKeyEncrypt(plainText, key []byte) ([]byte, error) {
	gorsa.RSA.SetPublicKey(string(key))

	var buffer bytes.Buffer
	pubenctypt, err := gorsa.RSA.PubKeyENCTYPT(plainText)
	if err != nil {
		return nil, err
	}

	buffer.Write(pubenctypt)

	return buffer.Bytes(), nil
}

//Rsa私钥解密

func RsaPriKeyDecrypt(cryptText, key []byte) ([]byte, error) {
	gorsa.RSA.SetPrivateKey(string(key))

	var buffer bytes.Buffer
	plainTextTemp, err := gorsa.RSA.PriKeyDECRYPT(cryptText)
	if err != nil {
		return nil, err
	}
	buffer.Write(plainTextTemp)

	return buffer.Bytes(), nil
}
