package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	database "urlshortener/Database"
	resource "urlshortener/Resource"
	common "urlshortener/common"

	"github.com/mitchellh/mapstructure"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	configBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(configBytes, &common.Config)
	if err != nil {
		fmt.Println(err)
	}

	mapstructure.Decode(common.Config["db"].(map[string]interface{}), &common.Db)

	err = database.Connect()
	if err != nil {
		fmt.Println("Error occured while connecting to database")
		return
	}
	resource.InitHandlers()
}
