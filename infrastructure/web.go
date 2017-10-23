package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ntrianta/cerndemo/interfaces"
)

type Web struct {
	EmployeeHandler *interfaces.EmployeeHandler
	LeaveHandler    *interfaces.LeaveHandler
}

func (web Web) Serve() {

	mainRouter := mux.NewRouter()
	getSubrouter := mainRouter.Methods("GET").Subrouter()
	postSubrouter := mainRouter.Methods("POST").Subrouter()
	deleteSubrouter := mainRouter.Methods("DELETE").Subrouter()
	putSubrouter := mainRouter.Methods("PUT").Subrouter()

	//Employee
	postSubrouter.HandleFunc("/employee", respond(web.EmployeeHandler.Add))
	getSubrouter.HandleFunc("/employee", respond(web.EmployeeHandler.List))
	getSubrouter.HandleFunc("/employee/{id:[0-9]+}", respond(web.EmployeeHandler.Read))
	deleteSubrouter.HandleFunc("/employee/{id:[0-9]+}", respond(web.EmployeeHandler.Delete))
	putSubrouter.HandleFunc("/employee", notImplemented)

	//LeaveGroup
	postSubrouter.HandleFunc("/leave", respond(web.LeaveHandler.Add))
	getSubrouter.HandleFunc("/leave", respond(web.LeaveHandler.List))
	getSubrouter.HandleFunc("/leave/{id:[0-9]+}", respond(web.LeaveHandler.Read))
	deleteSubrouter.HandleFunc("/leave/{id:[0-9]+}", respond(web.LeaveHandler.Delete))
	putSubrouter.HandleFunc("/leave", notImplemented)

	http.Handle("/", mainRouter)
	_ = http.ListenAndServe(":8080", nil)
}

func respond(fn func(r *http.Request, vars map[string]string) []byte) http.HandlerFunc {

	fmt.Println("Infrastructure layer. Web listener.")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		out := fn(r, mux.Vars(r))
		w.Write(out)
	})

}

func notImplemented(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("501 - Not Implemented"))
}
