import React from 'react'
import UserNameFirst from './UserNameFirst'
import { useSelector } from 'react-redux'

const User = () => {
  const {profile} = useSelector(state => state.user)

  return (
    <section className="hidden md:block">
      <div className='flex w-full gap-2 justify-start items-center shadow-lg p-4'>
                            <UserNameFirst letter={profile?.user?.name[0]} id={profile?.user?.id}/>
                            <div className='flex flex-col  items-start'>
                                
                                <h1 className='text-lg font-semibold text-slate-700'>{profile?.user?.name}</h1>
                                
                                <div>
                                <span className='text-sm text-slate-600'>follows {profile?.person_i_follow?.length}</span>
                                <span> , </span>
                                <span className='text-sm text-slate-600'>followers {profile?.my_followers?.length}</span>
                                </div>
                            </div>
                            
                        </div>
    </section>
  )
}

export default User