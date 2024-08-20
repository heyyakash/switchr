import React from 'react'
import { Button } from '../ui/button'
import ProjectCard, { ProjectCardInterface } from '../CardComponent/ProjectCard'


const Dashboard = () => {
    const Project: ProjectCardInterface[] = [
        {
            name:"Heavenly Project",
            link:"www.google.com",
            createdAt:"29th Jan, 2002",
            createdBy:"Akash Sharma",
            role:"Creator",
            flags:24,
        },
        {
            name:"Heavenly Project",
            link:"www.google.com",
            createdAt:"29th Jan, 2002",
            createdBy:"Akash Sharma",
            role:"Creator",
            flags:24,
        },
        {
            name:"Heavenly Project",
            link:"www.google.com",
            createdAt:"29th Jan, 2002",
            createdBy:"Akash Sharma",
            role:"Creator",
            flags:24,
        },

    ]
  return (
    <div>
        <div className='flex items-center justify-between mb-5'>
        <h2 className='text-xl font-semibold'>Your Projects</h2>
        <Button>Create New</Button>
        </div>

        <div className='w-full grid gap-6 grid-cols-3 grid-rows-auto'>
            {Project.map((x,i)=>{
                return (
                    <ProjectCard key = {i} name={x.name} link = {x.link} createdAt={x.createdAt} createdBy={x.createdBy} flags={x.flags} role = {x.role} /> 
                )
            })}
        </div>
    </div>
  )
}

export default Dashboard