package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//ToDo mock data
type ToDo struct {
	List []List `json:"to_do_list"`
}

//List implements to-do list
type List struct {
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

//Task struct
type Task struct {
	TaskName string `json:"taskName"`
	Status   string `json:"status"`
}

//GetTodoList function
func (l *List) GetTodoList(f *os.File) error {
	err := l.readFile(f)
	if err != nil {
		return err
	}
	return nil
}

//AddTask function
func (l *List) AddTask(f *os.File) error {
	err := l.updateFile(f, "add")
	if err != nil {
		return err
	}
	return nil
}

//RemoveTask function
func (l *List) RemoveTask(f *os.File) error {
	err := l.updateFile(f, "remove")
	if err != nil {
		return err
	}
	return nil
}

//UpdateTask function
func (l *List) UpdateTask(f *os.File) error {
	err := l.updateFile(f, "update")
	if err != nil {
		return err
	}
	return nil
}

func (l *List) readFile(f *os.File) error {
	var t ToDo
	data, err := ioutil.ReadAll(f)
	f.Seek(0, 0)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	for _, v := range t.List {
		if v.Name == l.Name {
			l.Tasks = v.Tasks
		}
	}
	if len(l.Tasks) == 0 {
		return fmt.Errorf("UserName %v not found", l.Name)
	}
	return nil
}

func (l *List) updateFile(f *os.File, opt string) error {
	var t ToDo
	data, err := ioutil.ReadAll(f)
	f.Seek(0, 0)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	for j, v := range t.List {
		if v.Name == l.Name {
			switch opt {
			case "add":
				t.List[j].Tasks = append(t.List[j].Tasks, Task{l.Tasks[0].TaskName, "pending"})
			case "remove":
				for i, v1 := range v.Tasks {
					if l.Tasks[0].TaskName == v1.TaskName {
						t.List[j].Tasks = append(t.List[j].Tasks[:i], t.List[j].Tasks[i+1:]...)
						break
					}
				}
			case "update":
				for i, v1 := range v.Tasks {
					if l.Tasks[0].TaskName == v1.TaskName {
						t.List[j].Tasks[i].Status = l.Tasks[0].Status
						break
					}
				}
			}
			break
		}
	}
	bs, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./data_mock/to_do.json", bs, 0644)
	if err != nil {
		return err
	}

	return nil
}
