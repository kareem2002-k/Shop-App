import React from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom';

function Home() {
    const navigate = useNavigate();
  return (
    <div>
        <button
        className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'
        onClick={
           async () => {
               await axios.post('http://localhost:8000/logout', {}, {
                    withCredentials: true
                }).then((response) => {
                    console.log(response)
                
                    if (response.status === 200) {
                        // User is authenticated, navigate to home page
                        navigate('/');
                    }
                    else {
                        console.log('User did not logout');
                    }

                    
                })
            }
           
        }

        >Logout</button>
    </div>
  )
}

export default Home