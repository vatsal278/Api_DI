package database

import "Api_DI/model"

type IDB interface {
	Get() ([]model.Employee, error)
	Insert(model.Employee) (string, error)
	Update(model.Employee) (string, error)
	Delete(id int64) error
}
