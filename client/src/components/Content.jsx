import React from 'react'

const Content = () => {
  return (
    <section className='p-4 bg-gray-200 h-screen'>
        <div className=''>
        <div>
            <form className='flex flex-col gap-4'>
                <input type='text' className='p-2 w-full h-32 rounded-xl shadow-md' />
                <button className='ml-auto px-4 py-2 bg-green-400 text-white text-lg font-semibold rounded-md'>Post</button>
            </form>
        </div>
        <div>
            
        </div>
        </div>
    </section>
  )
}

export default Content