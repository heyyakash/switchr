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
import TableRowComponent from './TableRow'



export interface FlagRowInterface {
  fid: string,
  name: string,
  value: string,
  createdAt: string,
  createdBy: string
  full_name: string
}


export interface FlagTableInterface {
  list: any
}

const TableComponent: React.FC<FlagTableInterface> = (props) => {


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
        {Array.isArray(props.list) && props.list?.map((x: any, i: number) =>  <TableRowComponent full_name= {x?.full_name} fid = {x?.fid} name = {x?.flag} value = {x?.value} createdAt= {x?.createdAt} createdBy= {x.createdBy} key= {i} /> )}

      </TableBody>
    </Table>

  )
}

export default TableComponent