package entities

import (
	"errors"
	"fmt"
)

type EmployeeRepository interface {
	Store(employee *Employee)
	List() []*Employee
	Read(id int) *Employee
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

func (e *Employee) Validate() error {

	fmt.Println(e.Age)

	if e.Age < 18 {
		fmt.Println("ti fasi")
		return errors.New("underage employee")
	} else if e.ID < 1 {
		return errors.New("invalid id")
	} else {
		return nil
	}
}
