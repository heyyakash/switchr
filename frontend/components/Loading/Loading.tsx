import { Key } from 'lucide-react'
import React from 'react'

const Loading = () => {
  return (
    <div className='absolute inset-0 flex items-center justify-center'>
        <Key size = {50} className='animate-ping' />
    </div>
  )
}

export default Loading


Loading.getLayout = () => {
    return (
        <></>
    )
}