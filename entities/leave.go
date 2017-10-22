package entities

import (
	"errors"
	"time"
)

type LeaveRepository interface {
	Store(leave *Leave)
	List() []*Leave
	Read(id int) *Leave
	Delete(id int)
	//Update()
}

type Leave struct {
	ID       int
	Employee int
	Start    string
	End      string
}

func (l *Leave) Validate() error {

	start, _ := time.Parse(time.RFC822, l.Start)
	end, _ := time.Parse(time.RFC822, l.End)
	duration := end.Sub(start).Hours()

	if duration < 0 {
		return errors.New("start time after end time")
	} else if duration > 168 {
		return errors.New("no leave for more than a week")
	} else {
		return nil
	}

}
