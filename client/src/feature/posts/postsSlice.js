import {createSlice, createAsyncThunk} from "@reduxjs/toolkit"
import { axiosBase } from "../../utils"
import {toast} from "react-toastify"

const initialState = {
    posts : [],
    
    loading: false,
    success: false,
    numOfPages: 0,
    page: 0,
    posting: false,
    likeCommentBtn: false
}



export const updatePost = createAsyncThunk("posts/updatePost",
    async ({postId, content}, thunkAPI) => {
        try {
            const {token} = thunkAPI.getState().user
            const resp = await axiosBase.put(`/updateposts/${postId}`, {content}, {
                headers : {
                    "Authorization" : `Bearer ${token}`
                }
            })

            // fetch current posts
            const {page} = thunkAPI.getState().posts
            thunkAPI.dispatch(getPosts(`/getpostsbyuser?page=${page}`))
            
            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const handleLikeAsync = createAsyncThunk("posts/handleLikeAsync",
    async ({postId, path}, thunkAPI) => {
        try {
            const {token} = thunkAPI.getState().user
            const resp = await axiosBase.get(`/addlike/${postId}`, {
                headers : {
                    "Authorization" : `Bearer ${token}`
                }
            })

            // fetch current posts
            const {page} = thunkAPI.getState().posts
            thunkAPI.dispatch(getPosts(`${path}?page=${page}`))
            
            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const createPost = createAsyncThunk("posts/createPost",
    async (data, thunkAPI) => {
        try {
            const {token} = thunkAPI.getState().user
            const resp = await axiosBase.post("/createposts",data, {
                headers : {
                    "Authorization" : `Bearer ${token}`
                }
            })

            return resp?.data 
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

// export const getUserPosts = createAsyncThunk("posts/getUserPosts", 
//     async (path, thunkAPI) => {
//         try {
//             const {token} = thunkAPI.getState().user
//             const resp = await axiosBase.get(`${path}`, {
//                 headers : {
//                     "Authorization" : `Bearer ${token}`
//                 }
//             })

//             return resp?.data
//         } catch (error) {
//             return thunkAPI.rejectWithValue(error?.response)
//         }
//     }
// )

export const getPosts = createAsyncThunk("posts/getPosts", 
    async (path, thunkAPI) => {
        try {
            const {token} = thunkAPI.getState().user
            const resp = await axiosBase.get(`${path}`, {
                headers : {
                    "Authorization" : `Bearer ${token}`
                }
            })

            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

const postsSlice = createSlice({
    name: "posts",
    initialState: JSON.parse(localStorage.getItem("posts")) || initialState,
    reducers: {
        
    },
    extraReducers: (builder) => {
        builder.addCase(getPosts.pending, (state, {payload}) => {
            state.loading = true
        }).addCase(getPosts.fulfilled, (state, {payload}) => {
            state.loading = false
            state.success = true
            state.posts = payload?.posts
            state.numOfPages = payload?.numOfPages
            state.page = payload?.page
            localStorage.setItem("posts", JSON.stringify(state))
        }).addCase(getPosts.rejected, (state, {payload}) => {
            state.loading = false
            state.success = false
            toast.error(payload?.data?.msg)
        }).addCase(createPost.pending, (state, {payload}) => {
            state.posting = true
        }).addCase(createPost.fulfilled, (state, {payload}) => {
            state.posting = false
            toast.success("posted successfully!")
        }).addCase(createPost.rejected, (state, {payload}) => {
            state.posting = false
            toast.error(payload?.data?.msg)
        }).addCase(handleLikeAsync.pending, (state, {payload}) => {
            state.likeCommentBtn = true
        }).addCase(handleLikeAsync.fulfilled, (state, {payload}) => {
            state.likeCommentBtn = false
        }).addCase(handleLikeAsync.rejected, (state, {payload}) => {
            state.likeCommentBtn = false
            toast.error(payload?.data?.msg)
        }).addCase(updatePost.pending, (state, {payload}) => {
            state.posting = true
        }).addCase(updatePost.fulfilled, (state, {payload}) => {
            state.posting = false
        }).addCase(updatePost.rejected, (state, {payload}) => {
            state.posting = false
            toast.error(payload?.data?.msg)
        })
    }
})

export default postsSlice.reducer