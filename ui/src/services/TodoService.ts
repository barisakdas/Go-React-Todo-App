import axios, { AxiosResponse } from 'axios';
import { Todo } from '../model/Todo';

const API_URL = 'http://localhost:8080'; // API URL'nizi buraya girin

export const getTodoList = async (): Promise<Todo[]> => {
  try {
    console.log("Get Todo List")
    const response: AxiosResponse<Todo[]> = await axios.get<Todo[]>(`${API_URL}/get_all_todos`);
    return response.data;
  } catch (error) {
    console.error('Error while fetching todo list:', error);
    return [];
  }
};

export const createTodo = async (todo: Todo): Promise<Todo | null> => {
  try {
    console.log("Create Todo")
    const response: AxiosResponse<Todo> = await axios.post<Todo>(`${API_URL}/add_todo`, todo);
    return response.data;
  } catch (error) {
    console.error('Error while creating todo:', error);
    return null;
  }
};

export const completeTodo = async (todoId: number, updatedTodo: Todo): Promise<Todo | null> => {
  try {
    console.log("Complete Todo")
    const response: AxiosResponse<Todo> = await axios.put<Todo>(`${API_URL}/complete_todo?id=/${todoId}`, updatedTodo);
    return response.data;
  } catch (error) {
    console.error('Error while updating todo:', error);
    return null;
  }
};

export const deleteTodo = async (todoId: number): Promise<Todo | null> => {
  try {
    console.log("Delete Todo")
    const response: AxiosResponse<Todo> = await axios.delete<Todo>(`${API_URL}/delete_todo?id=/${todoId}`);
    return response.data;
  } catch (error) {
    console.error('Error while deleting todo:', error);
    return null;
  }
};
