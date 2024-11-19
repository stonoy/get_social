import {createSlice, createAsyncThunk} from "@reduxjs/toolkit"
import { axiosBase } from "../../utils"
import {toast} from 'react-toastify'

const initialState = {
    submitting: false,
    errorMsg: "",
    token: "",
    user: {},
    success: false
}

export const login = createAsyncThunk("user/login",
    async(data, thunkAPI) => {
        try {
            const resp = await axiosBase.post("/login", data)
            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const register = createAsyncThunk("user/register",
    async(data, thunkAPI) => {
        try {
            const resp = await axiosBase.post("/register", data)
            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

const userSlice = createSlice({
    name: "user",
    initialState: JSON.parse(localStorage.getItem("user")) || initialState,
    reducers: {
        setUser : (state, {payload}) => {
            console.log(payload)
        },
        logout: () => {
            return initialState
        },
        setToDefault : (state, {payload}) => {
            state.submitting = false
            state.success = false
            state.errorMsg = ""
            localStorage.setItem("user", JSON.stringify(state))
        }
    },
    extraReducers: (builder) => {
        builder.addCase(login.pending, (state, action) => {
            state.submitting = true
        }).addCase(login.fulfilled, (state, {payload}) => {
            const {token, user} = payload
            state.submitting = false
            state.success = true
            state.token = token
            state.user = user
            toast.success("loggin in!")
            localStorage.setItem("user", JSON.stringify(state))
        }).addCase(login.rejected, (state, {payload}) => {
            state.submitting = false
            state.success = false
            toast.error(payload?.data?.msg)
            state.errorMsg = payload?.data?.msg
        }).addCase(register.pending, (state, action) => {
            state.submitting = true
        }).addCase(register.fulfilled, (state, {payload}) => {
            const {token, user} = payload
            state.submitting = false
            state.success = true
            state.token = token
            state.user = user
            toast.success("registered!")
            localStorage.setItem("user", JSON.stringify(state))
        }).addCase(register.rejected, (state, {payload}) => {
            state.submitting = false
            state.success = false
            toast.error(payload?.data?.msg)
            state.errorMsg = payload?.data?.msg
        })
    }
})

export const {setUser, logout, setToDefault} = userSlice.actions

export default userSlice.reducer