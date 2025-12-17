import { createFileRoute } from '@tanstack/react-router'
import { getTodosQueryOptions } from '@/features/todos/api/get-todos'
import { TodoList } from '@/features/todos/components/todo-list'

type RouteSearch = {
  page?: number
}

export const Route = createFileRoute('/todos/')({
  component: RouteComponent,
  validateSearch: (search: Record<string, unknown>): RouteSearch => {
    const pageNum = Number(search.page)
    return { page: pageNum >= 1 ? Math.floor(pageNum) : undefined }
  },

  // `loaderDeps` is required to enable search params in loaders.
  // This is designed this way so the router can track loader dependencies and prevent caching bugs,
  // ensuring that preloaded data is stored separately for each unique page.
  // https://tanstack.com/router/latest/docs/framework/react/guide/data-loading#using-search-params-in-loaders
  loaderDeps: ({ search }) => ({ page: search.page }),
  loader: ({ context, deps }) => {
    return context.queryClient.ensureQueryData(
      getTodosQueryOptions({ page: deps.page }),
    )
  },
})

function RouteComponent() {
  const { page } = Route.useSearch()

  return <TodoList page={page} />
}
