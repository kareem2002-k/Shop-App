import axios from 'axios';
import React, { useEffect } from 'react'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom';

function Register() {
    const navigate = useNavigate();
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [username, setUsername] = useState('')

    useEffect(() => {
        // Check if the user is already authenticated
        
        axios
            .get('http://localhost:8000/user', {
                withCredentials: true
            })
            .then((response) => {
                if (response.status === 200) {
                    // User is authenticated, navigate to home page
                    navigate('/Home');
                }
            })
            .catch((error) => {
                // User is not authenticated, continue rendering the login page
                console.log('User not authenticated');
            });
    }, []);


    const registerHandler =  () => {
        axios.post('http://localhost:8000/register', {
            "email": email,
            "password": password,
            "username": username

        }, {
            withCredentials: true
        }).then((response) => {
         if (response.status === 200) {
               // User is authenticated, navigate to home page
               navigate('/Home');
             }

            else {
                console.log(response)
                console.log('User not authenticated');
            }

        })
    }
  return (
    <div>   
        <div className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 flex flex-col">
        <div className="mb-4">
            <label className="block text-grey-darker text-sm font-bold mb-2" for="username">
            Username
            </label>
            <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker"
            id="username"
            type="text"
            placeholder="Username"
            onChange={
                (e) => {
                    setUsername(e.target.value)
                }
            }

            />
        </div>
        <div className="mb-4">
            <label className="block text-grey-darker text-sm font-bold mb-2" for="email">
            Email
            </label>
            <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker"
            id="email"
            type="text"
            placeholder="Email"
            onChange={
                (e) => {
                    setEmail(e.target.value)
                }
            }
            />
        </div>
        <div className="mb-6">
            <label className="block text-grey-darker text-sm font-bold mb-2" for="password">
            Password
            </label>
            <input

            className="shadow appearance-none border border-red rounded w-full py-2 px-3 text-grey-darker mb-3"
            id="password"
            type="password"
            placeholder="******************"
            onChange={
                (e) => {
                    setPassword(e.target.value)
                }
            }
            />
            <p className="text-red text-xs italic">Please choose a password.</p>
        </div>
        <div className="flex items-center justify-between">
            <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            type="button"
            onClick={registerHandler}
            >
            Sign In
            </button>
           
        </div>
        </div>
    </div>

  )
}

export default Register