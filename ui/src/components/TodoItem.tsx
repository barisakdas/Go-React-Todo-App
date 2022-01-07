import React from "react";
import { Todo } from "../model/Todo";
import "../style/TodoItem.scss";

export const TodoItem: React.FC<Todo> = ({
  Id,
  Title,
  Description,
  CreatedDate,
  EndDate,
  IsCompleted,
}) => {
  return (
    <article className="todo-item">
      <section className="todo-item-text">
        <h3>{Title}</h3>
        <p>{Description}</p>
      </section>
      <input type="checkbox" defaultChecked={IsCompleted} />
    </article>
  );
};
