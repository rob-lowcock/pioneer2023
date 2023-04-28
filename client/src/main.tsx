import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import Root from './routes/root'
import Login from './routes/login'
import './index.css'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
  },
  {
    path: '/login',
    element: <Login />,
  }
])

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
