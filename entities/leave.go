package entities

import (
	"errors"
	"fmt"
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

	fmt.Println("Entities layer. Leave entity. Business rules validation")

	timeFormat := "02-Jan-2006"

	start, _ := time.Parse(timeFormat, l.Start)
	end, _ := time.Parse(timeFormat, l.End)
	duration := end.Sub(start).Hours()

	if duration < 0 {
		return errors.New("start time after end time")
	} else if duration > 168 {
		return errors.New("no leave for more than a week")
	} else {
		return nil
	}

}
