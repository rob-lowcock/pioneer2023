import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import Root from './routes/root'
import Login, {
  action as loginAction
} from './routes/login'
import Retro, { loader as retroLoader, action as retroAction } from './routes/retro'
import './index.css'

const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
  },
  {
    path: '/login',
    element: <Login />,
    action: loginAction,
  },
  {
    path: '/retro',
    loader: retroLoader,
    action: retroAction,
    element: <Retro />,
  }
])

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)
