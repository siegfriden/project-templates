import { useState, type FormEvent } from 'react'
import { useCreateTodo } from '../api/create-todo'

export function CreateTodo() {
  const [title, setTitle] = useState('')
  const createTodo = useCreateTodo({
    mutationConfig: {
      onSuccess: (data) => alert('Todo created: ' + JSON.stringify(data)),
    },
  })

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault()
    createTodo.mutate({ data: { title, completed: false, userId: 1 } })
  }

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Title"
        required
      />
      <button type="submit" disabled={createTodo.isPending}>
        {createTodo.isPending ? 'Creating...' : 'Create'}
      </button>
    </form>
  )
}
