import React, { useState } from 'react'
import dayjs from "dayjs"
import {useDispatch, useSelector} from "react-redux"
import { AiTwotoneLike } from "react-icons/ai";
import { LiaCommentSolid } from "react-icons/lia";
import { BsThreeDotsVertical } from "react-icons/bs";
import { handleLikeAsync } from '../feature/posts/postsSlice';
import CommentBox from './CommentBox';

const Posts = ({posts}) => {
    const [openCommentBoxId, setOpenCommentBoxId] = useState("")
    const dispatch = useDispatch()
    const {likeCommentBtn} = useSelector(state => state.posts)

    const handleCommentBoxOpen = (postId) => {
        if (postId == openCommentBoxId){
            setOpenCommentBoxId("")
        } else {
            setOpenCommentBoxId(postId)
        }
    }

  return (
    <>
    {
        posts.map(post => {
            return (
                <article key={post.id} className='flex flex-col gap-2 bg-white my-4 shadow-md rounded-md p-2 md:my-6 md:p-4 hover:shadow-xl'>
                    <div className='flex justify-between items-center'>
                    <div className='flex gap-2 justify-start items-center'>
                        <h1 className='text-xl rounded-full bg-slate-300 capitalize px-4 py-2 font-semibold'>{post.name[0]}</h1>
                        <div className='flex flex-col  items-start'>
                            <h1 className='text-sm text-slate-700'>{post.name}</h1>
                            <h1 className='text-sm text-slate-600'>{dayjs(post.created_at).toString()}</h1>
                        </div>
                    </div>
                    <h1><BsThreeDotsVertical/></h1>
                    </div>
                    <div className='p-4 text-slate-600'>
                        <h1>{post.content}</h1>
                    </div>
                    <div className='flex px-2 justify-between items-center cursor-pointer'>
                        <div className='flex w-1/2 gap-1 bg-grey-300 border-r-2 border-gray-400 items-center justify-center'>
                        <h1 className='text-md' > {post.likes}</h1>
                        <button disabled={likeCommentBtn} onClick={() => dispatch(handleLikeAsync({path: "/getpostsuggestions", postId: post.id}))} className='text-md'> <AiTwotoneLike />
                        </button>
                        </div>
                        <div className='flex w-1/2 gap-1 bg-grey-300 items-center justify-center'>
                        <h1 className='text-md'> {post.comments}</h1>
                        <button onClick={() => handleCommentBoxOpen(post.id)} className='text-md'> <LiaCommentSolid />
                        </button>
                        </div>
                    </div>
                    {openCommentBoxId == post.id && <CommentBox postId={post.id} />}
                </article>
            )
        })
    }
    </>
  )
}

export default Posts