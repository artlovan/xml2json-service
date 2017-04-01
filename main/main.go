package main

import (
	"net/http"
	"io/ioutil"
	"strings"
	xml2json "github.com/artlovan/goxml2json"
)

/**
  * Service will convert any xml structure
  * to json and return ot back.
*/

func handler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	jsonData, err := xml2json.Convert(strings.NewReader(string(body)))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData.Bytes())
}

func main() {
	service := http.NewServeMux()
	service.HandleFunc("/api/xml2json", handler)
	http.ListenAndServe(":6061", service)
}
