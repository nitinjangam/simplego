package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/nitinjangam/simplego/models"
)

type inbound struct {
	Name  string `json:"name"`
	Tasks struct {
		TaskName string `json:"taskName"`
		Status   string `json:"status"`
	} `json:"tasks"`
}

//GetTodoList function
func (a *App) GetTodoList(w http.ResponseWriter, r *http.Request) {
	l := &models.List{}
	vars := mux.Vars(r)
	l.Name = vars["name"]
	err := l.GetTodoList(a.DbConn.TodoFile)
	if err != nil {
		log.Fatalf("Error while GetTodoList: %v", err.Error())
	}
	bs, err := json.Marshal(l)
	if err != nil {
		log.Fatalf("Error while json marshal: %v", err.Error())
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(bs)
}

//UpdateTodoList function
func (a *App) UpdateTodoList(w http.ResponseWriter, r *http.Request) {
	l := &models.List{}
	inp := &inbound{}
	err := inp.getInputData(r)
	if err != nil {
		log.Fatalf("Error while reading input..")
	}
	l.Name = inp.Name
	l.Tasks = append(l.Tasks, inp.Tasks)
	err = l.UpdateTask(a.DbConn.TodoFile)
	if err != nil {
		log.Fatalf("Error while UpdateTask: %v", err.Error())
	}
	bs, err := json.Marshal(l)
	if err != nil {
		log.Fatalf("Error while json marshal: %v", err.Error())
	}
	w.Write(bs)
}

//AddTaskToList function
func (a *App) AddTaskToList(w http.ResponseWriter, r *http.Request) {
	l := &models.List{}
	inp := &inbound{}
	err := inp.getInputData(r)
	if err != nil {
		log.Fatalf("Error while reading input..")
	}
	l.Name = inp.Name
	l.Tasks = []models.Task{}
	l.Tasks = append(l.Tasks, inp.Tasks)
	err = l.AddTask(a.DbConn.TodoFile)
	if err != nil {
		log.Fatalf("Error while AddTask: %v", err.Error())
	}
	bs, err := json.Marshal(l)
	if err != nil {
		log.Fatalf("Error while json marshal: %v", err.Error())
	}
	w.Write(bs)
}

//RemoveTaskFromList function
func (a *App) RemoveTaskFromList(w http.ResponseWriter, r *http.Request) {
	l := &models.List{}
	inp := &inbound{}
	err := inp.getInputData(r)
	if err != nil {
		log.Fatalf("Error while reading input..")
	}
	l.Name = inp.Name
	l.Tasks = append(l.Tasks, inp.Tasks)
	err = l.RemoveTask(a.DbConn.TodoFile)
	if err != nil {
		log.Fatalf("Error while RemoveTask: %v", err.Error())
	}
	bs, err := json.Marshal(l)
	if err != nil {
		log.Fatalf("Error while json marshal: %v", err.Error())
	}
	w.Write(bs)
}

func (i *inbound) getInputData(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	jsonErr := json.Unmarshal(body, &i)
	if err != nil {
		return jsonErr
	}
	return nil
}
