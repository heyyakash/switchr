import React from 'react'
import { Button } from '../ui/button'
import TableComponent, { FlagRowInterface } from './Table'
import {
    Sheet,
    SheetContent,
    SheetDescription,
    SheetHeader,
    SheetTitle,
    SheetTrigger,
} from "@/components/ui/sheet"

import { Input } from '../ui/input'
import { Label } from '../ui/label'
import { Flag } from 'lucide-react'
import { useQuery } from '@tanstack/react-query'
import { HTTPRequest } from '@/api/api'



const Project = () => {
    const Flags: FlagRowInterface[] = [
        {
            fid: "sadasdf-sada21-asd213-232134",
            name: "SHOW_FEATURE",
            value: "TRUE",
            createdAt: "25th JUL 2025",
            createdBy: "Akash Sharma",
        },
        {
            fid: "sadasdf-sada21-asd213-232134",
            name: "SHOW_FEATURE",
            value: "TRUE",
            createdAt: "25th JUL 2025",
            createdBy: "Akash Sharma",
        },
        {
            fid: "sadasdf-sada21-asd213-232134",
            name: "SHOW_FEATURE",
            value: "TRUE",
            createdAt: "25th JUL 2025",
            createdBy: "Akash Sharma",
        },
        {
            fid: "sadasdf-sada21-asd213-232134",
            name: "SHOW_FEATURE",
            value: "TRUE",
            createdAt: "25th JUL 2025",
            createdBy: "Akash Sharma",
        },
    ]

    return (
        <div>
            <div className='mb-4 flex items-center justify-between'>
                <h2 className='text-xl font-semibold'>Heavenly Project</h2>
                <Sheet>
                    <SheetTrigger><Button variant={"default"}>Create Flag</Button></SheetTrigger>
                    <SheetContent>
                        <SheetHeader className='flex flex-col gap-5'>
                            <SheetTitle className='flex gap-2 text-xl items-center'> <Flag />Create a new Flag</SheetTitle>
                            <SheetDescription className='mt-5'>
                            <div className='flex flex-col gap-4'>
                                <div className="grid w-full max-w-sm items-center gap-1.5">
                                    <Label htmlFor="Key" className='mb-1'>Key</Label>
                                    <Input type="Key" id="Key" placeholder="Key" />
                                </div>
                                <div className="grid w-full max-w-sm items-center gap-1.5">
                                    <Label htmlFor="Value" className='mb-1'>Value</Label>
                                    <Input type="Value" id="Value" placeholder="Value" />
                                </div>
                                <Button variant={"default"} className='w-full mt-2'>Create</Button>

                            </div>
                            </SheetDescription>
                        </SheetHeader>
                    </SheetContent>
                </Sheet>


            </div>
            <TableComponent list={Flags} />
        </div>
    )
}

export default Project