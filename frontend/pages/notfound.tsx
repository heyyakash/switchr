import { FileX2 } from 'lucide-react'
import React from 'react'

const NotFound = () => {
  return (
    <div className='w-full h-full flex items-center flex-col gap-5 justify-center'>
      <FileX2 size={70} />
        <h1 className='text-5xl font-extrabold'>Uh oh! Not Found</h1>
    </div>
  )
}

export default NotFound