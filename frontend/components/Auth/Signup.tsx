import React from 'react'
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import * as z from "zod"



const formSchema = z.object({
    fullname: z.string().min(5).max(50),
    email: z.string().min(2, {
        message: "Length of email address should be greater than 2"
    }),
    password: z.string().min(2).max(50)
})


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
import { toast } from 'sonner'
import { HTTPRequest } from '@/api/api'

const SignUp = () => {
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            fullname :"",
            email: "",
            password: ""
        },
    })

    async function onSubmit(values: z.infer<typeof formSchema>) {
        const payload = {
            email:values.email,
            fullname:values.fullname,
            password:values.password
        }
        const res = await HTTPRequest("/user/create", {
            body: JSON.stringify(payload)
        },"POST")
        
        if(res.response.success){
            toast.success("You have signed up successfully!!")
            form.reset({
                fullname:"",
                email:"",
                password:""
            })
        }else{
            toast.error("Sign Up failed")
        }
        
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="w-[90%] md:w-[450px] ">
                <div className="my-6">
                    <h3 className=" text-4xl font-bold text-center">Create an account</h3>
                    <p className="text-lg text-white/50 text-center mt-3">Enter your email and password to sign up</p>
                </div>

                <FormField
                    control={form.control}
                    name="fullname"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel className="">Full Name</FormLabel>
                            <FormControl>
                                <Input className="input-primary h-[40px] text-[1rem]" type="text" placeholder="John Doe" {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="email"
                    render={({ field }) => (
                        <FormItem className='mt-4'>
                            <FormLabel className="">Email</FormLabel>
                            <FormControl>
                                <Input className="input-primary h-[40px] text-[1rem]" type="email" placeholder="johndoe@gmail.com" {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <FormField
                    control={form.control}
                    name="password"
                    render={({ field }) => (
                        <FormItem className="mt-4">
                            <FormLabel className="">Password</FormLabel>
                            <FormControl>
                                <Input type="password" className="input-primary h-[40px] text-[1rem]" placeholder="password" {...field} />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <Button type="submit" size={"lg"} className="mt-6 text-lg w-full" variant={"default"}>Sign up</Button>
            </form>
        </Form>
    )
}

export default SignUp