import {createSlice, createAsyncThunk} from "@reduxjs/toolkit"
import { axiosBase } from "../../utils"
import {toast} from 'react-toastify'
import { getPosts } from "../posts/postsSlice"

const initialState = {
    submitting: false,
    errorMsg: "",
    token: "",
    user: {},
    success: false,
    loading: false,
    followSuggestions: [],
    searchName: "",
    searchUsers: [],
    profile: {
        user : {
            name : "Guest User"
        }
    },
    profileLoading: false
}

export const updateUser = createAsyncThunk("user/updateUser", 
    async(_, thunkAPI) => {
        try {
            const {id, name, location, age, username, bio} = thunkAPI.getState().user.profile?.user
            const {token} = thunkAPI.getState().user

            const resp = await axiosBase.put("/updateusers", {name, location, age, username, bio}, {
                headers : {
                    "Authorization" : `Bearer ${token}`
                }
            })

            // refetch new profile
            thunkAPI.dispatch(getProfile(id))

            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const getProfile = createAsyncThunk("user/getProfile", 
    async (userId, thunkAPI) => {
        try {
            const resp = await axiosBase.get(`/getusersdetails/${userId}`)

            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const getUsers = createAsyncThunk("user/getUsers",
    async ({name, location}, thunkAPI) => {
        try {

            const resp = await axiosBase.get(`/getusers?name=${name || ""}&location=${location || ""}`)

            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const followPerson = createAsyncThunk("user/followPerson",
    async (person, thunkAPI) => {
        try {
            const {token} = thunkAPI.getState().user
            const resp = await axiosBase.post("/followpersons",{person}, {
                headers : {
                    "Authorization" : `Bearer ${token}`
                }
            })

            // refetch follow suggestions
            thunkAPI.dispatch(getFollowSuggestions())
            thunkAPI.dispatch(getPosts("/getpostsuggestions"))

            return resp?.data
        } catch (error) {
            return thunkAPI.rejectWithValue(error?.response)
        }
    }
)

export const getFollowSuggestions = createAsyncThunk("user/getFollowSuggestions",
    async (_, thunkAPI) => {
        try {
            const {token} = thunkAPI.getState().user
            const resp = await axiosBase.get("/followsuggestions", {
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
        },
        setSearchName : (state, {payload}) => {
            state.searchName = payload
        },
        changeUserDetails: (state, {payload}) => {
            const {name, value} = payload

            state.profile.user[name] = value
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
        }).addCase(getFollowSuggestions.pending, (state, {payload}) => {
            state.loading = true
        }).addCase(getFollowSuggestions.fulfilled, (state, {payload}) => {
            state.loading = false
            state.followSuggestions = payload?.follow_suggestions
        }).addCase(getFollowSuggestions.rejected, (state, {payload}) => {
            state.loading = false
            toast.error(payload?.data?.msg)
        }).addCase(followPerson.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(followPerson.fulfilled, (state, {payload}) => {
            state.submitting = false
        }).addCase(followPerson.rejected, (state, {payload}) => {
            state.submitting = false
            toast.error(payload?.data?.msg)
        }).addCase(getUsers.pending, (state, {payload}) => {
            state.loading = true
        }).addCase(getUsers.fulfilled, (state, {payload}) => {
            state.loading = false
            // state.followSuggestions = payload?.follow_suggestions
            state.searchUsers = payload?.user
        }).addCase(getUsers.rejected, (state, {payload}) => {
            state.loading = false
            toast.error(payload?.data?.msg)
        }).addCase(getProfile.pending, (state, {payload}) => {
            state.profileLoading = true
        }).addCase(getProfile.fulfilled, (state, {payload}) => {
            state.profileLoading = false
            // state.followSuggestions = payload?.follow_suggestions
            // console.log(payload)
            state.profile = payload
        }).addCase(getProfile.rejected, (state, {payload}) => {
            state.profileLoading = false
            toast.error(payload?.data?.msg)
        }).addCase(updateUser.pending, (state, {payload}) => {
            state.submitting = true
        }).addCase(updateUser.fulfilled, (state, {payload}) => {
            state.submitting = false
        }).addCase(updateUser.rejected, (state, {payload}) => {
            state.submitting = false
            toast.error(payload?.data?.msg)
        })
    }
})

export const {setUser, logout, setToDefault,setSearchName, changeUserDetails} = userSlice.actions

export default userSlice.reducer