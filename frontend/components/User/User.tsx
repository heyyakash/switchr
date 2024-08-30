import React, { useState } from 'react'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { HTTPRequest } from '../../api/api'
import Loading from '../Loading/Loading'
import { useForm } from 'react-hook-form'
import { Label } from '@radix-ui/react-dropdown-menu'
import { Input } from '../ui/input'
import { Button } from '../ui/button'
import { RotateCw } from 'lucide-react'
import { toast } from 'sonner'

interface userSettings{
    fullname: string
}

interface passwordUpdate{
    current: string
    new : string
}

const User = () => {
    const [udpating, setUpdating] = useState(false)
    const [change, setChange] = useState(false)

    const queryClient = useQueryClient()
    const {handleSubmit, register} = useForm<userSettings>()
    const {handleSubmit:handlePasswordSubmit, register: passwordRegister, reset: changePasswordReset} = useForm<passwordUpdate>()
    const {data, isLoading} = useQuery({ queryKey: ["user"], queryFn: async () => { return (await HTTPRequest("/user", {}, "GET")) } })


    const handleUserNameChange = async (userData: userSettings) => {   
        setUpdating(true)
        const payload = {
            uid: data?.response.message?.uid,
            fullname : userData.fullname
        }
        const res = await HTTPRequest("/user", {body:JSON.stringify(payload)}, "PATCH")
        if(res?.response.success){
            queryClient.invalidateQueries({
                queryKey:[`user`]
            })
            toast.success(res?.response.message)

        }else{
            toast.error(res?.response.message)
        }
        setUpdating(false)
    }

    const handlePasswordUpdateChange = async (data: passwordUpdate) => {
        setChange(true)
        const res = await HTTPRequest("/user/password", {body:JSON.stringify(data)}, "PATCH")
        if(res?.response.success){
            queryClient.invalidateQueries({
                queryKey:[`user`]
            })
            toast.success(res?.response.message)

        }else{
            toast.error(res?.response.message)
        }
        changePasswordReset()
        setChange(false)
    }

    if (isLoading) return <Loading />
    return (
        <div>
            <h2 className='text-2xl '>{data?.response.message?.fullname || "User"}&apos;s Settings</h2>
            <div className='mt-12'>
                <h3 className='text-xl'>Edit User Details</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>You can change your current project name to another unique project name</p>
                <form onSubmit={handleSubmit(handleUserNameChange)} className='flex flex-col gap-3 items-start'>
                    <Label>Full Name</Label>
                    <Input type = "text" {...register("fullname")} placeholder= {data?.response.message?.fullname} className='max-w-[500px] w-full' />
                    <Button disabled = {udpating}>
                    {udpating? <RotateCw className='animate-spin' />: "Update"}
                </Button>
                </form>
            </div>
            <div className='mt-6'>
                <h3 className='text-xl'>Change Password</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>You can change your current project name to another unique project name</p>
                <form onSubmit={handlePasswordSubmit(handlePasswordUpdateChange)} className='flex flex-col gap-3 items-start'>
                    <Label>Current Password</Label>
                    <Input type = "password" {...passwordRegister("current")} placeholder= "Current Password" className='max-w-[500px] w-full' />
                    <Label>New Password</Label>
                    <Input type = "password" {...passwordRegister("new")} placeholder= "New Password" className='max-w-[500px] w-full' />
                    <Button disabled = {change}>
                    {change? <RotateCw className='animate-spin' />: "Update"}
                </Button>
                </form>
            </div>
        </div>
    )
}

export default User