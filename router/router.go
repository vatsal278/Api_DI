package router

import (
	"Api_DI/controller/employee"
	"Api_DI/database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	db := sql.NewDB()
	e := employee.NewController(db)
	router := mux.NewRouter()
	router.HandleFunc("/getEmployee", e.AllEmployee()).Methods("GET")
	router.HandleFunc("/insertEmployee", e.InsertEmployee()).Methods("POST")
	router.HandleFunc("/updateEmployee", e.UpdateEmployee()).Methods("PUT")
	router.HandleFunc("/deleteEmployee/{id}", e.DeleteEmployee()).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 9090")

	return router
}
