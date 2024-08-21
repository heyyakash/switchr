'use client'; 

import { HTTPRequest } from '@/api/api'
import Dashboard from '@/components/Dashboard/Dashboard'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { useRouter } from 'next/router'
import React from 'react'
import dynamic from 'next/dynamic';

const NoSSRComponent = dynamic(() => import('@/components/Dashboard/Dashboard'), { ssr: false });

const DashboardLayout = () => {

    return (
        <div className='p-4 px-6 max-w-[1200px] w-full mx-auto'>
            <NoSSRComponent />
        </div>
    )
}


export default DashboardLayout