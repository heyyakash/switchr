import Project from '@/components/Project/Project'
import { GetServerSideProps } from 'next'
import React, { useEffect, useState } from 'react'
import dynamic from 'next/dynamic';
import { useRouter } from 'next/router';
import Loading from '@/components/Loading/Loading';

const NoSSRComponent = dynamic(() => import('@/components/Project/Project'), { ssr: false });
interface props{
  id: string
}


const ProjectContainer: React.FC<props> = (props) => {
  const [id, setId] = useState<string | null>(null)
  const router = useRouter()
  useEffect(()=>{
    setId(router?.query?.id as string)
  },[router.query.id])

  if(!id) return <Loading />
  return (
    <div className='max-w-[1200px] w-full mx-auto p-4 px-6'>
   <NoSSRComponent id = {id} />
   </div>
  )
}

export default ProjectContainer

// export const getServerSideProps: GetServerSideProps<props> = async (context) => {
//   const id = context.query.id as string;
//   return {
//     props: {
//       id
//     }
//   };
// }

