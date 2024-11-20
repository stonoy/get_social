import React from 'react'
import { Content, FollowersList, User } from '../components'

const Landing = () => {
  return (
    <>
        <User />
        <Content/>
        <FollowersList isBigScreen={true}/>
    </>
  )
}

export default Landing