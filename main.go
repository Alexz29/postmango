package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type PostmanDocumentInfo struct {
	Name   string
	Schema string
}

type PostmanItemRequest struct {
	Url string
}

type PostmanItemResponse struct {
	Code int
	Body string
}

type PostmanItem struct {
	Name     string
	Request  PostmanItemRequest
	Response []PostmanItemResponse
}

type PostmanDocument struct {
	Info PostmanDocumentInfo
	Item []PostmanItem
}

type Param struct {
	name  string
	value string
}

type EmptyQuery struct {
	Item []Param
}

var document PostmanDocument

func handleRequest(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	if r.URL.RawQuery != "" {
		uri = r.URL.Path + "?" + cleanParam(r.URL.RawQuery)
	}
	fmt.Println("request:", uri)
	for _, value := range document.Item {
		urlClear := strings.Replace(value.Request.Url, "{{url}}", "", 1)
		rUrl := cleanParam(url.PathEscape(urlClear))
		rUrl, err := url.QueryUnescape(rUrl)

		if err != nil {
			fmt.Println("error:", err)
		}

		if rUrl == uri {
			w.WriteHeader(value.Response[0].Code)
			w.Write([]byte(value.Response[0].Body))

			return
		}
	}

	if "/" == uri {
		w.WriteHeader(200)
		w.Write([]byte(""))

		return
	}

	w.WriteHeader(404)
	w.Write([]byte("Not found"))
}

func cleanParam(rawQuery string) string {

	array := strings.Split(rawQuery, "&")
	var s []string

	for _, value := range array {
		tmpVar := strings.Split(value, "=")
		if tmpVar[0] != "" {
			var str = tmpVar[0] + "=null"
			s = append(s, str)
		}

	}

	return strings.Join(s, "&")
}

func main() {

	file := flag.String("f", "./fixture.json", "Mock server json file")
	port := flag.String("p", "8080", "Server port")
	host := flag.String("h", "localhost", "Server host address")

	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
	fmt.Println("file:", *file)

	jsonData, err := ioutil.ReadFile(*file)

	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(jsonData), &document)

	http.HandleFunc("/", handleRequest)

	if err := http.ListenAndServe(*host+":"+*port, nil); err != nil {
		panic(err)
	}
}
