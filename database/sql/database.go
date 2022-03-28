package sql

import (
	"Api_DI/config"
	"Api_DI/database"
	"Api_DI/model"
	"database/sql"
	"fmt"
)

type db struct {
	db *sql.DB
}

func NewDB() database.IDB {
	return &db{
		db: config.Connect(),
	}
}

func (d *db) Get() ([]model.Employee, error) {
	var employee model.Employee
	var arrEmployee []model.Employee
	rows, err := d.db.Query("SELECT id, name, city FROM employee")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&employee.Id, &employee.Name, &employee.City)
		if err != nil {
			return nil, err
		}
		arrEmployee = append(arrEmployee, employee)
	}
	return arrEmployee, nil
}

func (d *db) Insert(employee model.Employee) (string, error) {

	res, err := d.db.Exec("INSERT INTO employee(name, city) VALUES(?, ?)", employee.Name, employee.City)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()

	return fmt.Sprint(id), err
}

func (d *db) Update(employee model.Employee) (string, error) {

	res, err := d.db.Exec("UPDATE employee SET name=?, city=? WHERE id=?", employee.Id, employee.Name, employee.City)
	if err != nil {
		return "", err
	}
	id, err := res.LastInsertId()

	return fmt.Sprint(id), nil
}

func (d *db) Delete(id int64) error {
	_, err := d.db.Exec("DELETE FROM employee WHERE id=?", id)
	if err != nil {
		return err
	}

	return err
}
