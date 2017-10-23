package interfaces

import (
	"fmt"

	"github.com/ntrianta/cerndemo/entities"
)

type FileHandler interface {
	Store(writable string) error
	// 	List(query string) Rows
	// 	ReadByID(query string, id int, readable *string)
	// 	ReadByName(query string, name string, readable *int)
	// 	Delete(query string, id int) error
}

type FileRepo struct {
	fileHandler FileHandler
}

type FileEmployeeRepo FileRepo
type FileLeaveRepo FileRepo

func NewFileEmployeeRepo(handler FileHandler) *FileEmployeeRepo {
	fileEmployeeRepo := new(FileEmployeeRepo)
	fileEmployeeRepo.fileHandler = handler
	return fileEmployeeRepo
}

func NewFileLeaveRepo(handler FileHandler) *FileLeaveRepo {
	fileLeaveRepo := new(FileLeaveRepo)
	fileLeaveRepo.fileHandler = handler
	return fileLeaveRepo
}

func (repo *FileEmployeeRepo) Store(employee *entities.Employee) {
	repo.fileHandler.Store(fmt.Sprintf("=======================\n"))
	repo.fileHandler.Store(fmt.Sprintf("Entry: %d\n", employee.ID))
	repo.fileHandler.Store(fmt.Sprintf("First Name: %s\n", employee.Name))
	repo.fileHandler.Store(fmt.Sprintf("Last Name: %s\n", employee.Surname))
	repo.fileHandler.Store(fmt.Sprintf("Age: %d\n", employee.Age))
	repo.fileHandler.Store(fmt.Sprintf("Organic Unit: %s\n", employee.Unit))
	repo.fileHandler.Store(fmt.Sprintf("Supervisor: %s\n", employee.Supervisor))
}

func (repo *FileEmployeeRepo) List() []*entities.Employee {
	employees := make([]*entities.Employee, 0)
	return employees
}

func (repo *FileEmployeeRepo) Read(id int) *entities.Employee {
	employee := new(entities.Employee)
	return employee
}

func (repo *FileEmployeeRepo) Delete(id int) {
	fmt.Println("Deleted:", id)
}

func (repo *FileLeaveRepo) Store(leave *entities.Leave) {
	repo.fileHandler.Store(fmt.Sprintf("=======================\n"))
	repo.fileHandler.Store(fmt.Sprintf("Entry: %d\n", leave.ID))
	repo.fileHandler.Store(fmt.Sprintf("Employee: %d\n", leave.Employee))
	repo.fileHandler.Store(fmt.Sprintf("Start Date: %s\n", leave.Start))
	repo.fileHandler.Store(fmt.Sprintf("End Date: %s\n", leave.End))
}

func (repo *FileLeaveRepo) List() []*entities.Leave {
	leaves := make([]*entities.Leave, 0)
	return leaves
}

func (repo *FileLeaveRepo) Read(id int) *entities.Leave {
	leave := new(entities.Leave)
	return leave
}

func (repo *FileLeaveRepo) Delete(id int) {
	fmt.Println("Deleted:", id)
}
