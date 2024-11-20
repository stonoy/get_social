import React, { useEffect } from 'react'
import { followPerson, getUsers } from '../feature/user/userSlice'
import { useDispatch, useSelector } from 'react-redux'

const Users = () => {
  const {loading, submitting, searchUsers, profile} = useSelector(state => state.user)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(getUsers())
  },[])

 

  return (
    <section className='w-11/12 md:1/2 mx-auto shadow-lg h-screen'>
      {
        searchUsers?.map(user => {
           
 
          return (
            <article key={user.id} className='flex justify-between items-center border-b-2 border-gray-300 mt-2 md:mt-4 p-2 md:p-4'>
              <div className='flex gap-2 justify-start items-center'>
                        <h1 className='text-2xl rounded-full bg-slate-300 capitalize cursor-pointer px-4 py-2 font-semibold'>{user.name[0]}</h1>
                        <div className='flex flex-col  items-start'>
                            <h1 className='text-lg text-slate-700'>{user.name}</h1>
                            <h1 className='text-md text-slate-600'>@{user.username}</h1>
                        </div>
                    </div>
                    {
                      profile?.person_i_follow.some(searchUser => searchUser.user_id == user.id) ?
                      <button disabled className='py-2 px-4 bg-green-300 text-white rounded-lg'>followed</button>
                      :
                      <button disabled={submitting} onClick={() => dispatch(followPerson(user.id))} className='py-2 px-4 bg-green-400 text-white rounded-lg'>follow</button>
                    }
            </article>
          )
        })
      }
    </section>
  )
}

export default Users