import Nav from '@/components/LandingPage/Nav'
import { Badge } from '@/components/ui/badge'
import React from 'react'

const Docs = () => {
  return (
    <>
        <Nav />
        <div className='max-w-[1200px] py-[6rem] w-full mx-auto'>
            <h2 className='text-3xl font-semibold text-center pt-[4rem]'>
                Accessing feature flags using REST API
            </h2>
            <div className='text-primary/70 mt-12 text-xl px-4 flex flex-col gap-3'>
                <p>The following procedures can be followed once a project is created and some flags are added in it</p>
                <p><span className='font-bold'>Step 1</span>. Login and open the project whose keys you wish to access and the click on <Badge>API</Badge> button on the project page.</p>
                <p><span className='font-bold'>Step 2</span>. Once a dialog box appears click on <Badge>Generate Key</Badge> button, this will generate key to access the flag.</p>
                <p><span className='font-bold'>Step 3</span>. Once the key is generated, it is valid for 120 days and shall be sent in the header called <span className='font-extrabold text-primary'>token</span> in a GET Request to the following URL</p>
                <div className='p-3 w-full overflow-auto bg-secondary/50 rounded-lg'>
                    {process.env.NEXT_PUBLIC_BASE_URL}/api/get/&lt;FLAG_NAME&gt;
                </div>
                <b>Example Request</b>
                <div className='p-3 w-full overflow-auto bg-secondary/50 rounded-lg'>
                curl -X GET &apos;{process.env.NEXT_PUBLIC_BASE_URL}/api/get/FLAG_NAME&apos;
                --header &apos;token: &lt;YOUR_TOKEN / KEY&gt;&apos;;
                </div>
            </div>
        </div> 
    </>
  )
}

export default Docs


Docs.getLayout = () => {
    return <></>
  }