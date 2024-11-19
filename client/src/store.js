import {configureStore} from "@reduxjs/toolkit"
import userReducer from "./feature/user/userSlice"
import postsReducer from "./feature/posts/postsSlice"
import commentssReducer from "./feature/comments/commentsSlice"

export const store = configureStore({
    reducer : {
        user : userReducer,
        posts: postsReducer,
        comments: commentssReducer
    }
})