package common

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type DBConfig struct {
	User     string `json:"user"`
	Host     string `json:"host"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Port     string `json:"port"`
}

type HTTPResponse struct {
	Msg    string      `json:"_msg"`
	Status int         `json:"_status"`
	Data   interface{} `json:"data"`
}

type Data struct {
	URL    string        `json:"url"`
	NewURL string        `json:"newurl"`
	Expiry time.Duration `json:"expiry"`
}

var Db = DBConfig{}
var Config map[string]interface{}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
