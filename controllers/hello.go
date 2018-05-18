package controllers

import (
	"io"
	"net/http"
)

func Hello(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello World")
}

func Test(resp http.ResponseWriter, req *http.Request) {
	http.ServeFile(resp, req, "./test/test.html")
}
