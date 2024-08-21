import { Key, Menu, Search, StickyNote } from 'lucide-react'
import React, { ReactNode } from 'react'
import { Input } from '../ui/input'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import { Popover, PopoverContent, PopoverTrigger } from '../ui/popover'
import { Button } from '../ui/button'
import { useRouter } from 'next/router'
import Link from 'next/link'
import { ModeToggle } from '../Toggle'
import {
  Command,
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
  CommandSeparator,
  CommandShortcut,
} from "@/components/ui/command"
import { SearchBox } from '../Command/SearchBox'
import { Badge } from '../ui/badge'
import { HTTPRequest } from '@/api/api'
import { useQuery, useQueryClient } from '@tanstack/react-query'

interface props {
    children :ReactNode
}

const Layout: React.FC<props> = (props) => {
    const getUser = async () => {
        return (await HTTPRequest("/user",{}, "GET"))
    }
    const {data, error, isLoading} = useQuery({queryKey:["user"] ,queryFn:getUser  })
    const client = useQueryClient()
    const { data:userprojectmap } = useQuery({
        queryKey: ["projects"],
        queryFn: async () => {
            return (await HTTPRequest("/userprojectmap", {}, "GET"))
        }
    })
    const router = useRouter()
    if(data && !data.response.success || error){
        router.push('/login')
    }
    const logout = () => {
        document.cookie = `token= ;secure=true; path=/`
        router.push('/login')
    }
 

    return (
        <div className='w-full h-screen flex flex-col'>
            <div className='h-[70px] border-b-[1.4px] border-secondary grid grid-rows-1 grid-cols-2 md:grid-cols-3 items-center p-4 px-6'>

                <div className='flex items-center gap-4 w-full'>
                    <Link href = "/dashboard" className="flex gap-2 items-center text-primary text-[1.2rem] font-medium ">
                        <Key size={"20px"} /> 
                    </Link>
                    / <div className='text-md text-primary/50'>{data?.response?.message?.fullname}</div>
                    {data?.response?.message?.verified ? (<></>):(<Badge variant={"destructive"}>Unverified</Badge>)}
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
                                <Avatar>
                                    <AvatarImage src="https://github.com/shadcn.png" />
                                    <AvatarFallback>CN</AvatarFallback>
                                </Avatar>
                            </PopoverTrigger>
                            <PopoverContent className='mr-2 border-secondary flex flex-col gap-3'>
                                <div className='w-full grid grid-cols-2 rounded-lg overflow-hidden grid-rows-1'>
                                    <img src="https://github.com/shadcn.png" alt="" />
                                    <div className='bg-primary flex flex-col justify-center dark:text-black items-center'>
                                        <h2 className='text-4xl '>{userprojectmap?.response?.message?.length}</h2>
                                        <p>Projects</p>
                                    </div>
                                </div>
                                <Button variant={"secondary"} size = {"lg"} className='w-full'>Settings</Button>
                                <Button onClick={()=>logout()} variant={"destructive"} size = {"lg"} className='w-full bg-red-500'>Logout</Button>
                            </PopoverContent>
                        </Popover>
                    </div>
                </div>

            </div>
            <div className='w-full h-full flex'>
                {props.children}
            </div>
        </div>
    )
}

export default Layout