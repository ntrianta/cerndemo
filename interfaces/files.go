package interfaces

import (
	"fmt"

	"github.com/ntrianta/cerndemo/entities"
)

type FileHandler interface {
	Store(writable []byte) error
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
	repo.fileHandler.Store([]byte(employee.Name))
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
	fmt.Println(leave)
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
