import React from 'react'
import { Button } from '../ui/button'
import ProjectCard, { ProjectCardInterface } from '../CardComponent/ProjectCard'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { HTTPRequest } from '@/api/api'
import {
    Sheet,
    SheetContent,
    SheetDescription,
    SheetHeader,
    SheetTitle,
    SheetTrigger,
} from "@/components/ui/sheet"
import { Label } from '../ui/label'
import { Input } from '../ui/input'
import { DownloadCloud } from 'lucide-react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'
import { toast } from 'sonner'
import { Form } from '../ui/form'


const formSchema = z.object({
    name: z.string().min(8, {
        message: "Length of name should be greater than 8"
    }),
})


const Dashboard = () => {
    const queryClient = useQueryClient()
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: ""
        },
    })
    async function CreateProject(values: z.infer<typeof formSchema>) {
        console.log("clicked")
        const res = await HTTPRequest(
            "/project/create",
            { body: JSON.stringify(values) },
            "POST"
        )
        if (res.response.success) {
            toast.success(res.response.message)
            queryClient.invalidateQueries({
                queryKey:["projects"]
            })
        } else {
            toast.error(res.response.message)
        }
    }
    const Project: ProjectCardInterface[] = [
        {
            name: "Heavenly Project",
            link: "www.google.com",
            createdAt: "29th Jan, 2002",
            createdBy: "Akash Sharma",
            role: "Creator",
            flags: 24,
        },
        {
            name: "Heavenly Project",
            link: "www.google.com",
            createdAt: "29th Jan, 2002",
            createdBy: "Akash Sharma",
            role: "Creator",
            flags: 24,
        },
        {
            name: "Heavenly Project",
            link: "www.google.com",
            createdAt: "29th Jan, 2002",
            createdBy: "Akash Sharma",
            role: "Creator",
            flags: 24,
        },

    ]
    const { data, error, isLoading } = useQuery({
        queryKey: ["projects"],
        queryFn: async () => {
            return (await HTTPRequest("/userprojectmap", {}, "GET"))
        }
    })
    console.log(typeof data?.response.message, data?.response.message[0])
    return (
        <div>
            <div className='flex items-center justify-between mb-5'>
                <h2 className='text-xl font-semibold'>Your Projects</h2>

                <Sheet>
                    <SheetTrigger><Button>Create New</Button></SheetTrigger>
                    <SheetContent>
                        <SheetHeader className='flex flex-col gap-5'>
                            <SheetTitle className='flex gap-2 text-xl items-center'> <DownloadCloud />Create a new Project</SheetTitle>
                            <SheetDescription className='mt-5'>
                                    <form onSubmit={form.handleSubmit(CreateProject)} className='flex flex-col gap-4'>
                                        <div className="grid w-full max-w-sm items-center gap-1.5">
                                            <Label htmlFor="name" className='mb-1'>Name</Label>
                                            <Input type="name" id="name" {...form.register("name")} placeholder="name" />
                                        </div>

                                        <Button type="submit" variant={"default"} className='w-full mt-2'>Create Project</Button>

                                    </form>
               
                            </SheetDescription>
                        </SheetHeader>
                    </SheetContent>
                </Sheet>
            </div>

            <div className='w-full grid gap-6 grid-cols-3 grid-rows-auto'>
                {Project.map((x, i) => {
                    return (
                        <ProjectCard key={i} name={x.name} link={x.link} createdAt={x.createdAt} createdBy={x.createdBy} flags={x.flags} role={x.role} />
                    )
                })}
                {data && data.response.message.length>0 ? 
                
                data.response.message?.map((x:any,y:number)=>{
                    return(
                        <ProjectCard key = {y} name = {x?.Project.name} link = {x.pid} createdBy={x?.Project?.createdBy} owned = {x?.Project?.createdBy===x?.uid} createdAt='12 Jan 2002' flags = {1} role = {x.role_id} />
                    )
                })
                :(<></>)}
            </div>
        </div>
    )
}

export default Dashboard