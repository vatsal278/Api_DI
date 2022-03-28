package main

import (
	"Api_DI/router"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(":9090", r))

}
