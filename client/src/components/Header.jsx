import React from 'react'
import { GiHamburgerMenu } from "react-icons/gi";
import { useSelector } from 'react-redux';

const Header = () => {
    const {token, user:{name}} = useSelector((state) => state.user)
  return (
    <nav className='w-full border-b-2 border-slate-500 text-xl text-slate-600'>
        <div className='p-4 border flex justify-end items-center gap-2 md:p-4 md:w-5/6 md:mx-auto'>
            <h1 >{name}</h1>
            <GiHamburgerMenu/>
        </div>
    </nav>
  )
}

export default Header