import {createSlice, createAsyncThunk} from "@reduxjs/toolkit"
import { axiosBase } from "../../utils"
import {toast} from "react-toastify"

const initialState = {
    loading: false,
    submitting: false,
    commentsAll: {}
}

export const deleteComment = createAsyncThunk("comments/deleteComment",
    async (data, thunkAPI) => {
        try {
            const {token} = thunkAPI.getState().user
            const resp = await axiosBase.delete(`/deletecomments/${data.id}`, {
                headers : {
                    "Authorization" : `Bearer ${token}`
                }
            })

            // refetch new comments
            thunkAPI.dispatch(getComments(data.postId))

            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const getComments = createAsyncThunk("comments/getComments",
    async (postId, thunkAPI) => {
        try {
            const resp= await axiosBase.get(`/getpostcomments/${postId}`)

            return {data: resp?.data, postId}
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const createComment = createAsyncThunk("comments/createComment",
    async (data, thunkAPI) => {
        try {
            const {token} = thunkAPI.getState().user
            const resp = await axiosBase.post(`/createcomments`, data, {
                headers : {
                    "Authorization" : `Bearer ${token}`
                }
            })

            // refetch new comments
            thunkAPI.dispatch(getComments(data.post_id))

            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

const commentsSlice = createSlice({
    name: "comments",
    initialState: JSON.parse(localStorage.getItem("comments")) || initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(getComments.pending, (state, {payload}) => {
            state.loading = true
        }).addCase(getComments.fulfilled, (state, {payload}) => {
            const {data:{comments}, postId} = payload

            state.loading = false
            state.commentsAll[postId]=comments
        }).addCase(getComments.rejected, (state, {payload}) => {
            state.loading = false
            toast.error(payload?.data?.msg)
        }).addCase(createComment.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(createComment.fulfilled, (state, {payload}) => {
            state.submitting = false
        }).addCase(createComment.rejected, (state, {payload}) => {
            state.submitting = false
            toast.error(payload?.data?.msg)
        }).addCase(deleteComment.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(deleteComment.fulfilled, (state, {payload}) => {
            state.submitting = false
        }).addCase(deleteComment.rejected, (state, {payload}) => {
            state.submitting = false
            toast.error(payload?.data?.msg)
        })
    }
})

export default commentsSlice.reducer