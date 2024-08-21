import React from 'react'
import { Button } from '../ui/button'
import TableComponent, { FlagRowInterface } from './Table'
import {
    Sheet,
    SheetContent,
    SheetDescription,
    SheetHeader,
    SheetTitle,
    SheetTrigger,
} from "@/components/ui/sheet"

import { Input } from '../ui/input'
import { Label } from '../ui/label'
import { Flag } from 'lucide-react'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { HTTPRequest } from '@/api/api'
import { z } from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { toast } from 'sonner'
import { useRouter } from 'next/router'

interface props {
    id :string
}

const formSchema = z.object({
    flag: z.string().min(3, {
        message: "Length of name should be greater than 3"
    }),
    value: z.string().min(3, {
        message: "Length of name should be greater than 3"
    }),
})

const Project: React.FC<props> = ({id}) => {
    const queryClient = useQueryClient()
    const router = useRouter()
    const Flags: FlagRowInterface[] = [
        {
            fid: "sadasdf-sada21-asd213-232134",
            name: "SHOW_FEATURE",
            value: "TRUE",
            createdAt: "25th JUL 2025",
            createdBy: "Akash Sharma",
        },
        {
            fid: "sadasdf-sada21-asd213-232134",
            name: "SHOW_FEATURE",
            value: "TRUE",
            createdAt: "25th JUL 2025",
            createdBy: "Akash Sharma",
        },
        {
            fid: "sadasdf-sada21-asd213-232134",
            name: "SHOW_FEATURE",
            value: "TRUE",
            createdAt: "25th JUL 2025",
            createdBy: "Akash Sharma",
        },
        {
            fid: "sadasdf-sada21-asd213-232134",
            name: "SHOW_FEATURE",
            value: "TRUE",
            createdAt: "25th JUL 2025",
            createdBy: "Akash Sharma",
        },
    ]

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            flag: "",
            value:""
        },
    })

    const {data, error, isLoading} = useQuery({
        queryKey:["flags"],
        queryFn: async () => {
            return (await HTTPRequest(`/flags/pid/${id}`, {}, "GET"))
        }
    })

    const {data:projectdata, error :projectdataerror, isLoading: projectdataloading} = useQuery({
        queryKey:[`project-${id}`],
        queryFn:async() => {
            return (await HTTPRequest(`/project/${id}`, {}, "GET"))
        },
    })

    console.log(data)

    const createFlag = async (payload : z.infer<typeof formSchema>) => {
        const obj = {
            pid: id,
            flag: payload.flag,
            value : payload.value
        }
        const res = await HTTPRequest("/flags/create", {body:JSON.stringify(obj)}, "POST")
        if (res.response.success){
            queryClient.invalidateQueries({queryKey:["flags"]})
            toast.success("Flag created successfully")
        }else{
            toast.error(res.response.message)
        }

    }

    return (
        <div className='w-full'>
            <div className='mb-4 flex items-center justify-between'>
                <h2 className='text-xl font-semibold'>{projectdata?.response.message.name}</h2>
                <Sheet>
                    <SheetTrigger><Button variant={"default"}>Create Flag</Button></SheetTrigger>
                    <SheetContent>
                        <SheetHeader className='flex flex-col gap-5'>
                            <SheetTitle className='flex gap-2 text-xl items-center'> <Flag />Create a new Flag</SheetTitle>
                            <SheetDescription className='mt-5'>
                            <form onSubmit={form.handleSubmit(createFlag)} className='flex flex-col gap-4'>
                                <div className="grid w-full max-w-sm items-center gap-1.5">
                                    <Label htmlFor="Key" className='mb-1'>Key</Label>
                                    <Input type="Key" id="Key" {...form.register("flag")} placeholder="Key" />
                                </div>
                                <div className="grid w-full max-w-sm items-center gap-1.5">
                                    <Label htmlFor="Value" className='mb-1'>Value</Label>
                                    <Input type="Value" id="Value" {...form.register("value")} placeholder="Value" />
                                </div>
                                <Button variant={"default"} className='w-full mt-2'>Create</Button>

                            </form>
                            </SheetDescription>
                        </SheetHeader>
                    </SheetContent>
                </Sheet>


            </div>
            <TableComponent list={data?.response.message} />
        </div>
    )
}

export default Project