package interfaces

import (
	"fmt"

	"github.com/ntrianta/cerndemo/entities"
)

type RelationalHandler interface {
	Store(query string, writable ...interface{}) error
	List(query string) Rows
	Read(query string, id int, readable *string)
	Delete(query string, id int) error
}

type Rows interface {
	Scan(readable ...interface{}) error
	Next() bool
	Close() error
}

type RelationalRepo struct {
	relationalHandler RelationalHandler
}

type RelationalEmployeeRepo RelationalRepo
type RelationalLeaveRepo RelationalRepo

func NewRelationalEmployeeRepo(relationalHandler RelationalHandler) *RelationalEmployeeRepo {
	relationalEmployeeRepo := new(RelationalEmployeeRepo)
	relationalEmployeeRepo.relationalHandler = relationalHandler
	return relationalEmployeeRepo
}

func NewRelationalLeaveRepo(relationalHandler RelationalHandler) *RelationalLeaveRepo {
	relationalLeaveRepo := new(RelationalLeaveRepo)
	relationalLeaveRepo.relationalHandler = relationalHandler
	return relationalLeaveRepo
}

//Employee

func (repo *RelationalEmployeeRepo) Store(employee *entities.Employee) {

	fmt.Println("Interfaces layer. Employee relational repository, store.")

	query := "INSERT INTO employees (ID, Name, Surname, Unit, Age, Supervisor) VALUES(?, ?, ?, ?, ?, ?)"
	repo.relationalHandler.Store(query, employee.ID, employee.Name, employee.Surname, employee.Unit, employee.Age, employee.Supervisor)
}

func (repo *RelationalEmployeeRepo) List() []*entities.Employee {
	query := "SELECT id, name, surname FROM employees"
	employees := make([]*entities.Employee, 0)
	rows := repo.relationalHandler.List(query)
	defer rows.Close()

	for rows.Next() {
		e := new(entities.Employee)
		rows.Scan(&e.ID, &e.Name, &e.Surname)
		employees = append(employees, e)
	}

	return employees
}

func (repo *RelationalEmployeeRepo) Read(id int) *entities.Employee {

	fmt.Println("Interfaces layer. Employee relational repository, read.")

	query := "SELECT name FROM employees where ID = ?"
	employee := new(entities.Employee)
	repo.relationalHandler.Read(query, id, &employee.Name)
	return employee
}

func (repo *RelationalEmployeeRepo) Delete(id int) {
	query := "DELETE FROM employees WHERE ID = ?"
	repo.relationalHandler.Delete(query, id)
}

//Leave

func (repo *RelationalLeaveRepo) Store(leave *entities.Leave) {

	fmt.Println("Interfaces layer. Leave relational repository, store.")

	query := "INSERT INTO leaves (ID, Employee, Start, End) VALUES(?, ?, ?, ?)"
	repo.relationalHandler.Store(query, leave.ID, leave.Employee, leave.Start, leave.End)
}

func (repo *RelationalLeaveRepo) List() []*entities.Leave {
	query := "SELECT id, employee, start, end FROM leaves"
	leaves := make([]*entities.Leave, 0)
	rows := repo.relationalHandler.List(query)
	defer rows.Close()

	for rows.Next() {
		l := new(entities.Leave)
		rows.Scan(&l.ID, &l.Employee, &l.Start, &l.End)
		leaves = append(leaves, l)
	}

	return leaves
}

func (repo *RelationalLeaveRepo) Read(id int) *entities.Leave {

	fmt.Println("Interfaces layer. Leave relational repository, read.")

	query := "SELECT start FROM leaves where id = ?"
	leave := new(entities.Leave)
	repo.relationalHandler.Read(query, id, &leave.Start)
	return leave
}

func (repo *RelationalLeaveRepo) Delete(id int) {
	query := "DELETE FROM leaves WHERE id = ?"
	repo.relationalHandler.Delete(query, id)
}
