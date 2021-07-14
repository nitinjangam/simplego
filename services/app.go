package services

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	userFilePath = "./data_mock/user_data.json"
	todoListPath = "./data_mock/to_do.json"
	serveIP      = "127.0.0.1"
	servePort    = ":8081"
)

//App structure will consist of all properties of application
type App struct {
	Router *mux.Router
	//in future we can include loggers, databases here
	//mocking database connection as files
	DbConn files
}

type files struct {
	UserFile *os.File
	TodoFile *os.File
}

//Initialize function to initializing connections with database or any network
func (a *App) Initialize() {
	//reading mock data from json files it will be replaced by connection to db in future
	var err error
	a.DbConn.UserFile, err = os.Open(userFilePath)
	if err != nil {
		log.Fatalf("Unable to open user data file: %v", err.Error())
	}
	a.DbConn.TodoFile, err = os.Open(todoListPath)
	if err != nil {
		log.Fatalf("Unable to open to-do list file: %v", err.Error())
	}
	a.Router = mux.NewRouter()
	a.setRoutes()
}

//Run function to start server
func (a *App) Run() {
	log.Printf("Server started on %v", serveIP+servePort)
	log.Fatal(http.ListenAndServe(servePort, a.Router))
}

func (a *App) setRoutes() {
	a.Router.HandleFunc("/getlist/{name}", a.GetTodoList).Methods("GET")
	a.Router.HandleFunc("/updatelist", a.UpdateTodoList).Methods("PUT")
	a.Router.HandleFunc("/addtask", a.AddTaskToList).Methods("POST")
	a.Router.HandleFunc("/removetask", a.RemoveTaskFromList).Methods("DELETE")
}
