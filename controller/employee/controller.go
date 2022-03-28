package employee

import (
	"Api_DI/controller"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Api_DI/database"
	"Api_DI/model"
	"io/ioutil"
)

type employee struct {
	db database.IDB
}

func NewController(db database.IDB) controller.IController {
	return &employee{
		db: db,
	}
}

// AllEmployee = Select Employee API
func (e employee) AllEmployee() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var response model.Response

		rows, err := e.db.Get()
		if err != nil {
			log.Println(err.Error())
			return
		}

		response.Status = 200
		response.Message = "Success"
		response.Data = rows

		w.WriteHeader(202)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

// InsertEmployee = Insert Employee API
func (e employee) InsertEmployee() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.Response
		var employees model.Employee
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			return
		}
		err = json.Unmarshal(body, &employees)
		if err != nil {
			log.Println(err.Error())
			return
		}
		id, err := e.db.Insert(employees)
		if err != nil {
			log.Println(err.Error())
			return
		}

		employees.Id = id
		response.Status = 200
		response.Message = "Insert data successfully"
		response.Data = []model.Employee{
			employees,
		}
		fmt.Print("Insert data to database")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(response)
	}
}

func (e employee) UpdateEmployee() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.Response
		var employees model.Employee
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			return
		}
		err = json.Unmarshal(body, &employees)
		if err != nil {
			log.Println(err.Error())
			return
		}

		id, err := e.db.Update(employees)
		if err != nil {
			log.Println(err.Error())
			return
		}
		employees.Id = id
		response.Status = 200
		response.Message = "Update data successfully"
		response.Data = []model.Employee{
			employees,
		}
		fmt.Print("Update data successfully")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// DeleteEmployee = Delete employee API
func (e employee) DeleteEmployee() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var response model.Response
		var id int64
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			return
		}
		err = json.Unmarshal(body, &id)
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = e.db.Delete(id)
		if err != nil {
			log.Println(err.Error())
			return
		}

		response.Status = 200
		response.Message = "Delete data successfully"
		fmt.Print("Delete data successfully")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
