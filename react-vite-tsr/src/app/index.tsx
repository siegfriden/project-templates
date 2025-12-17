import { useState } from 'react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { RouterProvider, createRouter } from '@tanstack/react-router'
import { queryClientConfig } from '@/lib/react-query'
import { routeTree } from './routeTree.gen'

const router = createRouter({
  routeTree,
  context: {
    queryClient: undefined!, // will be set later in RouterProvider
  },
  scrollRestoration: true,
  defaultPreload: 'intent', // preload pages when the user hovers over a link
})

// Register the router instance for type safety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

export function App() {
  // Initialize queryClient here to ensure isolation when using SSR
  // https://tanstack.com/query/latest/docs/framework/react/guides/ssr
  const [queryClient] = useState(() => new QueryClient(queryClientConfig))

  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} context={{ queryClient }} />
      <ReactQueryDevtools />
    </QueryClientProvider>
  )
}
