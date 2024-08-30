import React, { useState } from 'react'

import TableComponent from './Table'
import { useQuery} from '@tanstack/react-query'
import { HTTPRequest } from '@/api/api'
import Loading from '../Loading/Loading'
import ProjectNav from '../nav/ProjectNav'
import {
    Sheet,
    SheetContent,
    SheetDescription,
    SheetHeader,
    SheetTitle,
    SheetTrigger,
  } from "@/components/ui/sheet"
import { Menu, Settings } from 'lucide-react'
import Link from 'next/link'
import { Button } from '../ui/button'
  


interface props {
    id: string
}

const Project: React.FC<props> = ({ id }) => {
    const [open, setOpen] = useState<boolean>(false)

    const { data, error, isLoading } = useQuery({
        queryKey: ["flags",id],
        queryFn: async () => {
            return (await HTTPRequest(`/flags/pid/${id}`, {}, "GET"))
        },
        refetchOnMount:true
    })

    const { data: projectdata, error: projectdataerror, isLoading: projectdataloading } = useQuery({
        queryKey: [`project-${id}`,id],
        queryFn: async () => {
            return (await HTTPRequest(`/project/${id}`, {}, "GET"))
        },
        refetchOnMount:true
    })



    if (isLoading) return <Loading />

    return (
        <div className='w-full'>
            <div className='mb-4 flex items-center justify-between'>
                <h2 className='text-xl font-semibold'>{projectdata?.response.message.name}</h2>
                <div className='hidden md:flex items-center gap-3'>
                    <ProjectNav id={id} />
                    <Link href = {`/settings/${id}`}><Button variant={"secondary"} ><Settings /></Button></Link>
                </div>
                <div className='md:hidden flex  items-center gap-3'>
                    <Sheet open = {open} onOpenChange={setOpen}>
                        <SheetTrigger><Menu /></SheetTrigger>
                        <SheetContent>
                            <SheetHeader>
                                <SheetTitle>Project Options</SheetTitle>
                                <SheetDescription>
                                    <ProjectNav id={id} />
                                </SheetDescription>
                            </SheetHeader>
                        </SheetContent>
                    </Sheet>
                </div>
                

            </div>
            <TableComponent list={data?.response.message} />
        </div>
    )
}

export default Project