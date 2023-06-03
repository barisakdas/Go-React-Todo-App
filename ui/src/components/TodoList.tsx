// TodoList.tsx
import React from "react";
import { Todo } from "../model/Todo";

interface TodoListProps {
  todos: Todo[];
}

export const TodoList: React.FC<TodoListProps> = ({ todos }) => {
  return (
    <ul>
      {todos.map((todo) => (
        <li key={todo.Id}>
          <h3>{todo.Title}</h3>
          {todo.Description && <p>{todo.Description}</p>}
          <p>Duration: {todo.Duration} minutes</p>
          <p>End Date: {todo.EndDate}</p>
        </li>
      ))}
    </ul>
  );
};
