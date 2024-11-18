import React from 'react'
import { Header } from '../components'
import { Outlet } from 'react-router-dom'

const HomeLayOut = () => {
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
