import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import {Login, Register, HomeLayOut, ErrorPage, Landing} from '../src/pages'


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