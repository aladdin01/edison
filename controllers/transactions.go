package controllers

import (
	"net/http"
	"encoding/json"
	"edison/kvstore"
	"log"
	"fmt"
)

type Error struct {
	Field   string
	Message string
	Code    int
}

func Put(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	// todo validation

	key := req.Form.Get("key") // file name
	value := req.Form.Get("value") // file name

	kvstore.Put(key, value)

	c := map[string]interface{}{"success" : true, "key" : key, "value": value, "msg": fmt.Sprintf("Key %s saved with value %s", key, value)}
	rr, _ := json.Marshal(c)
	resp.Write([]byte(rr))
}

func Get(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	key := req.Form.Get("key") // file name
	// todo validation
	value, err := kvstore.Get(key)

	if (err != nil) {
		responseError(400, err.Error(), resp)
		return
	}

	c := map[string]interface{}{"success" : true, "key" : key, "value": value}
	rr, _ := json.Marshal(c)
	resp.Write([]byte(rr))
}

func Delete(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	key := req.Form.Get("key") // file name
	// todo validation

	err := kvstore.Delete(key)

	if (err != nil) {
		responseError(400, err.Error(), resp)
		return
	}

	c := map[string]interface{}{"success" : true, "key" : key, "msg": fmt.Sprintf("Key %s deleted", key)}
	rr, _ := json.Marshal(c)
	resp.Write([]byte(rr))
}

func responseError(HttpError int, s string, resp http.ResponseWriter) {
	log.Println(s)
	resp.WriteHeader(HttpError)
	resp.Header().Set("Content-Type", "application/json")
	c := map[string]interface{}{"success" : false, "msg" : s}
	rr, _ := json.Marshal(c)
	resp.Write(rr)
}
