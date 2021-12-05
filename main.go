package main

import (
	"net/http"

	"github.com/chunlinwang/golang-gin-docker-boilerplate/config"
	"github.com/chunlinwang/golang-gin-docker-boilerplate/routes"
)

func main() {

	port := config.GetConfig("APP_PORT")

	router := routes.Init()

	http.ListenAndServe(port, router)
}
