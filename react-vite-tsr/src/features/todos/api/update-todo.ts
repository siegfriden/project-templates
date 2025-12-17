import { useMutation, useQueryClient } from '@tanstack/react-query'
import type { MutationConfig } from '@/lib/react-query'
import type { Todo } from '@/types/api'

type UpdateTodoInput = Omit<Todo, 'id'>

async function updateTodo({ id, data }: { id: string; data: UpdateTodoInput }) {
  const response = await fetch(
    `https://jsonplaceholder.typicode.com/todos/${id}`,
    {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    },
  )
  if (!response.ok) {
    throw new Error('Failed to update todo.')
  }
  return response.json()
}

export function useUpdateTodo({
  mutationConfig,
}: {
  mutationConfig?: MutationConfig<typeof updateTodo>
} = {}) {
  const queryClient = useQueryClient()
  const { onSuccess, ...restConfig } = mutationConfig || {}

  return useMutation({
    onSuccess: (data, ...args) => {
      queryClient.refetchQueries({ queryKey: ['todos', data.id] })
      onSuccess?.(data, ...args)
    },
    ...restConfig,
    mutationFn: updateTodo,
  })
}
