package entities

import "errors"

type EmployeeRepository interface {
	Store(employee Employee)
	List() []*Employee
	Read(id int) Employee
	Delete(id int)
	//Update()
}

type Employee struct {
	ID         int
	Name       string
	Surname    string
	Unit       string
	Age        int
	Supervisor string
}

func (ag *AcctGroup) Validate(e Employee) error {

	if e.Age < 18 {
		return errors.New("underage employee")
	} else if e.ID < 1 {
		return errors.New("invalid id")
	} else {
		return nil
	}
}
