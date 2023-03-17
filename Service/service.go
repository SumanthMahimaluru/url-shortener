package Service

import (
	"net/http"
	repository "urlshortener/Repository"
	common "urlshortener/common"
)

func AddLink(data common.Data) common.HTTPResponse {
	link, err := repository.Find(data.URL)
	if err != nil {
		return common.HTTPResponse{Msg: "error occured while fetching data", Status: 500}
	}
	if link.URL != "" {
		return common.HTTPResponse{Msg: "Url already exists", Status: http.StatusConflict}
	}

	genString := common.RandStringBytes(10)

	domain := common.Config["host"].(string)
	port := common.Config["port"].(string)

	newurl := "http://" + domain + ":" + port + "/short/" + genString
	data.NewURL = genString
	err = repository.Insert(data)
	if err != nil {
		return common.HTTPResponse{Msg: "error occured while inserting data", Status: http.StatusInternalServerError}
	}
	return common.HTTPResponse{Msg: "Added link ", Data: newurl, Status: http.StatusAccepted}
}

func GetLink(newurl string) common.HTTPResponse {
	data, err := repository.GetNewUrl(newurl)
	if err != nil {
		return common.HTTPResponse{Msg: "error occured while fetching data", Status: 500}
	}
	if data.NewURL == "" {
		return common.HTTPResponse{Msg: "Url is expired or not found", Status: 400}
	}
	return common.HTTPResponse{Msg: "Fetched linked successfully ", Data: *data, Status: 200}
}
