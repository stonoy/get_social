import axios from 'axios';
import React, { useEffect, useState } from 'react'
import { ImSpinner9 } from "react-icons/im";
import { axiosBase } from '../utils';
import { useParams } from 'react-router-dom';
import { toast } from 'react-toastify';
import { UserNameFirst } from '../components';

const AnyProfile = () => {
    const [profile, setProfile] = useState(null)
    const {userId} = useParams()

    console.log(userId)

    useEffect(() => {
        const fetchUserProfile = async () => {
           try {
            const resp = await axiosBase.get(`/getusersdetails/${userId}`)
            setProfile(resp?.data)
           } catch (error) {
            toast.error(error?.response?.data?.msg)
           }
        }

        fetchUserProfile()
    }, [userId])

  return (
    profile ? 

    (
        <section className='w-11/12 md:1/2 mx-auto p-4 md:p-8 flex flex-col gap-4 items-start'>
          <div className='flex w-full gap-2 justify-start items-center'>
                            <UserNameFirst letter={profile?.user?.name[0]} id={profile?.user?.id}/>
                            <div className='flex flex-col  items-start'>
                                
                                <h1 className='text-xl font-semibold text-slate-700'>{profile?.user?.name}</h1>
                                
                                <div>
                                <span className='text-md text-slate-600'>follows {profile?.person_i_follow?.length}</span>
                                <span> , </span>
                                <span className='text-md text-slate-600'>followers {profile?.my_followers?.length}</span>
                                </div>
                            </div>
                            
                        </div>
                        {/* other details */}
                        <div>
                          <label className='font-semibold text-lg text-slate-700'>Location</label>
                          <h1 className='text-md text-slate-500 '>{profile?.user?.location}</h1>
                          
                        </div>
                        <div>
                          <label className='font-semibold text-lg text-slate-700'>Age</label>
                          <h1 className='text-md text-slate-500 '>{profile?.user?.age}</h1>
                          
                        </div>
                        <div>
                          <label className='font-semibold text-lg text-slate-700'>Username</label>
                          <h1 className='text-md text-slate-500 '>{profile?.user?.username}</h1>
                         
                        </div>
                        <div>
                          <label className='font-semibold text-lg text-slate-700'>Bio</label>
                          <h1 className='text-md text-slate-500 '>{profile?.user?.bio}</h1>
                         
                        </div>
                        
        </section>
      )
      :
      (
        <div className='flex w-full h-screen justify-center items-center'>
            <ImSpinner9 className='animate-spin'/>
        </div>
      )
  )
}

export default AnyProfile