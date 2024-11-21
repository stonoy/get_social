import React, { useRef, useState } from 'react'
import dayjs from "dayjs"
import {useDispatch, useSelector} from "react-redux"
import { AiTwotoneLike } from "react-icons/ai";
import { LiaCommentSolid } from "react-icons/lia";
import { BsThreeDotsVertical } from "react-icons/bs";
import { handleLikeAsync, updatePost } from '../feature/posts/postsSlice';
import CommentBox from './CommentBox';
import UserNameFirst from './UserNameFirst';

const Posts = ({posts, isTimeLine}) => {
    const [openCommentBoxId, setOpenCommentBoxId] = useState("")
    const [openBoxId, setOpenBoxId] = useState("")
    const [openUpdateId, setOpenUpdateId] = useState("")
    const dispatch = useDispatch()
    const inputRef = useRef(null)
    const {likeCommentBtn, posting} = useSelector(state => state.posts)

    const handlePostUpdate = (e, postId) => {
        e.preventDefault()
        console.log("hi")

        const formData = new FormData(inputRef.current)
        const {content} = Object.fromEntries(formData)

        dispatch(updatePost({postId, content}))
    }

    const handleCommentBoxOpen = (postId) => {
        if (postId == openCommentBoxId){
            setOpenCommentBoxId("")
        } else {
            setOpenCommentBoxId(postId)
        }
    }

    const handleUpdateOpen = (postId) => {
        if (postId == openUpdateId){
            setOpenUpdateId("")
            setOpenBoxId("")
        } else {
            setOpenUpdateId(postId)
            setOpenBoxId("")
        }
    }

    const handleBoxOpen = (postId) => {
        if (postId == openBoxId){
            setOpenBoxId("")
        } else {
            setOpenBoxId(postId)
        }
    }

  return (
    <>
    {
        posts.map(post => {
            return (
                <article key={post.id} className='flex flex-col relative gap-2 bg-white my-4 shadow-md rounded-md p-2 md:my-6 md:p-4 hover:shadow-xl'>
                    <div className='flex justify-between items-center'>
                    <div className='flex gap-2 justify-start items-center'>
                        <UserNameFirst letter={post.name[0]} id={post.author} />
                        <div className='flex flex-col  items-start'>
                            <h1 className='text-sm text-slate-700'>{post.name}</h1>
                            <h1 className='text-sm text-slate-600'>{dayjs(post.created_at).toString()}</h1>
                        </div>
                    </div>
                    {isTimeLine && <h1 onClick={() => handleBoxOpen(post.id)}><BsThreeDotsVertical/></h1>}
                    </div>
                    <div className='p-4 text-slate-600'>
                        { openUpdateId == post.id ? <form ref={inputRef} onSubmit={(e) => handlePostUpdate(e, post.id)} className='flex flex-col w-full justify-center items-start'>
                            <input type='text' name="content" className='w-full h-24'/>
                            <div className='flex justify-between items-center w-full'>
                            <button disabled={posting} className='my-2 py-1 px-2 rounded-md bg-slate-300 md:my-4'>Update</button>
                            <button type='button' onClick={() => handleUpdateOpen("")} className='my-2 py-1 px-2 rounded-md bg-red-200 md:my-4'>Cancel</button>
                            </div>
                        </form>
                        :
                        <h1>{post.content}</h1>    
                    }
                    </div>
                    <div className='flex px-2 justify-between items-center cursor-pointer'>
                        <div className='flex w-1/2 gap-1 bg-grey-300 border-r-2 border-gray-400 items-center justify-center'>
                        <h1 className='text-md' > {post.likes}</h1>
                        <button disabled={likeCommentBtn} onClick={() => dispatch(handleLikeAsync({path: "/getpostsbyuser", postId: post.id}))} className='text-md'> <AiTwotoneLike />
                        </button>
                        </div>
                        <div className='flex w-1/2 gap-1 bg-grey-300 items-center justify-center'>
                        <h1 className='text-md'> {post.comments}</h1>
                        <button onClick={() => handleCommentBoxOpen(post.id)} className='text-md'> <LiaCommentSolid />
                        </button>
                        </div>
                    </div>
                    {openCommentBoxId == post.id && <CommentBox postId={post.id} />}
                    {openBoxId == post.id && <ul className='w-2/12 absolute left-3/4 top-1/4 bg-slate-700 text-white rounded-md'>
                            <li onClick={() => handleUpdateOpen(post.id)} className='py-0.5 px-1'>Update</li>
                            <li className='py-0.5 px-1'>Delete</li>
                        </ul>}
                </article>
            )
        })
    }
    </>
  )
}

export default Posts