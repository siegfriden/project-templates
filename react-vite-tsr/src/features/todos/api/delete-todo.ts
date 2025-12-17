import { useMutation, useQueryClient } from '@tanstack/react-query'
import type { MutationConfig } from '@/lib/react-query'

async function deleteTodo({ id }: { id: string }) {
  const response = await fetch(
    `https://jsonplaceholder.typicode.com/todos/${id}`,
    {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    },
  )
  if (!response.ok) {
    throw new Error('Failed to delete todo.')
  }
  return response.json()
}

export function useDeleteTodo({
  mutationConfig,
}: {
  mutationConfig?: MutationConfig<typeof deleteTodo>
} = {}) {
  const queryClient = useQueryClient()
  const { onSuccess, ...restConfig } = mutationConfig || {}

  return useMutation({
    onSuccess: (...args) => {
      queryClient.invalidateQueries({ queryKey: ['todos'] })
      onSuccess?.(...args)
    },
    ...restConfig,
    mutationFn: deleteTodo,
  })
}
