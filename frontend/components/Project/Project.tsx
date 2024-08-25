import React from 'react'

import TableComponent, { FlagRowInterface } from './Table'



import { useQuery, useQueryClient } from '@tanstack/react-query'
import { HTTPRequest } from '@/api/api'
import { z } from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { toast } from 'sonner'

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

    const { data, error, isLoading } = useQuery({
        queryKey: ["flags"],
        queryFn: async () => {
            return (await HTTPRequest(`/flags/pid/${id}`, {}, "GET"))
        }
    })

    const { data: projectdata, error: projectdataerror, isLoading: projectdataloading } = useQuery({
        queryKey: [`project-${id}`],
        queryFn: async () => {
            return (await HTTPRequest(`/project/${id}`, {}, "GET"))
        },
    })



    if (isLoading) return <Loading />

    return (
        <div className='w-full'>
            <div className='mb-4 flex items-center justify-between'>
                <h2 className='text-xl font-semibold'>{projectdata?.response.message.name}</h2>
                <div className='hidden md:flex items-center gap-3'>
                    <ProjectNav id={id} />
                </div>
                <div className='md:hidden flex flex-col items-center gap-3'>
                    <Sheet>
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
                <Link href = {`/settings/${id}`}><Button variant={"secondary"} ><Settings /></Button></Link>

            </div>
            <TableComponent list={data?.response.message} />
        </div>
    )
}

export default Project