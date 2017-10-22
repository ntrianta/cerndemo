package usecases

import "github.com/ntrianta/cerndemo/entities"

type LeaveInteractor struct {
	LeaveRepository entities.LeaveRepository
}

func (interactor *LeaveInteractor) Store(Leave *entities.Leave) error {
	interactor.LeaveRepository.Store(Leave)
	return nil
}

func (interactor *LeaveInteractor) List() []*entities.Leave {
	l := interactor.LeaveRepository.List()
	return l
}

func (interactor *LeaveInteractor) Read(id int) *entities.Leave {
	l := interactor.LeaveRepository.Read(id)
	return l
}

func (interactor *LeaveInteractor) Delete(id int) error {
	interactor.LeaveRepository.Delete(id)
	return nil
}
