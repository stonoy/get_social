import React, { useState } from 'react'
import { GiHamburgerMenu } from "react-icons/gi";
import { useDispatch, useSelector } from 'react-redux';
import { getUsers, setSearchName } from '../feature/user/userSlice';
import { useNavigate } from 'react-router-dom';

const Header = () => {
    const {token, profile:{user: {name}}, searchName} = useSelector((state) => state.user)
    const dispatch = useDispatch()
    const navigate = useNavigate()

    const handleSearch = () => {
      if (searchName == ""){return}
      dispatch(getUsers({name: searchName})).then(() => navigate("/searchusers"))
    }

  return (
    <nav className='w-full  border-b-2 border-slate-500 text-xl text-slate-600'>
      <div className='flex justify-between items-center w-11/12 md:w-9/12 mx-auto'>
      <div className='flex gap-2'>
        <input className='w-3/4 border-2' type='text' name='name' value={searchName} onChange={(e) => dispatch(setSearchName(e.target.value))}/>
        <button onClick={handleSearch} className='py-1 px-2 text-slate-600 bg-gray-200 rounded-sm'>Search</button>
      </div>
        <div className='p-4 border-2 flex justify-end items-center gap-2 md:p-4'>
            <h1 >{name}</h1>
            <GiHamburgerMenu/>
        </div>
      </div>
    </nav>
  )
}

export default Header