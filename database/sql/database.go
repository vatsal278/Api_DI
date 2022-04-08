package sql

import (
	"Api_DI/config"
	"Api_DI/database"
	"Api_DI/model"
	"database/sql"
	"fmt"
	"strconv"
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

	id, err := strconv.Atoi(employee.Id)
	if err != nil {
		return "", err
	}
	res, err := d.db.Exec("UPDATE employee SET name=?, city=? WHERE id=?", employee.Name, employee.City, id)
	if err != nil {
		return "", err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return "", err
	}
	if rows == 0 {
		return "", fmt.Errorf("%v rows impacted", rows)
	}

	return fmt.Sprint(id), nil
}

func (d *db) Delete(id int64) error {
	result, err := d.db.Exec("DELETE FROM employee WHERE id=?", id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 1 {
		return nil
	}

	return fmt.Errorf("%v rows impacted", rows)
}
