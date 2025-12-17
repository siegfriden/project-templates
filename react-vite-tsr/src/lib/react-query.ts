import type {
  QueryClientConfig,
  UseMutationOptions,
  UseQueryOptions,
} from '@tanstack/react-query'

export const queryClientConfig: QueryClientConfig = {
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
      retry: false,
      staleTime: 1000 * 60, // 1 minute
    },
  },
}

/* eslint-disable @typescript-eslint/no-explicit-any */

export type QueryConfig<
  QueryFnType extends (...args: any) => Promise<unknown>,
> = Omit<
  UseQueryOptions<
    Awaited<ReturnType<QueryFnType>>,
    Error,
    Awaited<ReturnType<QueryFnType>>,
    any[]
  >,
  'queryKey' | 'queryFn'
>

export type MutationConfig<
  MutationFnType extends (...args: any) => Promise<unknown>,
> = Omit<
  UseMutationOptions<
    Awaited<ReturnType<MutationFnType>>,
    Error,
    Parameters<MutationFnType>[0]
  >,
  'mutationFn'
>

/* eslint-enable @typescript-eslint/no-explicit-any */
