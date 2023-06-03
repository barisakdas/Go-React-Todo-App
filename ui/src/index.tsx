import React, { useState, useEffect } from 'react';
import { getTodoList, createTodo, completeTodo, deleteTodo } from './services/TodoService';
import { Todo } from './model/Todo';

type TodoFormState = {
  title: string;
  description: string;
  duration: number;
  endDate: string;
};

const App: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTodo, setNewTodo] = useState<TodoFormState>({
    title: '',
    description: '',
    duration: 0,
    endDate: '',
  });

  useEffect(() => {
    fetchTodoList();
  }, []);

  const fetchTodoList = async () => {
  try {
    const todoList: Todo[] = await getTodoList();
    setTodos(todoList);
  } catch (error) {
    console.error('Error while fetching todo list:', error);
  }
};

  const handleAddTodo = async () => {
    if (!newTodo.title.trim() || !newTodo.endDate.trim()) {
      return;
    }

    const todoToAdd: Todo = {
      Id: Date.now(),
      Title: newTodo.title.trim(),
      Description: newTodo.description.trim(),
      CreatedDate: new Date().toISOString(),
      Duration: newTodo.duration,
      EndDate: newTodo.endDate.trim(),
      IsCompleted: false,
    };

    try {
      const createdTodo: Todo | null = await createTodo(todoToAdd);
      if (createdTodo) {
        setTodos((prevTodos) => [...prevTodos, createdTodo]);
        setNewTodo({
          title: '',
          description: '',
          duration: 0,
          endDate: '',
        });
      }
    } catch (error) {
      console.error('Error while creating todo:', error);
    }
  };

  const handleDeleteTodo = async (todoId: number) => {
    try {
      const deletedTodo: Todo | null = await deleteTodo(todoId);
      if (deletedTodo) {
        setTodos((prevTodos) => prevTodos.filter((todo) => todo.Id !== todoId));
      }
    } catch (error) {
      console.error('Error while deleting todo:', error);
    }
  };

  return (
    <div>
      <h2>Todo List</h2>
      <ul>
        {todos.map((todo) => (
          <li key={todo.Id}>
            <h3>{todo.Title}</h3>
            {todo.Description && <p>{todo.Description}</p>}
            <p>Duration: {todo.Duration} minutes</p>
            <p>End Date: {todo.EndDate}</p>
            <button onClick={() => handleDeleteTodo(todo.Id)}>Delete</button>
          </li>
        ))}
      </ul>
      <h2>Add Todo</h2>
      <div>
        <label>Title:</label>
        <input
          type="text"
          value={newTodo.title}
          onChange={(e) => setNewTodo({ ...newTodo, title: e.target.value })}
        />
      </div>
      <div>
        <label>Description:</label>
        <textarea
          value={newTodo.description}
          onChange={(e) => setNewTodo({ ...newTodo, description: e.target.value })}
        />
      </div>
      <div>
        <label>Duration (minutes):</label>
        <input
          type="number"
          value={newTodo.duration}
          onChange={(e) => setNewTodo({ ...newTodo, duration: parseInt(e.target.value) })}
        />
      </div>
      <div>
        <label>End Date:</label>
        <input
          type="text"
          value={newTodo.endDate}
          onChange={(e) => setNewTodo({ ...newTodo, endDate: e.target.value })}
        />
      </div>
      <button onClick={handleAddTodo}>Add</button>
    </div>
  );
};

export default App;
