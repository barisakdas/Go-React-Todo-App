import { Todo } from "../model/Todo";
import { TodoItem } from "./TodoItem";
import "../style/TodoList.scss";

type Props = {
  todos: Todo[];
};

export const TodoList: React.FC<Props> = ({ todos }) => {
  return (
    <ul className="todo-list">
      {todos.length > 0 ? (
        todos.map((todo: Todo, i) => (
          <li key={i}>
            <TodoItem
              Id={todo.Id}
              Title={todo.Title}
              Description={todo.Description}
              CreatedDate={todo.CreatedDate}
              EndDate={todo.EndDate}
              IsCompleted={todo.IsCompleted}
            />
          </li>
        ))
      ) : (
        <></>
      )}
    </ul>
  );
};
