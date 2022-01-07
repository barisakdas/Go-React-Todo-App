package repository

import (
	. "ToDoApp/domain/entity"
	"time"
)

type ITodoRepository interface {
	GetAllTodos() ([]TodoDTO, error)
	GetTodoById(id int) (TodoDTO, error)
	AddTodo(todoDto *TodoDTO) (TodoDTO, error)
	UpdateTodo(todoDto *TodoDTO) (TodoDTO, error)
	DeleteTodo(id int) error
}

type TodoRepository struct{}

type TodoDTO struct {
	Id          int
	Title       string
	Description string
	CreatedDate string
	Duration    int
	EndDate     string
	IsCompleted bool
}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (t TodoRepository) GetAllTodos() ([]TodoDTO, error) {
	data, err := Todo{}.GetAll()
	if err != nil {
		return nil, err
	}

	var dataDtos []TodoDTO

	for _, value := range data {
		newData := TodoDTO{
			Id:          value.Id,
			Title:       value.Title,
			Description: value.Description,
			CreatedDate: value.CreatedDate,
			EndDate:     value.EndDate,
			IsCompleted: value.IsCompleted,
		}
		dataDtos = append(dataDtos, newData)
	}
	return dataDtos, nil
}

func (t TodoRepository) GetTodoById(id int) (TodoDTO, error) {
	data, err := Todo{}.Get(id)
	if err != nil {
		return TodoDTO{}, err
	}
	return TodoDTO{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		CreatedDate: data.CreatedDate,
		EndDate:     data.EndDate,
		IsCompleted: data.IsCompleted,
	}, nil
}

func (t TodoRepository) AddTodo(todoDto *TodoDTO) (TodoDTO, error) {
	result, err := Todo{
		Title:       todoDto.Title,
		Description: todoDto.Description,
		CreatedDate: time.Now().Format("02-01-2006"),
		EndDate:     time.Now().AddDate(0, 0, todoDto.Duration).Format("02-01-2006"),
		IsCompleted: false,
		IsDeleted:   false,
	}.Add()
	return TodoDTO{
		Id:          result.Id,
		Title:       result.Title,
		Description: result.Description,
		CreatedDate: result.CreatedDate,
		EndDate:     result.EndDate,
		IsCompleted: result.IsCompleted,
	}, err
}

func (t TodoRepository) UpdateTodo(todoDto *TodoDTO) (TodoDTO, error) {
	data, err := Todo{}.Get(todoDto.Id)
	if err != nil {
		return TodoDTO{}, err
	}
	result, err := data.Update(Todo{
		Id:          todoDto.Id,
		Title:       todoDto.Title,
		Description: todoDto.Description,
		CreatedDate: todoDto.CreatedDate,
		EndDate:     todoDto.EndDate,
		IsCompleted: todoDto.IsCompleted,
	})
	return TodoDTO{
		Id:          result.Id,
		Title:       result.Title,
		Description: result.Description,
		CreatedDate: result.CreatedDate,
		EndDate:     result.EndDate,
		IsCompleted: result.IsCompleted,
	}, err
}

func (t TodoRepository) DeleteTodo(id int) error {
	data, err := Todo{}.Get(id)
	if err != nil {
		return err
	}
	return data.Delete()
}
