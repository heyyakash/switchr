import { HTTPRequest } from '@/api/api'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import React, { useState } from 'react'
import Loading from '../Loading/Loading'
import { Button } from '../ui/button'
import { Label } from '../ui/label'
import { Input } from '../ui/input'
import { SubmitHandler, useForm } from 'react-hook-form'
import Members from './Members'
import { toast } from 'sonner'
import { useRouter } from 'next/router'
import { RotateCw } from 'lucide-react'

interface props {
    pid: string
}

interface ProjectName {
    name :string
}

const Settings: React.FC<props> = ({ pid }) => {
    const queryClient = useQueryClient()
    const { data, isError, isLoading } = useQuery({
        queryKey: [`project-${pid}`],
        queryFn: async () => {
            return (await HTTPRequest(`/project/${pid}`, {}, "GET"))
        },
    })

    const {handleSubmit, register } = useForm<ProjectName>()


    const HandleProjectNameChange: SubmitHandler<ProjectName> =async (data) => {
        console.log(data)
    }
    const router = useRouter()

    const [deleting, setDeleting] = useState(false)

    const DeleteProject = async () => {
        setDeleting(true)
        const res = await HTTPRequest(`/project/${pid}`,{},"DELETE")
        if(res?.response.success){
            queryClient.invalidateQueries({
                queryKey:["projects"]
            })
            toast.success(res?.response.message)
            router.push("/dashboard")

        }else{
            toast.error(res?.response.message)
        }
        setDeleting(false)
    }


    const [editBoxOpen , setEditBoxOpen] = useState(false)
    if (isLoading) return <Loading />
    if (isError) return <div>Error Occuered</div>
    return (
        <div className='w-full'>
            <div className='mb-4 flex items-center justify-between'>
                <h2 className='text-xl font-semibold'>{data?.response.message.name} Settings</h2>
            </div>
            <div className='mt-12'>
                <h3 className='text-xl'>Edit Project Name</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>You can change your current project name to another unique project name</p>
                <form onSubmit={handleSubmit(HandleProjectNameChange)} className='flex flex-col gap-3 items-start'>
                    <Label>Name</Label>
                    <Input type = "text" {...register("name")} placeholder= {data?.response.message?.name} className='max-w-[500px] w-full' />
                    <Button >Update</Button>
                </form>
            </div>
            
            <div className='mt-6'>
                <h3 className='text-xl'>Edit Members</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>Owners and Editors have the authority to remove members from the project</p>
                <Members pid= {pid} />            
            

            </div>
            <div className='mt-6'>
                <h3 className='text-xl'>Delete Project</h3>
                <hr className='my-3' />
                <p className='text-primary/70 mb-3'>Once you've deleted your project you and its members cannot access it again. This action can only be taken by the owner(s) of the project.</p>
                <Button onClick={()=>DeleteProject()} disabled = {deleting} variant={"destructive"}>
                    {deleting? <RotateCw className='animate-spin' />: "Delete Project"}
                </Button>
                
            </div>
        </div>
    )
}

export default Settings