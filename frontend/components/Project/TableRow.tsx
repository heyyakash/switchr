import { HTTPRequest } from '@/api/api'
import { zodResolver } from '@hookform/resolvers/zod'
import { useQueryClient } from '@tanstack/react-query'
import React, { useState } from 'react'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'
import { FlagRowInterface } from './Table'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Label } from '../ui/label'
import { Input } from '../ui/input'
import { Button } from '../ui/button'
import { Pause } from 'lucide-react'
import { TableCell, TableRow } from '../ui/table'

const formSchema = z.object({
    value: z.string().min(3, {
      message: "Length of name should be greater than 3"
    }),
})

const TableRowComponent: React.FC<FlagRowInterface> = (props) => {
    const [open, setOpen] = useState(false)
    const queryClient = useQueryClient()
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
          value: "",
        },
      })
    
  const UpdateFlag = async  (payload:any) => {
    const obj = {
      value : payload.value
    }
    const res = await HTTPRequest(`/flags/${props.fid}`, {body:JSON.stringify(obj)}, "PATCH")
    if (res?.response.success){
      toast.success(res.response.message)
      queryClient.invalidateQueries({
        queryKey:["flags"]
      })
      setOpen(false)
    }else{
      toast.error(res?.response.message)
    }
  }
  
  const DeleteFlag = async  () => {
    const res = await HTTPRequest(`/flags/${props.fid}`,{}, "DELETE")
    if (res?.response.success){
      toast.success(res.response.message)
      queryClient.invalidateQueries({
        queryKey:["flags"]
      })
      setOpen(false)
    }else{
      toast.error(res?.response.message)
    }
  }
  return (
    <>
    <TableRow className='cursor-pointer' onClick={() => setOpen(true)}>
      <TableCell className="font-medium  whitespace-nowrap">{props.fid?.length > 10 ? props.fid?.substring(0, 10) + "..." : props.fid}</TableCell>
      <TableCell>{props.name}</TableCell>
      <TableCell>{props.value}</TableCell>
      <TableCell>{props.full_name}</TableCell>
      <TableCell>{(new Date(props.createdAt)).toString().substring(0, 15)}</TableCell>
    </TableRow>
    <Dialog defaultOpen={false} onOpenChange={setOpen} open={open} >
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Flag</DialogTitle>
          <DialogDescription>
            <form onSubmit={form.handleSubmit(UpdateFlag)} className='flex w-full mt-4 flex-col gap-4'>
              <div className="grid w-full  items-center gap-1.5">
                <Label htmlFor="fid" className='mb-1'>fid</Label>
                <Input type="fid" id="fid"  value={props.fid} disabled placeholder="fid" />
              </div>
              <div className="grid w-full  items-center gap-1.5">
                <Label htmlFor="flag" className='mb-1'>Flag</Label>
                <Input type="input" id="flag" value={props.name} disabled placeholder="Key" />
              </div>
              <div className="grid w-full  items-center gap-1.5">
                <Label htmlFor="Value" className='mb-1'>Value</Label>
                <Input type="Value" id="Value" {...form.register("value")} placeholder={props.value} />
              </div>
              <div className='grid grid-cols-2 gap-3'>
                <Button type = "submit" variant={"default"} className='w-full mt-2'>Update</Button>
                <Button variant={"destructive"} onClick={()=>DeleteFlag()} className='w-full mt-2'>Delete</Button>
              </div>

            </form>
          </DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>


  </>
  )
}

export default TableRowComponent