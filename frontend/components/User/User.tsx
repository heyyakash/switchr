import React, { useState } from 'react'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { HTTPRequest } from '../../api/api'
import Loading from '../Loading/Loading'
import { useForm } from 'react-hook-form'
import { Label } from '@radix-ui/react-dropdown-menu'
import { Input } from '../ui/input'
import { Button } from '../ui/button'
import { Check, RotateCw } from 'lucide-react'
import { toast } from 'sonner'
import { useRouter } from 'next/router'
import { Badge } from '../ui/badge'
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
  } from "@/components/ui/dialog"

interface userSettings {
    fullname: string
}

interface passwordUpdate {
    current: string
    new: string
}

const User = () => {
    const [open, setOpen] = useState(false)
    const [emailSent, setEmailSent] = useState(false)
    const [udpating, setUpdating] = useState(false)
    const [change, setChange] = useState(false)
    const [deleting, setDeleting] = useState(false)
    const queryClient = useQueryClient()
    const router = useRouter()
    const { handleSubmit, register } = useForm<userSettings>()
    const { handleSubmit: handlePasswordSubmit, register: passwordRegister, reset: changePasswordReset } = useForm<passwordUpdate>()
    const { data, isLoading } = useQuery({ queryKey: ["user"], queryFn: async () => { return (await HTTPRequest("/user", {}, "GET")) } })


    const handleUserNameChange = async (userData: userSettings) => {
        setUpdating(true)
        const payload = {
            uid: data?.response.message?.uid,
            fullname: userData.fullname
        }
        const res = await HTTPRequest("/user", { body: JSON.stringify(payload) }, "PATCH")
        if (res?.response.success) {
            queryClient.invalidateQueries({
                queryKey: [`user`]
            })
            toast.success(res?.response.message)

        } else {
            toast.error(res?.response.message)
        }
        setUpdating(false)
    }

    const DeleteUser = async () => {
        setDeleting(true)
        const res = await HTTPRequest("/user", {}, "DELETE")
        if (res?.response.success) {
            queryClient.invalidateQueries({
                queryKey: [`user`]
            })
            router.push('/login')
            toast.success(res?.response.message)

        } else {
            toast.error(res?.response.message)
        }
        setDeleting(false)
    }

    const handlePasswordUpdateChange = async (data: passwordUpdate) => {
        setChange(true)
        const res = await HTTPRequest("/user/password", { body: JSON.stringify(data) }, "PATCH")
        if (res?.response.success) {
            queryClient.invalidateQueries({
                queryKey: [`user`]
            })
            toast.success(res?.response.message)

        } else {
            toast.error(res?.response.message)
        }
        changePasswordReset()
        setChange(false)
    }

    const VerifyUser = async () => {
        setEmailSent(false)
        setOpen(true)
        const res = await HTTPRequest("/user/verify", {}, "POST")
        if (res?.response.success) {
            toast.success(res.response.message)
            setEmailSent(true)
        } else {
            toast.error(res?.response.message)
        }
    }

    if (isLoading) return <Loading />
    return (
        <div>
            <h2 className='text-2xl '>{data?.response.message?.fullname || "User"}&apos;s Settings</h2>
            <Dialog open={open} onOpenChange={setOpen}>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Verifying User</DialogTitle>
                        <DialogDescription className='flex items-center '>
                            <div>
                                {!emailSent ? (<RotateCw size={20} className='animate-spin' />) : (<Check />)}
                            </div>
                            <div className='w-full'>{emailSent ? "Verification email sent, kindly check your email" : "Sending verification email"}</div>
                        </DialogDescription>
                    </DialogHeader>
                </DialogContent>
            </Dialog>
            {data?.response?.message?.verified ? (<></>) : (<>
                <div className='mt-12'>
                <h3 className='text-xl'>Verify Your Account</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>Before accessing any of the features of the app, you need to verify your email first</p>
                <Button onClick={() => VerifyUser()} className='cursor-pointer' variant={"default"}>Verify</Button>
                
            </div>
                
            </>)}
            <div className='mt-6'>
                <h3 className='text-xl'>Edit User Details</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>You can change your current project name to another unique project name</p>
                <form onSubmit={handleSubmit(handleUserNameChange)} className='flex flex-col gap-3 items-start'>
                    <Label>Full Name</Label>
                    <Input type="text" {...register("fullname")} placeholder={data?.response.message?.fullname} className='max-w-[500px] w-full' />
                    <Button disabled={udpating}>
                        {udpating ? <RotateCw className='animate-spin' /> : "Update"}
                    </Button>
                </form>
            </div>
            <div className='mt-6'>
                <h3 className='text-xl'>Change Password</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>You can change your current project name to another unique project name</p>
                <form onSubmit={handlePasswordSubmit(handlePasswordUpdateChange)} className='flex flex-col gap-3 items-start'>
                    <Label>Current Password</Label>
                    <Input type="password" {...passwordRegister("current")} placeholder="Current Password" className='max-w-[500px] w-full' />
                    <Label>New Password</Label>
                    <Input type="password" {...passwordRegister("new")} placeholder="New Password" className='max-w-[500px] w-full' />
                    <Button disabled={change}>
                        {change ? <RotateCw className='animate-spin' /> : "Update"}
                    </Button>
                </form>
            </div>
            <div className='mt-12'>
                <h3 className='text-xl'>Delete Account</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>Be carefull, deleting the account is a permanent action</p>
                <form onSubmit={handleSubmit(handleUserNameChange)} className='flex flex-col gap-3 items-start'>
                    <Button onClick={() => DeleteUser()} variant={"destructive"} disabled={deleting}>
                        {deleting ? <RotateCw className='animate-spin' /> : "Delete Account"}
                    </Button>
                </form>
            </div>

        </div>
    )
}

export default User