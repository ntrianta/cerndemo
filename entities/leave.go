package entities

import (
	"errors"
	"time"
)

type LeaveRepository interface {
	Store(leave Leave)
	List() []*Leave
	Read(id int) Leave
	Delete(id int)
	//Update()
}

type Leave struct {
	ID       int
	Employee int
	Start    string
	End      string
}

func (ag *AcctGroup) Validate(l Leave) error {

	start, _ := time.Parse(time.RFC822, l.Start)
	end, _ := time.Parse(time.RFC822, l.End)

	if start.After(end) {
		return errors.New("start time after end time")
	} else {
		return nil
	}
}
