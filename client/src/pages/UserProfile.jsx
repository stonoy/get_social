import React, { useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { LuPenLine } from "react-icons/lu";
import { changeUserDetails, updateUser } from '../feature/user/userSlice';

const UserProfile = () => {
  const {profile, submitting} = useSelector(state => state.user)
  const [openUpdate, setOpenUpdate] = useState(false)
  const dispatch = useDispatch()

  const handleUpdate = (e) => {
    const {name, value} = e.target

    dispatch(changeUserDetails({name, value}))
  }

  const updateUserAsync = () => {
    dispatch(updateUser()).then(() => setOpenUpdate(false))
  }
  
  // console.log(profile)

  return (
    <section className='w-11/12 md:1/2 mx-auto p-4 md:p-8 flex flex-col gap-4 items-start'>
      <div className='flex w-full gap-2 justify-between items-center'>
                        <h1 className='text-2xl rounded-full bg-slate-300 capitalize px-4 py-2 font-semibold'>{profile?.user?.name[0]}</h1>
                        <div className='flex flex-col  items-start'>
                            {openUpdate && <label className='font-semibold text-md text-slate-700'>Name</label>}
                            {!openUpdate && <h1 className='text-xl font-semibold text-slate-700'>{profile?.user?.name}</h1>}
                            {openUpdate && <input type='text' value={profile?.user?.name} onChange={handleUpdate} name='name' className='block border-2'/>}
                            <div>
                            <span className='text-md text-slate-600'>follows {profile?.person_i_follow?.length}</span>
                            <span> , </span>
                            <span className='text-md text-slate-600'>followers {profile?.my_followers?.length}</span>
                            </div>
                        </div>
                        <h1 onClick={() => setOpenUpdate(prev => !prev)} className='ml-auto text-lg'><LuPenLine/></h1>
                    </div>
                    {/* other details */}
                    <div>
                      <label className='font-semibold text-lg text-slate-700'>Location</label>
                      {!openUpdate && <h1 className='text-md text-slate-500 '>{profile?.user?.location}</h1>}
                      {openUpdate && <input type='text' value={profile?.user?.location} onChange={handleUpdate} name='location' className='block border-2'/>}
                    </div>
                    <div>
                      <label className='font-semibold text-lg text-slate-700'>Age</label>
                      {!openUpdate && <h1 className='text-md text-slate-500 '>{profile?.user?.age}</h1>}
                      {openUpdate && <input type='number' value={profile?.user?.age} onChange={handleUpdate} name='age' className='block border-2'/>}
                    </div>
                    <div>
                      <label className='font-semibold text-lg text-slate-700'>Username</label>
                      {!openUpdate && <h1 className='text-md text-slate-500 '>{profile?.user?.username}</h1>}
                      {openUpdate && <input type='text' value={profile?.user?.username} onChange={handleUpdate} name='username' className='block border-2'/>}
                    </div>
                    <div>
                      <label className='font-semibold text-lg text-slate-700'>Bio</label>
                      {!openUpdate && <h1 className='text-md text-slate-500 '>{profile?.user?.bio}</h1>}
                      {openUpdate && <input type='text' value={profile?.user?.bio} onChange={handleUpdate} name='bio' className='block border-2'/>}
                    </div>
                    <button disabled={submitting || !openUpdate} onClick={updateUserAsync} className='ml-auto py-2 px-4 text-md bg-green-400 text-white rounded-md'>Update</button>
    </section>
  )
}

export default UserProfile