import { queryOptions, useQuery } from '@tanstack/react-query'
import type { QueryConfig } from '@/lib/react-query'
import type { Todo } from '@/types/api'

async function getTodo(id: string): Promise<Todo> {
  const response = await fetch(
    `https://jsonplaceholder.typicode.com/todos/${id}`,
  )
  if (!response.ok) {
    throw new Error('Failed to fetch todo.')
  }
  return await response.json()
}

export function getTodoQueryOptions({ id }: { id: string }) {
  return queryOptions({
    queryKey: ['todos', id],
    queryFn: () => getTodo(id),
  })
}

export function useGetTodo({
  id,
  queryConfig,
}: {
  id: string
  queryConfig?: QueryConfig<typeof getTodo>
}) {
  return useQuery({
    ...getTodoQueryOptions({ id }),
    ...queryConfig,
  })
}
