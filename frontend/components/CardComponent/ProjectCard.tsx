import React, { useEffect } from 'react'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Badge } from '../ui/badge'
import { ArrowUpRight, Flag, Link } from 'lucide-react'
import { Button } from '../ui/button'

export interface ProjectCardInterface {
  name: string
  createdAt: string
  createdBy: string
  role: string
  link: string
  flags: number
}

const ProjectCard: React.FC<ProjectCardInterface> = (props) => {
  return (
    <div className='relative overflow-hidden w-full '>
      <Card className='relative bg-transparent z-[1]'>
        <CardHeader>
          <CardTitle>{props.name}</CardTitle>
          <CardDescription className='pt-2 flex items-center justify-between'>{props.createdBy} <Badge>{props.role}</Badge></CardDescription>
        </CardHeader>
        <CardContent>
          <p>{props.flags} Flags <Flag></Flag></p>
        </CardContent>
        <CardFooter className='w-full'>
          <Button variant={"default"} className='w-full'>Open <ArrowUpRight /></Button>
        </CardFooter>
      </Card>
  <div className='absolute z-[0] top-0 -rotate-90  text-[10rem] text-primary/10 font-bold '>{props.name}</div>
    </div>
  )
}

export default ProjectCard