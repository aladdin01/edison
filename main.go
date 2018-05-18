package main

import (
	"net/http"
	"edison/controllers"
	"edison/app"
	"fmt"
	"log"
)

func main() {
	//db := app.GetConnection()

	http.HandleFunc("/api/hello", controllers.Hello)
	http.HandleFunc("/test", controllers.Test)
	http.HandleFunc("/api/put", controllers.Put)
	http.HandleFunc("/api/get", controllers.Get)
	http.HandleFunc("/api/delete", controllers.Delete)

	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		app.Log(err.Error())
		fmt.Println("______________________________")
		fmt.Println("Port is busy")
		log.Fatal()
	}
}
