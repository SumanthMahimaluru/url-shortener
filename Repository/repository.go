package repository

import (
	"time"
	database "urlshortener/Database"
	common "urlshortener/common"
)

var query string = "INSERT INTO urlshortener (url, newurl, expirydate) VALUES (?, ?, ?)"
var findUrlQuery string = "SELECT * from urlshortener where url = ?"
var GetNewUrlQuery string = "SELECT * from urlshortener where newurl = ? and expirydate > now()"

type ResData struct {
	URL        string    `json:"url"`
	NewURL     string    `json:"newurl"`
	ExpiryDate time.Time `json:"expirydate"`
}

func Insert(data common.Data) error {
	result, err := database.GetClient().Query(query, data.URL, data.NewURL, time.Now().Add(data.Expiry*time.Hour))
	if err != nil {
		return err
	}
	defer result.Close()
	return nil
}

func Find(url string) (*common.Data, error) {
	result, err := database.GetClient().Query(findUrlQuery, url)
	if err != nil {
		return nil, err
	}
	var data ResData
	for result.Next() {
		err = result.Scan(&data.URL, &data.NewURL, &data.ExpiryDate)
		if err != nil {
			return nil, err
		}
	}

	return &common.Data{URL: data.URL, NewURL: data.NewURL}, nil
}

func GetNewUrl(newurl string) (*common.Data, error) {
	result, err := database.GetClient().Query(GetNewUrlQuery, newurl)
	if err != nil {
		return nil, err
	}
	var data ResData
	for result.Next() {
		err = result.Scan(&data.URL, &data.NewURL, &data.ExpiryDate)
		if err != nil {
			return nil, err
		}
	}
	return &common.Data{URL: data.URL, NewURL: data.NewURL}, nil
}
