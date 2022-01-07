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
    GetAllTodosAsync();
  }, []);

  const GetAllTodosAsync = async () => {
    try {
      const response = await axios.get("http://localhost:8080/get_all_todos");
      const data = await response.data;
      setTodos(data);
    } catch (err) {
      console.error(err);
    }
  };

  // this function get all data from api.
  async function CreateTodoAsync() {
    var body =
      `{
      "Title": "` +
      title +
      `",
      "Description": "` +
      description +
      `",
      "Duration": ` +
      duration +
      `,
      "IsCompleted": false
    }`;

    console.log("body: ", body);

    const response = await fetch("http://localhost:8080/add_todo", {
      method: "post",
      mode: "no-cors",
      body: body,
    });
    GetAllTodosAsync();
    return response.json();
  }

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
            onChange={(event) => setTitle(event.target.value)}
          />
        </div>
        <div className="add-todo-item">
          <label>Description: </label>
          <input
            type="text"
            name="title"
            id="description"
            onChange={(event) => setDescription(event.target.value)}
          />
        </div>
        <div className="add-todo-item">
          <label>Duration: </label>
          <input
            type="number"
            name="title"
            id="duration"
            onChange={(event) => setDuration(parseInt(event.target.value))}
          />
        </div>
        <button className="add-todo-item-button" onClick={CreateTodoAsync}>
          Add New
        </button>
      </div>
    </div>
  );
}

export default App;
