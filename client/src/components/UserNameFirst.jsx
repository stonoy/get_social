import React from 'react'
import { Link } from 'react-router-dom'

const UserNameFirst = ({letter, id}) => {
  return (
    <Link to={`/profile/${id}`} className='text-2xl rounded-full bg-slate-300 capitalize px-4 py-2 font-semibold'>{letter}</Link>
  )
}

export default UserNameFirst