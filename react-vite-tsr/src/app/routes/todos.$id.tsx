import { createFileRoute } from '@tanstack/react-router'
import { getTodoQueryOptions } from '@/features/todos/api/get-todo'
import { TodoDetails } from '@/features/todos/components/todo-details'

type RouteSearch = {
  returnPage?: number
}

export const Route = createFileRoute('/todos/$id')({
  component: RouteComponent,
  validateSearch: (search: Record<string, unknown>): RouteSearch => {
    const pageNum = Number(search.returnPage)
    return { returnPage: pageNum >= 1 ? Math.floor(pageNum) : undefined }
  },
  loader: ({ context, params }) => {
    return context.queryClient.ensureQueryData(
      getTodoQueryOptions({ id: params.id }),
    )
  },
})

function RouteComponent() {
  const { id } = Route.useParams()
  const { returnPage } = Route.useSearch()

  return <TodoDetails id={id} returnPage={returnPage} />
}
