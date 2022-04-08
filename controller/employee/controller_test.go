package employee_test

import (
	"Api_DI/controller/employee"
	"Api_DI/model"
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
)

type mockDB struct {
	f   func() (interface{}, error)
	age int64
}

func (m mockDB) Get() ([]model.Employee, error) {
	fmt.Println("Mock DB")
	resp, err := m.f()
	return resp.([]model.Employee), err
}

func (m mockDB) Insert(employee model.Employee) (string, error) {
	fmt.Println("Mock DB")
	resp, err := m.f()
	return resp.(string), err
}

func (m mockDB) Update(employee model.Employee) (string, error) {
	fmt.Println("Mock DB")
	resp, err := m.f()
	return resp.(string), err
}

func (m mockDB) Delete(id int64) error {
	fmt.Println("Mock DB")
	_, err := m.f()
	return err
}

func TestAllEmployee(t *testing.T) {
	errorCase := true

	db := &mockDB{
		f: func() (interface{}, error) {
			if errorCase {
				return []model.Employee{}, errors.New("error case")
			}
			return []model.Employee{
				{
					Id:   "100",
					Name: "Vatsal",
					City: "Jaipur",
				},
			}, nil
		},
	}

	eController := employee.NewController(db)

	allEmployeeController := eController.AllEmployee()

	// table driven tests
	tests := []struct {
		name       string
		isError    bool
		statusCode int
	}{
		{
			name:       "Failure Case",
			isError:    true,
			statusCode: 200,
		},
		{
			name:       "Success Case",
			isError:    false,
			statusCode: 202,
		},
	}

	// to execute the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			r := httptest.NewRequest("GET", "/getAllEmployees", nil)

			errorCase = tt.isError

			allEmployeeController(w, r)

			t.Log(w.Header())
			t.Log(w.Code)
			t.Log(w.Body.String())

			if w.Code != tt.statusCode {
				t.Errorf("Want: %v, Got: %v", tt.statusCode, w.Code)
			}
		})
	}
}

func TestInsertEmployee(t *testing.T) {
	errorCase := true

	db := &mockDB{
		f: func() (interface{}, error) {
			if errorCase {
				return "", errors.New("error case")
			}
			return "101", nil
		},
	}

	eController := employee.NewController(db)

	insertEmployeeController := eController.InsertEmployee()

	// table driven tests
	tests := []struct {
		name       string
		isError    bool
		statusCode int
	}{
		{
			name:       "Failure Case",
			isError:    true,
			statusCode: 200,
		},
		{
			name:       "Success Case",
			isError:    false,
			statusCode: 200,
		},
	}

	// to execute the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var newstr = []byte(`{"Id":101, "Name": yogesh, "City":bombay}`)
			r := httptest.NewRequest("POST", "/insertEmployees", bytes.NewBuffer(newstr))

			errorCase = tt.isError

			insertEmployeeController(w, r)

			t.Log(w.Header())
			t.Log(w.Code)
			t.Log(w.Body.String())

			if w.Code != tt.statusCode {
				t.Errorf("Want: %v, Got: %v", tt.statusCode, w.Code)
			}
		})
	}
}

func TestUpdateEmployee(t *testing.T) {
	errorCase := true

	db := &mockDB{
		f: func() (interface{}, error) {
			if errorCase {
				return "", errors.New("error case")
			}
			return "101", nil
		},
	}

	eController := employee.NewController(db)

	updateEmployeeController := eController.UpdateEmployee()

	// table driven tests
	tests := []struct {
		name       string
		isError    bool
		statusCode int
	}{
		{
			name:       "Failure Case",
			isError:    true,
			statusCode: 200,
		},
		{
			name:       "Success Case",
			isError:    false,
			statusCode: 200,
		},
	}

	// to execute the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var newstr = []byte(`{"Id":101, "Name":yogesh, "City":bombay}`)
			r := httptest.NewRequest("PUT", "/updateEmployees", bytes.NewBuffer(newstr))

			errorCase = tt.isError

			updateEmployeeController(w, r)

			t.Log(w.Header())
			t.Log(w.Code)
			t.Log(w.Body.String())

			if w.Code != tt.statusCode {
				t.Errorf("Want: %v, Got: %v", tt.statusCode, w.Code)
			}
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	errorCase := true

	db := &mockDB{
		f: func() (interface{}, error) {
			if errorCase {
				return 0, errors.New("error case")
			}
			return "", nil
		},
	}

	eController := employee.NewController(db)

	deleteEmployeeController := eController.DeleteEmployee()

	// table driven tests
	tests := []struct {
		name       string
		isError    bool
		statusCode int
	}{
		{
			name:       "Failure Case",
			isError:    true,
			statusCode: 200,
		},
		{
			name:       "Success Case",
			isError:    false,
			statusCode: 200,
		},
	}

	// to execute the tests in the table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			var newstr = []byte(`{"Id":100}`)
			r := httptest.NewRequest("DELETE", "/deleteEmployees", bytes.NewBuffer(newstr))

			errorCase = tt.isError

			deleteEmployeeController(w, r)

			t.Log(w.Header())
			t.Log(w.Code)
			t.Log(w.Body.String())

			if w.Code != tt.statusCode {
				t.Errorf("Want: %v, Got: %v", tt.statusCode, w.Code)
			}
		})
	}
}
