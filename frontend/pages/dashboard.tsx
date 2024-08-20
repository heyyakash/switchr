"use client"


import { HTTPRequest } from '@/api/api'
import Dashboard from '@/components/Dashboard/Dashboard'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { useRouter } from 'next/router'
import React from 'react'

const DashboardLayout = () => {

    return (
        <div className='p-4 px-6 max-w-[1200px] w-full mx-auto'>
            <Dashboard />
        </div>
    )
}


export default DashboardLayout