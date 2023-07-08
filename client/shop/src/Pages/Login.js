import React, { useEffect } from 'react'
import { useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom';
 
function Login() {

    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const navigate = useNavigate();

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

    



    const loginHandler =  () => {

        axios.post('http://localhost:8000/login', {
            "email": email,
            "password": password
            
        }, {
            withCredentials: true
        }).then((response) => {
            console.log(response)
            navigate('/Home')
        })
    }
    

        



  return (
    <div className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 flex flex-col">
    <div className="mb-4">
      <label className="block text-grey-darker text-sm font-bold mb-2" for="username">
        Username
      </label>
      <input 
        onChange={
            (e) => {
                setEmail(e.target.value)
            }
        }
      className="shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker" id="username" type="text" placeholder="Username"/>
    </div>
    <div className="mb-6">
      <label className="block text-grey-darker text-sm font-bold mb-2" for="password"  >
        Password
      </label>
      <input onChange={
            (e) => {
                setPassword(e.target.value)
            }

      }  className="shadow appearance-none border border-red rounded w-full py-2 px-3 text-grey-darker mb-3" id="password" type="password" placeholder="******************"/>
      <p className="text-red text-xs italic">Please choose a password.</p>
    </div>
    <div className="flex items-center justify-between">
      <button onClick={loginHandler} className=" bg-blue-500 hover:bg-blue-dark text-white font-bold py-2 px-4 rounded" type="button">
        Sign In
      </button>
      <a className="inline-block align-baseline font-bold text-sm text-blue hover:text-blue-darker" href="#">
        Forgot Password?
      </a>
    </div>

    <button className="bg-blue-500 hover:bg-blue-dark text-white font-bold py-2 px-4 rounded" type="button" onClick={
        () => {
            axios.get('http://localhost:8000/user', {
                withCredentials: true
            }).then((response) => {
                console.log(response)
            })

        }
    }>
        Get User
        </button>
</div>
  )
}

export default Login