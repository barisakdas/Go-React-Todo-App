package presentation

import (
	"ToDoApp/application"
	"ToDoApp/domain/repository"
	"ToDoApp/infrastructure/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	data, err := application.TodoApplication{}.GetAllTodos()
	if (err != nil) || (data == nil) {
		util.HttpNotFound(w)
		return
	}
	jsonResp, _ := json.Marshal(&data)
	w.Write(jsonResp)
}
func GetTodoById(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	id := r.URL.Query().Get("id")
	todoId, _ := strconv.Atoi(id)
	data, err := application.TodoApplication{}.GetTodoById(todoId)
	if (err != nil) || ((repository.TodoDTO{}) == data) {
		util.HttpNotFound(w)
		return
	}

	jsonResp, _ := json.Marshal(&data)
	w.Write(jsonResp)
}
func AddTodo(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	var todoDto repository.TodoDTO
	bodyText, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(bodyText, &todoDto)

	fmt.Println("bodyText: ",todoDto)
	data, err := application.TodoApplication{}.AddTodo(&todoDto)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Data added: " + data.Title))
}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todoDto repository.TodoDTO
	bodyText, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(bodyText, &todoDto)

	data, err := application.TodoApplication{}.UpdateTodo(&todoDto)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Data updated: " + data.Title))
}
func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	todoId, _ := strconv.Atoi(id)

	data, err := application.TodoApplication{}.CompleteTodo(todoId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Data updated: " + data.Title))
}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	todoId, _ := strconv.Atoi(id)
	application.TodoApplication{}.DeleteTodo(todoId)
	w.Write([]byte("Deleting operation is success"))
}
