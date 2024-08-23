import React, { useState } from 'react'
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { useQuery } from '@tanstack/react-query'
import { HTTPRequest } from '@/api/api'
import { z } from "zod"
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { Button } from "@/components/ui/button"
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { toast } from 'sonner'
import { RotateCw } from 'lucide-react'


const formSchema = z.object({
    email: z.string().min(2).max(50),
    role: z.string(),
})

interface props {
    pid: string
}

const Invite: React.FC<props> = ({ pid }) => {
    const [inviting, setInviting] = useState(false)
    const { data, isLoading, isError } = useQuery({
        queryKey: ["roles"],
        queryFn: async () => {
            return (await HTTPRequest("/roles/list", {}, "GET"))
        }
    })

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            email: "",
            role: "2",
        },
    })

    const onSubmit = async (values: z.infer<typeof formSchema>) => {
        setInviting(true)
        const payload = {
            ...values,
            pid : pid,
            role: parseInt(data?.response.message[values.role], 10)
        }
        const res = await HTTPRequest(`/share`, { body: JSON.stringify(payload) }, "POST")
        if (res?.response.success) {
            toast.success(res.response.message)
        } else {
            toast.error(res?.response.message)
        }
        setInviting(false)
    }

    return (
        <Dialog>
            <DialogTrigger>
                <Button>Add Member</Button>
            </DialogTrigger>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle>Invite a person</DialogTitle>
                </DialogHeader>
                <DialogDescription>
                    <Form {...form}>
                        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                            <FormField
                                control={form.control}
                                name="email"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Email Address</FormLabel>
                                        <FormControl>
                                            <Input placeholder="Email Address" {...field} />
                                        </FormControl>
                                        <FormDescription>
                                            Enter the registered email of the person you want to invite to the project
                                        </FormDescription>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="role"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Role</FormLabel>
                                        <FormControl>
                                            <Select
                                                value={field.value}
                                                onValueChange={field.onChange} 
                                                disabled={isLoading || isError}
                                            >
                                                <SelectTrigger className="w-full    ">
                                                    <SelectValue placeholder="Select Role" />
                                                </SelectTrigger>
                                                <SelectContent>
                                                    {data?.response.message &&
                                                        Object.keys(data.response.message).map((x, i) => (
                                                            <SelectItem value={x} className="capitalize" key={i}>
                                                                {x}
                                                            </SelectItem>
                                                        ))}
                                                </SelectContent>
                                            </Select>
                                        </FormControl>
                                        <FormDescription>Enter the role you want to assign them.</FormDescription>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <Button type="submit" disabled = {inviting}>{inviting? (<RotateCw size={20} className='animate-spin' />):"Invite"}</Button>
                        </form>
                    </Form>

                </DialogDescription>

            </DialogContent>
        </Dialog>

    )
}

export default Invite