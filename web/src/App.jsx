// import { useState } from 'react'
// import reactLogo from './assets/react.svg'
// import viteLogo from '/vite.svg'
import './index.css'; 
import './app.css';
import UrlForm from './components/UrlForm';


export default function App(){
  return(
    <div className='min-w-full h-screen bg-beige-taupe grid grid-rows-3 auto-rows-fr'>
      <div className='flex flex-col row-start-2 items-center'>

        <div className=" ">
          <h1 className="text-5xl text-sage-green font-bold drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,0.8)]">URL SHORTENER</h1>
        </div>

        
          <UrlForm/>
      

      </div>
     
    
    </div>
  )
}