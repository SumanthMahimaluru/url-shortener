package resource

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	urlService "urlshortener/Service"
	"urlshortener/common"
)

func AddLink(w http.ResponseWriter, r *http.Request) {
	log.Println("Add Link")
	var data common.Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := urlService.AddLink(data)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(res.Status)

	fmt.Fprintf(w, res.Msg+"\n")
	if res.Status == http.StatusAccepted {
		fmt.Fprintf(w, res.Data.(string))
	}
	return
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	res := urlService.GetLink(pathArgs[2])
	if res.Status != 200 {
		return
	}
	var data map[string]interface{}
	bytes, _ := json.Marshal(res.Data)
	_ = json.Unmarshal(bytes, &data)

	http.Redirect(w, r, data["url"].(string), http.StatusPermanentRedirect)
	return
}

func InitHandlers() {
	http.HandleFunc("/addLink", AddLink)
	http.HandleFunc("/short/", GetLink)
	log.Fatal(http.ListenAndServe(":"+common.Config["port"].(string), nil))
}
