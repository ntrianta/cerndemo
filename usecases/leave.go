package usecases

import (
	"fmt"

	"github.com/ntrianta/cerndemo/entities"
)

type LeaveInteractor struct {
	LeaveRepository entities.LeaveRepository
}

func (interactor *LeaveInteractor) Store(leave *entities.Leave) error {

	fmt.Println("Usecases layer. Leave Interactor, store.")

	err := leave.Validate()

	if err != nil {
		return err
	}

	fmt.Println("Usecases layer. Leave Interactor, store.")

	interactor.LeaveRepository.Store(leave)
	return nil

}

func (interactor *LeaveInteractor) List() []*entities.Leave {
	l := interactor.LeaveRepository.List()
	return l
}

func (interactor *LeaveInteractor) Read(id int) *entities.Leave {

	fmt.Println("Usecases layer. Leave Interactor, read.")
	fmt.Println("Usecases layer. Leave Interactor, application rules applied.")

	l := interactor.LeaveRepository.Read(id)

	return l
}

func (interactor *LeaveInteractor) Delete(id int) error {
	interactor.LeaveRepository.Delete(id)
	return nil
}
