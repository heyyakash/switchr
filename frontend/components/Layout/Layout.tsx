import { Check, Key, Menu, RotateCw, Search, StickyNote } from 'lucide-react'
import React, { ReactNode, useState } from 'react'
import { Input } from '../ui/input'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import { Popover, PopoverContent, PopoverTrigger } from '../ui/popover'
import { Button } from '../ui/button'
import { useRouter } from 'next/router'
import Link from 'next/link'
import { ModeToggle } from '../Toggle'
import { SearchBox } from '../Command/SearchBox'
import { Badge } from '../ui/badge'
import { HTTPRequest } from '@/api/api'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
  } from "@/components/ui/dialog"
import Loading from '../Loading/Loading'

interface props {
    children: ReactNode
}

const Layout: React.FC<props> = (props) => {
    const getUser = async () => {
        return (await HTTPRequest("/user", {}, "GET"))
    }
    const [open, setOpen] = useState(false)
    const [emailSent, setEmailSent] = useState(false)
    const { data, error, isLoading } = useQuery({ queryKey: ["user"], queryFn: getUser })
    const client = useQueryClient()
    const { data: userprojectmap } = useQuery({
        queryKey: ["projects"],
        queryFn: async () => {
            return (await HTTPRequest("/userprojectmap", {}, "GET"))
        }
    })
    const router = useRouter()
    if (data && !data.response.success || error) {
        router.push('/login')
    }
    const logout = async() => {
        await HTTPRequest("/user/logout", {}, "POST")
    
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

    if(isLoading) return <Loading />
    return (
        <>
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

            <div className='w-full h-screen flex flex-col'>
                <div className='h-[70px] border-b-[1.4px] border-secondary grid grid-rows-1 grid-cols-2 md:grid-cols-3 items-center p-4 px-6'>

                    <div className='flex items-center gap-4 w-full'>
                        <Link href="/dashboard" className="flex gap-2 items-center text-primary text-[1.2rem] font-medium ">
                            <Key size={"20px"} />
                        </Link>
                        / <Link href="/dashboard" className='text-md text-primary/50'>{data?.response?.message?.fullname}</Link>
                        {data?.response?.message?.verified ? (<></>) : (<Badge onClick={() => VerifyUser()} className='cursor-pointer' variant={"destructive"}>Unverified</Badge>)}
                    </div>

                    <div className='hidden md:flex items-center gap-3 px-2 rounded-lg justify-center'>
                        Press <Badge variant="default" className=' font-semibold'>CTRL + K</Badge>

                        <SearchBox />
                        {/* <Input placeholder='Search your notes' className='border border-primary' type="text" /> */}
                    </div>


                    <div className='justify-self-end'>
                        <div className='flex gap-2  w-[90px] items-center '>
                            <ModeToggle />
                            <Popover>
                                <PopoverTrigger>
                                    <div className='flex items-center justify-center h-10 w-10 rounded-full bg-primary text-secondary'>{data?.response?.message?.fullname?.split(" ").map((x:any)=>x[0])}</div>
                                </PopoverTrigger>
                                <PopoverContent className='mr-2 border-secondary flex flex-col gap-3'>
                                    <div className='w-full grid grid-cols-2 rounded-lg overflow-hidden grid-rows-1'>
                                        <div className='flex flex-col justify-center h-[150px] bg-secondary text-primary items-center'>
                                            <h2 className='text-4xl '>{data?.response?.message?.fullname?.split(" ").map((x:any)=>x[0])}</h2>
                                        </div>
                                        <div className='bg-primary flex flex-col justify-center dark:text-black items-center'>
                                            <h2 className='text-4xl '>{userprojectmap?.response?.message?.length}</h2>
                                            <p>Projects</p>
                                        </div>
                                    </div>
                                    <Button variant={"secondary"} size={"lg"} className='w-full'>Settings</Button>
                                    <Button onClick={() => logout()} variant={"destructive"} size={"lg"} className='w-full bg-red-500'>Logout</Button>
                                </PopoverContent>
                            </Popover>
                        </div>
                    </div>

                </div>
                <div className='w-full h-full flex'>
                    {props.children}
                </div>
            </div>
        </>
    )
}

export default Layout