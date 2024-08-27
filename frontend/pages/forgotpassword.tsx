import { HTTPRequest } from '@/api/api'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { RotateCw } from 'lucide-react'
import React, { useState } from 'react'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'

interface payload {
  email: string
}

const ForgotPassword = () => {
  const [sending, setSending] = useState(false)

  const {register , handleSubmit, reset} = useForm<payload>()
  const onSubmit = async (data: payload) => {
    setSending(true)
    const res = await HTTPRequest(`/user/forgot`, {body: JSON.stringify(data)}, "POST" )
    if(res?.response.success){
        toast.success("Kindly check your registered email inbox")
    }else{
        toast.error(res?.response.message)
    }
    reset()
    setSending(false)
  }


  return (
    <div className='max-w-[1200px] w-full mx-auto min-h-[100vh] flex justify-center items-center'>
    
    <form onSubmit={handleSubmit(onSubmit)} className=" w-[90%] md:w-[450px]">
        <div className="my-6">
            <h3 className=" text-2xl font-bold text-center">Forgot Password</h3>
            <p className="text-md text-white/50 text-center mt-3">Enter your email to receive reset password link</p>
        </div>
        <Label htmlFor='email'>Email</Label>
        <Input className="input-primary h-[40px] text-[1rem]" type="email" placeholder="johndoe@gmail.com" {...register("email")} id = "email" />
        <Button type="submit" size={"lg"} className="mt-6 text-lg w-full" variant={"default"}>{sending ? (<RotateCw size={20} className='animate-spin' />) : ("Send Reset Link")}</Button>
    </form>
    </div>
  )
}

export default ForgotPassword



ForgotPassword.getLayout = () => {
  return (
      <></>
  )
}