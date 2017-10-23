package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ntrianta/cerndemo/entities"
)

type EmployeeInteractor interface {
	Store(Employee *entities.Employee) error
	List() []*entities.Employee
	Read(id int) *entities.Employee
	Delete(id int) error
}

type LeaveInteractor interface {
	Store(Leave *entities.Leave) error
	List() []*entities.Leave
	Read(id int) *entities.Leave
	Delete(id int) error
}

type EmployeeHandler struct {
	EmployeeInteractor EmployeeInteractor
}

type LeaveHandler struct {
	LeaveInteractor LeaveInteractor
}

//Employees

func (handler EmployeeHandler) Add(r *http.Request, vars map[string]string) []byte {

	fmt.Println("Interfaces layer. Employee web handler, store.")

	decoder := json.NewDecoder(r.Body)
	employee := new(entities.Employee)
	_ = decoder.Decode(employee)

	err := handler.EmployeeInteractor.Store(employee)
	if err != nil {
		out := err.Error()
		return []byte(out)
	}
	out := fmt.Sprintf("Added\n")
	return []byte(out)

}

func (handler EmployeeHandler) List(r *http.Request, vars map[string]string) []byte {

	res := handler.EmployeeInteractor.List()
	out := fmt.Sprintf("The result is: %s", res)
	return []byte(out)
}

func (handler EmployeeHandler) Read(r *http.Request, vars map[string]string) []byte {

	fmt.Println("Interfaces layer. Employee web handler, read.")

	id, _ := strconv.Atoi(vars["id"])

	res := handler.EmployeeInteractor.Read(id)
	out := fmt.Sprintf("The result is: %v \n", res.Name)
	return []byte(out)
}

func (handler EmployeeHandler) Delete(r *http.Request, vars map[string]string) []byte {

	id, _ := strconv.Atoi(vars["id"])

	handler.EmployeeInteractor.Delete(id)
	out := fmt.Sprintf("Deleted\n")
	return []byte(out)
}

//Leave
func (handler LeaveHandler) Add(r *http.Request, vars map[string]string) []byte {

	fmt.Println("Interfaces layer. Leave web handler, store.")

	decoder := json.NewDecoder(r.Body)
	leave := new(entities.Leave)
	_ = decoder.Decode(leave)

	err := handler.LeaveInteractor.Store(leave)
	if err != nil {
		out := err.Error()
		return []byte(out)
	}
	out := fmt.Sprintf("Added\n")
	return []byte(out)
}

func (handler LeaveHandler) List(r *http.Request, vars map[string]string) []byte {

	res := handler.LeaveInteractor.List()
	out := fmt.Sprintf("The result is: %s", res)
	return []byte(out)
}

func (handler LeaveHandler) Read(r *http.Request, vars map[string]string) []byte {

	id, _ := strconv.Atoi(vars["id"])

	res := handler.LeaveInteractor.Read(id)
	out := fmt.Sprintf("The result is: %v \n", res.ID)
	return []byte(out)
}

func (handler LeaveHandler) Delete(r *http.Request, vars map[string]string) []byte {

	id, _ := strconv.Atoi(vars["id"])

	handler.LeaveInteractor.Delete(id)
	out := fmt.Sprintf("Deleted\n")
	return []byte(out)
}
