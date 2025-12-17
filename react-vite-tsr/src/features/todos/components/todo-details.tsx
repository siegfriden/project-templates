import { Link } from '@tanstack/react-router'
import { useGetTodo } from '../api/get-todo'
import { useUpdateTodo } from '../api/update-todo'
import { useDeleteTodo } from '../api/delete-todo'

export function TodoDetails({
  id,
  returnPage,
}: {
  id: string
  returnPage?: number
}) {
  const { data, isLoading, error } = useGetTodo({ id })
  const updateTodo = useUpdateTodo({
    mutationConfig: {
      onSuccess: (data) => alert('Todo updated: ' + JSON.stringify(data)),
    },
  })
  const deleteTodo = useDeleteTodo({
    mutationConfig: {
      onSuccess: () => alert('Todo deleted: ' + JSON.stringify(data)),
    },
  })

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error: {error.message}</div>
  if (!data) return null

  const handleToggleCompleted = () => {
    updateTodo.mutate({ id, data: { ...data, completed: !data.completed } })
  }
  const handleDeleteTodo = () => {
    deleteTodo.mutate({ id })
  }

  return (
    <>
      <h1>Todo Details</h1>
      <Link to="/todos" search={{ page: returnPage }}>
        Back
      </Link>
      <div>
        <h3>{data.title}</h3>
        <p>Status: {data.completed ? 'Completed' : 'Pending'}</p>
        <button onClick={handleToggleCompleted}>
          Mark as {data.completed ? 'Pending' : 'Completed'}
        </button>
        <button onClick={handleDeleteTodo}>Delete</button>
      </div>
    </>
  )
}
