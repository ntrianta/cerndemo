package main

import (
	"database/sql"
	"fmt"

	"github.com/ntrianta/cerndemo/infrastructure"
	"github.com/ntrianta/cerndemo/interfaces"
	"github.com/ntrianta/cerndemo/usecases"

	_ "gopkg.in/go-sql-driver/mysql.v1"
)

func initDB() *sql.DB {
	user := "root"
	password := "root"
	host := "localhost:3306"
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/cerndemo?tls=false", user, password, host)

	db, err := sql.Open("mysql", connectString)
	fmt.Println(err)
	return db
}

func main() {

	// employeeFileHandler := infrastructure.FileHandler{
	// 	File: "./employees",
	// }

	// leaveFileHandler := infrastructure.FileHandler{
	// 	File: "./leave",
	// }

	db := initDB()

	databaseHandler := infrastructure.MySQL{
		Db: db,
	}

	//Employees

	//fileEmployeeRepo := interfaces.NewFileEmployeeRepo(employeeFileHandler)
	relationalEmployeeRepo := interfaces.NewRelationalEmployeeRepo(databaseHandler)
	employeeInteractor := new(usecases.EmployeeInteractor)
	employeeInteractor.EmployeeRepository = relationalEmployeeRepo
	employeeWebHandler := new(interfaces.EmployeeHandler)
	employeeWebHandler.EmployeeInteractor = employeeInteractor

	//Leave

	//fileLeaveRepo := interfaces.NewFileLeaveRepo(leaveFileHandler)
	relationalLeaveRepo := interfaces.NewRelationalLeaveRepo(databaseHandler)
	leaveInteractor := new(usecases.LeaveInteractor)
	leaveInteractor.LeaveRepository = relationalLeaveRepo
	leaveWebHandler := new(interfaces.LeaveHandler)
	leaveWebHandler.LeaveInteractor = leaveInteractor

	web := &infrastructure.Web{
		EmployeeHandler: employeeWebHandler,
		LeaveHandler:    leaveWebHandler,
	}

	web.Serve()

}
