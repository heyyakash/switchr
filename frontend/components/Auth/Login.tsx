import React from 'react'
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import * as z from "zod"

const formSchema = z.object({
    email: z.string().min(2, {
        message: "Length of email address should be greater than 2"
    }),
    password: z.string().min(2).max(50)
})

import { toast } from 'sonner'
import { useRouter } from 'next/router'
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import EmailPassword from './EmailPassword'
import MagicLinks from './MagicLink'


const Login = () => {
    const router = useRouter()

    return (
        <>
            <Tabs defaultValue="password" className="w-[90%] md:w-[450px]">
            <TabsContent value='password'><EmailPassword /></TabsContent>
            <TabsContent value='magiclink'><MagicLinks /></TabsContent>
                <TabsList className='w-full grid grid-cols-2 mt-10   '>
                    <TabsTrigger value="password">Password</TabsTrigger>
                    <TabsTrigger value="magiclink">Magic Link</TabsTrigger>
                </TabsList>
            </Tabs>
        </>
    )
}


export default Login