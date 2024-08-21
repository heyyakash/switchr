import React, { useState } from 'react'
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { responseInterface } from '@/api/api'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"


export interface FlagRowInterface {
  fid: string,
  name: string,
  value: string,
  createdAt: string,
  createdBy: string
}

export interface FlagTableInterface {
  list: any
}

const TableComponent: React.FC<FlagTableInterface> = (props) => {
  const [open, setOpen] = useState(false)

  return (
    <Table className='w-full'>
      <TableCaption>A list of flags associated with Project</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">F-ID</TableHead>
          <TableHead>Name</TableHead>
          <TableHead>Value</TableHead>
          <TableHead>Created By</TableHead>
          <TableHead>Created At</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {props.list?.map((x: any, i: number) => {
          return (
            <>
              <TableRow onClick={() => setOpen(true)} key={i}>
                <TableCell className="font-medium  whitespace-nowrap">{x.fid?.length > 10 ? x.fid?.substring(0, 10) + "..." : x.fid}</TableCell>
                <TableCell>{x.flag}</TableCell>
                <TableCell>{x.value}</TableCell>
                <TableCell>{x.full_name}</TableCell>
                <TableCell>{(new Date(x.createdAt)).toString().substring(0, 15)}</TableCell>
              </TableRow>
              <Dialog defaultOpen={false} onOpenChange={setOpen} open={open} >
                <DialogContent>
                  <DialogHeader>
                    <DialogTitle>Are you absolutely sure?</DialogTitle>
                    <DialogDescription>
                      This action cannot be undone. This will permanently delete your account
                      and remove your data from our servers.
                    </DialogDescription>
                  </DialogHeader>
                </DialogContent>
              </Dialog>


            </>
          )
        })}

      </TableBody>
    </Table>

  )
}

export default TableComponent