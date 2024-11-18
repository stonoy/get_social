import { useEffect, useRef } from "react"
import { TbLoader3 } from "react-icons/tb"
import { useDispatch, useSelector } from "react-redux"
import {Link, useNavigate} from "react-router-dom"
import { register, setToDefault } from "../feature/user/userSlice"

const Register = () => {
    const formRef = useRef(null)
    const dispatch = useDispatch()
    const {submitting, success, errorMsg} = useSelector((state) => state.user)
    const navigate = useNavigate()

    // if (errorMsg) {
    //     toast.error(errorMsg)
    // }

    const handleSubmit = (e) => {
        e.preventDefault()
        const formData = new FormData(formRef.current)
        dispatch(register(Object.fromEntries(formData)))
    }

    useEffect(() => {
        if (success) {
            navigate("/")
        }

        return () => {
            dispatch(setToDefault())
        }
    },[success])

    return (
        <section className="flex justify-center items-center h-screen">
            <div className="p-4 shadow-2xl text-grey-700">
                <h1 className="text-xl font-semibold">Login</h1>
                <form ref={formRef} onSubmit={handleSubmit} className="flex flex-col gap-2 my-4">
                    <div className="my-2">
                        <label>Name</label>
                        <input type="text" name="name" className="w-full p-1 my-2 outline outline-1 rounded-sm"/>
                    </div>
                    <div className="my-2">
                        <label>Email</label>
                        <input type="text" name="email" className="w-full p-1 my-2 outline outline-1 rounded-sm"/>
                    </div>
                    <div className="my-2">
                        <label>Password</label>
                        <input type="text" name="password" className="w-full p-1 my-2 outline outline-1 rounded-sm"/>
                    </div>
                    <button className="w-full my-4 rounded-sm text-lg bg-green-400 text-white p-1">{submitting ? <TbLoader3 className="animate-spin text-xl inline-block mx-auto"/> : "Register"}</button>
                </form>

                <p className="text-slate-700">Already a member? <Link className="underline" to="/login">Login</Link></p>
            </div>
        </section>
    )
}

export default Register