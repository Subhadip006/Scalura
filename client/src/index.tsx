/* @refresh reload */
import { render } from 'solid-js/web'
import './index.css'
import App from './App.tsx'
import { SolidQueryDevtools } from '@tanstack/solid-query-devtools'
import { QueryClient, QueryClientProvider } from '@tanstack/solid-query'

const root = document.getElementById('root')

render(() => 
    <QueryClientProvider client={new QueryClient()}>
        <App />
        <SolidQueryDevtools />
    </QueryClientProvider>, root!)
