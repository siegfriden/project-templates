import { Link } from '@tanstack/react-router'
import { useGetTodos } from '../api/get-todos'
import { CreateTodo } from './create-todo'

export function TodoList({ page = 1 }: { page?: number }) {
  const { data, error, isLoading } = useGetTodos({ page })

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error: {error.message}</div>

  return (
    <>
      <h1>Todos</h1>
      <CreateTodo />
      <ul>
        {data?.map((todo) => (
          <li key={todo.id}>
            <Link
              to="/todos/$id"
              params={{ id: todo.id.toString() }}
              search={{ returnPage: page }}
            >
              {todo.title} {todo.completed ? '(Completed)' : ''}
            </Link>
          </li>
        ))}
      </ul>
      <div style={{ display: 'flex', gap: '10px' }}>
        <Link
          to="/todos"
          search={{ page: Math.max(1, page - 1) }}
          disabled={page <= 1}
        >
          Previous
        </Link>
        <span>Page {page}</span>
        <Link to="/todos" search={{ page: page + 1 }}>
          Next
        </Link>
      </div>
    </>
  )
}
