package cmd

import (
	"ToDoApp/domain/entity"
	"ToDoApp/infrastructure/auth"
	"ToDoApp/presentation"
	"net/http"
)

func Run() {
	entity.CreateDb() // create new db if not exist
	entity.DbMigration()

	http.HandleFunc("/", presentation.GetAllEndpoints)        // For endpoints
	http.HandleFunc("/get-access-token", auth.GetAccessToken) // For access token

	// with access token
	// http.Handle("/get_all_todos", auth.IsAuthorized(presentation.GetAllTodos))
	// http.Handle("/get_todo", auth.IsAuthorized(presentation.GetTodoById))
	// http.Handle("/add_todo", auth.IsAuthorized(presentation.AddTodo))
	// http.Handle("/update_todo", auth.IsAuthorized(presentation.UpdateTodo))
	// http.Handle("/complete_todo", auth.IsAuthorized(presentation.CompleteTodo))
	// http.Handle("/delete_todo", auth.IsAuthorized(presentation.DeleteTodo))

	// without access token
	http.HandleFunc("/get_all_todos", presentation.GetAllTodos)
	http.HandleFunc("/get_todo", presentation.GetTodoById)
	http.HandleFunc("/add_todo", presentation.AddTodo)
	http.HandleFunc("/update_todo", presentation.UpdateTodo)
	http.HandleFunc("/complete_todo", presentation.CompleteTodo)
	http.HandleFunc("/delete_todo", presentation.DeleteTodo)

	http.ListenAndServe(":8080", nil)
}
