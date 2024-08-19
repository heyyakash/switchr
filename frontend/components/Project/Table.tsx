import React from 'react'
import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
  } from "@/components/ui/table"
  
export interface FlagRowInterface {
    fid: string,
    name: string,
    value: string,
    createdAt: string,
    createdBy: string
}

export interface FlagTableInterface{
    list : FlagRowInterface[]
}

const TableComponent: React.FC<FlagTableInterface> = (props) => {


  return (
    <Table className=''>
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
    {props.list.map((x,i)=>{
        return(
            <TableRow key = {i}>
            <TableCell className="font-medium text-ellipsis whitespace-nowrap">{x.fid}</TableCell>
            <TableCell>{x.name}</TableCell>
            <TableCell>{x.value}</TableCell>
            <TableCell>{x.createdBy}</TableCell>
            <TableCell>{x.createdAt}</TableCell>
          </TableRow>
        )
    })}

  </TableBody>
</Table>

  )
}

export default TableComponent