import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import {Login, Register, HomeLayOut, ErrorPage, Landing, TimeLine, Users} from '../src/pages'
import { User } from './components'


const router = createBrowserRouter([
  {
    path: "/",
    element: <HomeLayOut/>,
    errorElement: <ErrorPage/>,
    children: [
      {
        index: true,
        element: <Landing/>
      },
      {
        path: "/timeline",
        element: <TimeLine/>
      },
      {
        path: "/searchusers",
        element: <Users/>
      }
    ]
  },
  {
    path: "/login",
    element: <Login/>, 
  },
  {
    path: "/register",
    element: <Register/>,
  }
])


export default function App() {
  return <RouterProvider router={router} />
}