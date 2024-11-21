import React, { useEffect } from 'react'
import { Header } from '../components'
import { Outlet } from 'react-router-dom'
import { useDispatch, useSelector } from 'react-redux'
import { getProfile } from '../feature/user/userSlice'

const HomeLayOut = () => {
  const {user:{id}} = useSelector((state) => state.user)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(getProfile(id))
  }, [])

  return (
    <main>
      <Header/>
      <section className='md:w-10/12 md:mx-auto md:flex md:gap-4 md:justify-center'>
        <Outlet/>
      </section>
    </main>
  )
}

export default HomeLayOut
