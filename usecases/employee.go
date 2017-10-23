package usecases

import (
	"github.com/ntrianta/cerndemo/entities"
)

type EmployeeInteractor struct {
	EmployeeRepository entities.EmployeeRepository
}

func (interactor *EmployeeInteractor) Store(employee *entities.Employee) error {
	err := employee.Validate()
	if err != nil {
		return err
	}
	interactor.EmployeeRepository.Store(employee)
	return err
}

func (interactor *EmployeeInteractor) List() []*entities.Employee {
	e := interactor.EmployeeRepository.List()
	return e
}

func (interactor *EmployeeInteractor) Read(id int) *entities.Employee {
	e := interactor.EmployeeRepository.Read(id)
	return e
}

func (interactor *EmployeeInteractor) Delete(id int) error {
	interactor.EmployeeRepository.Delete(id)
	return nil
}
