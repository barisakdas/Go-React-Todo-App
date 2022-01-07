package application

import . "ToDoApp/domain/repository"

type TodoApplication struct{}

var todoRepo = NewTodoRepository()

func (t TodoApplication) GetAllTodos() ([]TodoDTO, error) {
	return todoRepo.GetAllTodos()
}
func (t TodoApplication) GetTodoById(id int) (TodoDTO, error) {
	return todoRepo.GetTodoById(id)
}
func (t TodoApplication) AddTodo(todoDto *TodoDTO) (TodoDTO, error) {
	return todoRepo.AddTodo(todoDto)
}
func (t TodoApplication) UpdateTodo(todoDto *TodoDTO) (TodoDTO, error) {
	return todoRepo.UpdateTodo(todoDto)
}
func (t TodoApplication) CompleteTodo(id int) (TodoDTO, error) {
	data, err := todoRepo.GetTodoById(id)
	if err != nil {
		return TodoDTO{}, err
	}
	data.IsCompleted = true
	return todoRepo.UpdateTodo(&data)
}
func (t TodoApplication) DeleteTodo(id int) error {
	return todoRepo.DeleteTodo(id)
}
