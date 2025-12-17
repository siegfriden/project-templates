import { queryOptions, useQuery } from '@tanstack/react-query'
import type { QueryConfig } from '@/lib/react-query'
import type { Todo } from '@/types/api'

async function getTodos(page = 1): Promise<Todo[]> {
  const response = await fetch(
    `https://jsonplaceholder.typicode.com/todos?_limit=10&_page=${page}`,
  )
  if (!response.ok) {
    throw new Error('Failed to fetch todos.')
  }
  return await response.json()
}

export function getTodosQueryOptions({ page }: { page?: number } = {}) {
  return queryOptions({
    queryKey: page ? ['todos', { page }] : ['todos'],
    queryFn: () => getTodos(page),
  })
}

export function useGetTodos({
  page,
  queryConfig,
}: {
  page?: number
  queryConfig?: QueryConfig<typeof getTodos>
} = {}) {
  return useQuery({
    ...getTodosQueryOptions({ page }),
    ...queryConfig,
  })
}
