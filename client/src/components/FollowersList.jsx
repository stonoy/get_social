import React, { useEffect } from 'react'
import {useDispatch, useSelector} from "react-redux" 
import { getFollowSuggestions, followPerson } from '../feature/user/userSlice'

const FollowersList = ({isBigScreen}) => {
  const {submitting, followSuggestions} = useSelector(state => state.user)
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(getFollowSuggestions())
  },[])


  if (isBigScreen) {
    return (
      <section className="hidden md:block">FollowersList</section>
    )
  }

  return (
    <section className="flex gap-4 md:hidden">
      {
        followSuggestions?.map(follow => {
          return (
            <div className='m-2 p-2 shadow-lg hover:shadow-xl' key={follow.person_id}>
              <div className='flex gap-2 justify-start items-center'>
                        <h1 className='text-xl rounded-full bg-slate-300 capitalize px-4 py-2 font-semibold'>{follow.name[0]}</h1>
                        <div className='flex flex-col  items-start'>
                            <h1 className='text-md font-semibold text-slate-700'>{follow.name}</h1>
                            <h1 className='text-sm text-slate-600'>followers <span>{follow.followers}</span></h1>
                        </div>
                    </div>
                    <button disabled={submitting} onClick={() => dispatch(followPerson(follow.person_id))} className='m-2 py-0.5 px-1 bg-green-400 text-white rounded-md'>follow</button>
            </div>
          )
        })
      }
    </section>
  )
}

export default FollowersList