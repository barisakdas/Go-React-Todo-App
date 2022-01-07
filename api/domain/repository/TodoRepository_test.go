package repository

import (
	"reflect"
	"testing"
)

func TestTodoRepository_GetAllTodos(t *testing.T) {
	var tests []struct {
		name    string
		tr      TodoRepository
		want    []TodoDTO
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.GetAllTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.GetAllTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.GetAllTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_GetTodoById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		tr      TodoRepository
		args    args
		want    TodoDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.GetTodoById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.GetTodoById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.GetTodoById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_AddTodo(t *testing.T) {
	type args struct {
		todoDto *TodoDTO
	}
	tests := []struct {
		name    string
		tr      TodoRepository
		args    args
		want    TodoDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.AddTodo(tt.args.todoDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.AddTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.AddTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_UpdateTodo(t *testing.T) {
	type args struct {
		todoDto *TodoDTO
	}
	tests := []struct {
		name    string
		tr      TodoRepository
		args    args
		want    TodoDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.UpdateTodo(tt.args.todoDto)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TodoRepository.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodoRepository_DeleteTodo(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		tr      TodoRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.DeleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
