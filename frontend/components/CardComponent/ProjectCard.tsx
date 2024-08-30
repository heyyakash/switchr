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
import { ArrowUpRight, Flag } from 'lucide-react'
import { Button } from '../ui/button'
import Link from 'next/link'

export interface ProjectCardInterface {
  name: string
  createdAt: string
  owned: boolean
  role: string
  link: string
  flags: number
}

const ProjectCard: React.FC<ProjectCardInterface> = (props) => {
  return (
    <div className='relative overflow-hidden w-full '>
      <Card className='relative min-h-[270px] bg-transparent z-[1]'>
        <CardHeader>
          <CardTitle>{props.name}</CardTitle>
          <CardDescription className='pt-2 flex items-center justify-between'>{props.owned?(<Badge>Owned</Badge>):(<Badge variant={"secondary"}>Shared</Badge>)} <Badge className='ml-auto' variant={"outline"}>{props.createdAt}</Badge></CardDescription>
        </CardHeader>
        <CardContent>
          {/* <p>{props.flags} Flags <Flag></Flag></p> */}
        </CardContent>
        <CardFooter className='w-full absolute bottom-0 '>
          <Link href={`/project/${props.link}`}>
          <Button variant={"default"} className='w-full'>Open <ArrowUpRight /></Button>
          </Link>
        </CardFooter>
      </Card>
  <div className='absolute z-[0] top-0 -rotate-90  text-[10rem] text-primary/10 font-bold '>{props.name}</div>
    </div>
  )
}

export default ProjectCard