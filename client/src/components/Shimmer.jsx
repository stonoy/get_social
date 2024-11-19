import React from 'react'

const Shimmer = () => {
  return (
    <article  className='flex flex-col animate-pulse gap-2 bg-white my-4 shadow-md rounded-md p-2 md:my-6 md:p-4 hover:shadow-xl'>
                    <div className='flex gap-2 justify-start items-center'>
                        <h1 className='text-xl h-4 rounded-full bg-slate-300 capitalize px-4 py-2 font-semibold'></h1>
                        <div className='flex flex-col h-2 items-start'>
                            <h1 className='text-sm text-slate-700'></h1>
                            <h1 className='text-sm text-slate-600'></h1>
                        </div>
                    </div>
                    <div className='p-4 h-12 text-slate-600'>
                        <h1></h1>
                    </div>
                    <div className='flex px-2 h-4 justify-between items-center cursor-pointer'>
                        <div className='flex w-1/2 gap-1 bg-grey-300 border-r-2 border-gray-400 items-center justify-center'>
                        <h1 className='text-md'> </h1>
                        <h1 className='text-md'> 
                        </h1>
                        </div>
                        <div className='flex w-1/2 gap-1 bg-grey-300 items-center justify-center'>
                        <h1 className='text-md'></h1>
                        <h1 className='text-md'> 
                        </h1>
                        </div>
                    </div>
                </article>
  )
}

export default Shimmer