package main

import (
	"github.com/ntrianta/cerndemo/infrastructure"
	"github.com/ntrianta/cerndemo/interfaces"
	"github.com/ntrianta/cerndemo/usecases"
)

func main() {

	employeeFileHandler := infrastructure.FileHandler{
		File: "./employees",
	}

	leaveFileHandler := infrastructure.FileHandler{
		File: "./leave",
	}

	//Employees

	fileEmployeeRepo := interfaces.NewFileEmployeeRepo(employeeFileHandler)
	employeeInteractor := new(usecases.EmployeeInteractor)
	employeeInteractor.EmployeeRepository = fileEmployeeRepo
	employeeWebHandler := new(interfaces.EmployeeHandler)
	employeeWebHandler.EmployeeInteractor = employeeInteractor

	//Leave

	fileLeaveRepo := interfaces.NewFileLeaveRepo(leaveFileHandler)
	leaveInteractor := new(usecases.LeaveInteractor)
	leaveInteractor.LeaveRepository = fileLeaveRepo
	leaveWebHandler := new(interfaces.LeaveHandler)
	leaveWebHandler.LeaveInteractor = leaveInteractor

	web := &infrastructure.Web{
		EmployeeHandler: employeeWebHandler,
		LeaveHandler:    leaveWebHandler,
	}

	web.Serve()

}
