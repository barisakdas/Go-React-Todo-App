import axios from "axios";
import React, { useEffect, useState } from "react";
import "./App.css";
import { Header } from "./components/Header";
import { TodoList } from "./components/TodoList";
import { Todo } from "./model/Todo";

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [title, setTitle] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [duration, setDuration] = useState<number>(1);

  useEffect(() => {
    fetchTodoList();
  }, []);

  const fetchTodoList = async () => {
    try {
      const response = await axios.get<Todo[]>("http://localhost:8080/get_all_todos");
      setTodos(response.data);
    } catch (err) {
      console.error(err);
    }
  };

  const createTodo = async () => {
    const newTodo: Todo = {
      Id: Date.now(),
      Title: title,
      Description: description,
      CreatedDate: new Date().toISOString(),
      Duration: duration,
      EndDate: "",
      IsCompleted: false,
    };

    try {
      await axios.post<Todo>("http://localhost:8080/add_todo", newTodo);
      fetchTodoList();
    } catch (error) {
      console.error("Error while creating todo:", error);
    }
  };

  return (
    <div className="App">
      <Header />
      <TodoList todos={todos} />
      <div className="add-todo-item">
        <div className="add-todo-item">
          <label>Title: </label>
          <input
            type="text"
            name="title"
            id="title"
            value={title}
            onChange={(event) => setTitle(event.target.value)}
          />
        </div>
        <div className="add-todo-item">
          <label>Description: </label>
          <input
            type="text"
            name="title"
            id="description"
            value={description}
            onChange={(event) => setDescription(event.target.value)}
          />
        </div>
        <div className="add-todo-item">
          <label>Duration: </label>
          <input
            type="number"
            name="title"
            id="duration"
            value={duration}
            onChange={(event) => setDuration(parseInt(event.target.value))}
          />
        </div>
        <button className="add-todo-item-button" onClick={createTodo}>
          Add New
        </button>
      </div>
    </div>
  );
}

export default App;
