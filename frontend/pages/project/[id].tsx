import Project from '@/components/Project/Project'
import { GetServerSideProps } from 'next'
import React from 'react'
import dynamic from 'next/dynamic';

const NoSSRComponent = dynamic(() => import('@/components/Project/Project'), { ssr: false });
interface props{
  id: string
}


const ProjectContainer: React.FC<props> = (props) => {
  return (
    <div className='max-w-[1200px] w-full mx-auto p-4 px-6'>
   <NoSSRComponent id = {props.id} />
   </div>
  )
}

export default ProjectContainer

export const getServerSideProps: GetServerSideProps<props> = async (context) => {
  // Fetch data using router.query.id
  const id = context.query.id as string;

  // Pass data to the component as props
  return {
    props: {
      id // Pass id as props
    }
  };
}

