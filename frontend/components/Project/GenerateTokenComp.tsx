import React, { useState } from 'react'
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog"
import { Button } from '../ui/button'
import { RotateCw } from 'lucide-react'
import { Input } from '../ui/input'
import { toast } from 'sonner'
import { CLIENT_STATIC_FILES_RUNTIME_REACT_REFRESH } from 'next/dist/shared/lib/constants'
import { HTTPRequest } from '@/api/api'

interface props {
    id: string
}

const GenerateTokenComp: React.FC<props> = (props) => {
    const [generating, setGenerating] = useState(false)
    const [open,setOpen] = useState(false)
    const [token, setToken] = useState<string | null>(null)
    const copyToClipBoard = (e: React.MouseEvent<HTMLInputElement, MouseEvent>) => {
        e.preventDefault()
        if (token) {
            navigator.clipboard.writeText(token).then(() => {
                toast.success("Copied")
            }).catch((err) => {
                console.log(err)
                toast.error("Failed to copy")
            })
        }
    }
    const generateToken = async () => {
        setGenerating(false)
        setToken(null)
        setGenerating(true)
        const res = await HTTPRequest(`/api/create/${props.id}`, {}, "GET")
        if (res?.response.success) {
            setGenerating(false)
            setToken(res.response.message)
        } else {
            toast.error("Some Error Occuered")
        }
        setGenerating(false)
    }
    return (
        <Dialog open = {open } onOpenChange={setOpen}>
            <DialogTrigger>
                <Button>How to access?</Button>
            </DialogTrigger>
            <DialogContent>
                <DialogHeader>
                    <DialogTitle className='text-2xl w-full'>Retrieving Flags</DialogTitle>
                    </DialogHeader>
                    <hr />
                    <div className=' flex flex-col gap-3'>
                        In order to retrieve the flags you need to generate an access token and send it with a GET request
                        <br />
                        <Button onClick={() => generateToken()} className=''>{generating ? (<RotateCw size={20} className='animate-spin' />) : ("Generate Token")}</Button>
                        {token && <Input className='cursor-pointer' type="text" value={token} onClick={(e) => copyToClipBoard(e)} />}
                        <hr />
                        <p>Example</p>
                        <textarea value={
                            `curl  -X GET \
      'http://localhost:8020/api/get/CREATE_FLAG'
      --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQ2ODA1OTksIlVpZCI6IiIsIlR5cGUiOiJhcGktdG9rZW4iLCJFbWFpbCI6IiIsIlBpZCI6IjcwNDhkMzE0LTUwNzQtNDBhZi05YTgxLWI1MjA5MmIwMmNkMCJ9.5at5eYJwbolz4M1eIS5GfT6RzHUw4xZzZ5t52z6FkB4'`
                        } className='p-2 bg-secondary h-[250px] rounded-lg w-full overflow-auto whitespace-pre-wrap'>

                        </textarea>
                    </div>
                
            </DialogContent>
        </Dialog>

    )
}

export default GenerateTokenComp