# React TS + TanStack Router Template

This template is built on top of Vite's React TS template.

```bash
npm create vite@latest -- --template react-ts
```

## Tech Stack

- **Framework:** [React 19](https://react.dev/) + [Vite](https://vitejs.dev/)
- **Routing:** [TanStack Router](https://tanstack.com/router) (File-based routing)
- **State Management:** [TanStack Query](https://tanstack.com/query)

## Project Structure

The project structure is inspired by [Bulletproof React](https://github.com/alan2207/bulletproof-react).

```text
src/
├── app/                 # App-layer configuration (providers, router)
│   └── routes/          # File-based route definitions
├── features/            # Feature-based modules (api, components, hooks per feature)
├── lib/                 # Library configurations (axios, query-client, etc.)
├── types/               # Shared TS types
└── ...
```
