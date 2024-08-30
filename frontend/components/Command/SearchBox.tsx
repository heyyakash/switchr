import React, { useEffect, useState } from 'react'
import {
  Command,
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command"
import { useQuery } from '@tanstack/react-query'
import { HTTPRequest } from '@/api/api'
import Loading from '../Loading/Loading'
import { useRouter } from 'next/router'


export function SearchBox() {
  const [open, setOpen] = useState(false)
  const router = useRouter()

  useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === "k" && (e.metaKey || e.ctrlKey)) {
        e.preventDefault()
        setOpen((open) => !open)
      }
    }
    document.addEventListener("keydown", down)
    return () => document.removeEventListener("keydown", down)
  }, [])

  const PushToProject = (pid: string )=> {
    window.location.href = `/project/${pid}`; // Forces a full page reload
    setOpen(false)
  }

  const { data, error, isLoading } = useQuery({
    queryKey: ["projects"],
    queryFn: async () => {
      return (await HTTPRequest("/userprojectmap", {}, "GET"))
    }
  })


  return (
    <Command className='w-auto'>
    <CommandDialog open={open} onOpenChange={setOpen} >
      <CommandInput placeholder="Type a command or search..." />
      <CommandList>
        <CommandEmpty>No results found.</CommandEmpty>
        <CommandGroup heading="Projects" className='text-xl'>
          {isLoading ? <Loading /> :
            Array.isArray(data?.response.message) && data?.response.message.map((project, index) => {
              return (
                <CommandItem onSelect={()=>PushToProject(project.pid)} key={index}>{project.Project.name}</CommandItem>
              )
            })}
        </CommandGroup>
      </CommandList>
    </CommandDialog>
    </Command>
  )
}
