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
      <section className='md:w-9/12 md:mx-auto'>
        <Outlet/>
      </section>
    </main>
  )
}

export default HomeLayOut
