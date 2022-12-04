package main

import (
	"go-rest/helper"
	"go-rest/middleware"

	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func NewServer(middleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    ":6969",
		Handler: middleware,
	}
}

func main() {

	errEnv := godotenv.Load("local.env")
	helper.CheckError(errEnv)

	server := InitServer()

	errServer := server.ListenAndServe()
	helper.CheckError(errServer)
}
