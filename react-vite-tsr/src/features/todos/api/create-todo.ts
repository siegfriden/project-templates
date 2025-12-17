import { useMutation, useQueryClient } from '@tanstack/react-query'
import type { MutationConfig } from '@/lib/react-query'
import type { Todo } from '@/types/api'

type CreateTodoInput = Omit<Todo, 'id'>

async function createTodo({ data }: { data: CreateTodoInput }) {
  const response = await fetch('https://jsonplaceholder.typicode.com/todos', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
  if (!response.ok) {
    throw new Error('Failed to create todo.')
  }
  return response.json()
}

export function useCreateTodo({
  mutationConfig,
}: {
  mutationConfig?: MutationConfig<typeof createTodo>
} = {}) {
  const queryClient = useQueryClient()
  const { onSuccess, ...restConfig } = mutationConfig || {}

  return useMutation({
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: ['todos'] })
      onSuccess?.(...args)
    },
    ...restConfig,
    mutationFn: createTodo,
  })
}
