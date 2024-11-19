import React, { useEffect, useRef } from 'react'
import Shimmer from './Shimmer'
import { useDispatch, useSelector } from 'react-redux'
import { createPost, getPosts } from '../feature/posts/postsSlice'
import Posts from './Posts'
import Pagination from './Pagination'

const Content = () => {
  const {loading, posts,numOfPages,page,posting} = useSelector(state => state.posts)
  const dispatch = useDispatch()
  const postRef = useRef(null)

  useEffect(() => {
    dispatch(getPosts(`/getpostsuggestions`))
  }, [])

  const handlePostSubmit = (e) => {
    e.preventDefault()
    const content = new FormData(postRef.current)
    dispatch(createPost(Object.fromEntries(content)))
  }
  
  return (
    <section className='p-4 bg-gray-200 h-screen'>
        <div className=''>
        <div>
            <form ref={postRef} onSubmit={handlePostSubmit} className='flex flex-col gap-4'>
                <input type='text'  className='p-2 text-xl w-full h-32 rounded-xl shadow-md' />
                <button disabled={posting} className='ml-auto px-4 py-2 bg-green-400 text-white text-lg font-semibold rounded-md'>Post</button>
            </form>
        </div>
        <div>
            {loading ? <Shimmer/>
            :
            <Posts posts={posts} />
          }
        </div>
        <Pagination path="/getpostsuggestions" numOfPages={numOfPages} page={page}/>
        </div>
    </section>
  )
}

export default Content