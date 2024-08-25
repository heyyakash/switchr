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
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { HTTPRequest } from '@/api/api'
import Loading from '../Loading/Loading'
import { RotateCw, Trash2 } from 'lucide-react'
import { Button } from '../ui/button'
import { toast } from 'sonner'


interface props {
    pid: string
}

const Members: React.FC<props> = ({ pid }) => {
    const queryClient = useQueryClient()
    const { data: roles, isLoading: rolesIsLoading } = useQuery({
        queryKey: ["roles"],
        queryFn: async () => {
            return (await HTTPRequest(`/roles/list`, {}, "GET"))
        }
    })

    const { data, isLoading } = useQuery({
        queryKey: [`pid-${pid}-members`],
        queryFn: async () => {
            return (await HTTPRequest(`/userprojectmap/members/${pid}`, {}, "GET"))
        }
    })

    const [deleting, setDeleting] = useState(false)

    const DeleteMembers = async (id: string) => {
        setDeleting(true)
        if (id.length) {
            const payload = {
                memid: id
            }
            const res = await HTTPRequest(`/userprojectmap/members/delete/${pid}`, { body: JSON.stringify(payload) }, "POST")
            if (res?.response.success) {
                toast.success("Member removed")
                queryClient.invalidateQueries({
                    queryKey: [`pid-${pid}-members`]
                })
            } else {
                toast.error(res?.response.message)
            }
        }
        setDeleting(false)
    }

    if (isLoading || rolesIsLoading) return <Loading />

    return (
        <Table>
            <TableCaption>A list of all members of the project</TableCaption>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">Sl no.</TableHead>
                    <TableHead>Member Name</TableHead>
                    <TableHead>Member Email</TableHead>
                    <TableHead>Role</TableHead>
                    <TableHead className="text-right">Delete</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {Array.isArray(data?.response.message) && data?.response.message.map((x, i) => {
                    return (
                        <TableRow key={i}>
                            <TableCell className="font-medium">{i + 1}</TableCell>
                            <TableCell>{x?.User?.fullname}</TableCell>
                            <TableCell>{x?.User?.email}</TableCell>
                            <TableCell className='capitalize'>{Object.keys(roles?.response.message).find(key => roles?.response.message[key] === x?.role_id) || "Reader"}</TableCell>
                            <TableCell className="text-right ml-auto"><Button variant={"destructive"} disabled={deleting} onClick={() => DeleteMembers(x?.uid)}>
                                {deleting ? <RotateCw className='animate-spin' /> : <Trash2 />}

                            </Button></TableCell>
                        </TableRow>
                    )
                })}

            </TableBody>
        </Table>

    )
}

export default Members