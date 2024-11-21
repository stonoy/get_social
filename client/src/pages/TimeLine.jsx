import React, { useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import Posts from '../components/Posts'
import Shimmer from '../components/Shimmer'
import Pagination from '../components/Pagination'
import { getPosts } from '../feature/posts/postsSlice'

const TimeLine = () => {
    const {loading, posts,numOfPages,page} = useSelector(state => state.posts)
    const dispatch = useDispatch()
  
    useEffect(() => {
      dispatch(getPosts(`/getpostsbyuser`))
    }, [])

    
    return (
      <section className='p-4 bg-gray-200 h-screen md:w-3/4'>
          <div className=''>
          
          <div>
              {loading ? <Shimmer/>
              :
              <Posts posts={posts} isTimeLine={true}/>
            }
          </div>
          <Pagination path="/getpostsbyuser" numOfPages={numOfPages} page={page}/>
          </div>
      </section>
    )
}

export default TimeLine