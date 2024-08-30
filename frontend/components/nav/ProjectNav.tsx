import React, { useState } from 'react'
import { useRouter } from 'next/router'
import { Input } from '../ui/input'
import { Label } from '../ui/label'
import { Flag } from 'lucide-react'
import {
    Sheet,
    SheetContent,
    SheetDescription,
    SheetHeader,
    SheetTitle,
    SheetTrigger,
} from "@/components/ui/sheet"
import { Button } from '../ui/button'
import GenerateTokenComp from '../Project/GenerateTokenComp'
import Invite from '../Project/Invite'
import { z } from 'zod'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { HTTPRequest } from '@/api/api'
import { toast } from 'sonner'
import { useQueryClient } from '@tanstack/react-query'

const formSchema = z.object({
    flag: z.string().min(3, {
        message: "Length of name should be greater than 3"
    }),
    value: z.string().min(3, {
        message: "Length of name should be greater than 3"
    }),
})



interface props {
    id: string
}

const ProjectNav: React.FC<props> = ({ id }) => {
    const [open,setOpen] = useState<boolean>(false)
    const queryClient = useQueryClient()

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            flag: "",
            value: ""
        },
    })

    const createFlag = async (payload: z.infer<typeof formSchema>) => {
        const obj = {
            pid: id,
            flag: payload.flag,
            value: payload.value
        }
        const res = await HTTPRequest("/flags/create", { body: JSON.stringify(obj) }, "POST")
        if (res?.response.success) {
            queryClient.invalidateQueries({ queryKey: ["flags"] })
            toast.success("Flag created successfully")
        } else {
            toast.error(res?.response.message)
        }
        form.reset()
        setOpen(false)
    }


    return (
        <div className='flex md:flex-row flex-col gap-2 items-center '>
            <GenerateTokenComp id={id} />
            <Invite pid={id} />
            <Sheet open = {open} onOpenChange={setOpen}>
                <SheetTrigger className='w-full md:w-auto'><Button className='w-full md:w-auto' variant={"secondary"}>Create Flag</Button></SheetTrigger>
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
    )
}

export default ProjectNav