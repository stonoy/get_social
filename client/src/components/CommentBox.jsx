import React, { useEffect, useRef } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { createComment, getComments } from '../feature/comments/commentsSlice'

const CommentBox = ({postId}) => {
    const {commentsAll, loading, submitting} = useSelector(state => state.comments)
    const dispatch = useDispatch()
    const commentRef = useRef(null)

    useEffect(() => {
        dispatch(getComments(postId))
    }, [])

    const handleComment = (e) => {
        e.preventDefault()
        
        const formData = new FormData(commentRef.current)
        const {comment} = Object.fromEntries(formData)

        if (!comment){return}

        dispatch(createComment({comment, post_id:postId}))
    }

  return (
    <div className='w-full p-2  md:p-4'>
        <div className='border-2 max-h-32 overflow-auto'>
            {
                !loading ?
                
                    commentsAll[postId]?.length > 0 ?
                    <>
                       {commentsAll[postId].map(comment => {
                        return (
                            <div className='flex my-2 gap-2 justify-start items-center md:my-4'>
                            <h1 className='text-xl rounded-full bg-slate-300 capitalize px-4 py-2 font-semibold'>{comment.name[0]}</h1>
                            <div className='flex flex-col  items-start'>
                                <h1 className='text-md font-semibold text-slate-700'>{comment.name}</h1>
                                <h1 className='text-sm text-slate-600'>{comment.comment}</h1>
                            </div>
                        </div>
                        )
                       })}
                    </>
                    :
                    <h1 className='text-sm text-slate-600'>No comments to show</h1>
                
                :
                <div className='flex gap-2 justify-start items-center animate-pulse'>
                            <h1 className='text-xl h-2 rounded-full bg-slate-300 capitalize px-4 py-2 font-semibold'></h1>
                            <div className='flex flex-col  items-start h-2'>
                                <h1 className='text-sm text-slate-700'></h1>
                                <h1 className='text-sm text-slate-600'></h1>
                            </div>
                        </div>
            }
        </div>
        <form ref={commentRef} onSubmit={handleComment} className='flex w-full'>
        <input type='text' name='comment' className='w-9/12 p-2 border-2 '/>
        <button className='w-3/12 py-0.5 px-1'>comment</button>
        </form>
    </div>
  )
}

export default CommentBox