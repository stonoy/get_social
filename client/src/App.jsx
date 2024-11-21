import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import {Login, Register, HomeLayOut, ErrorPage, Landing, TimeLine, Users, UserProfile, AnyProfile} from '../src/pages'


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
      },
      {
        path: "/profile",
        element: <UserProfile/>
      },
      {
        path: "/profile/:userId",
        element: <AnyProfile/>
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