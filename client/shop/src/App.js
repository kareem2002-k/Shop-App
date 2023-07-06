import logo from './logo.svg';
import './App.css';
import { useState } from 'react';
import axios from 'axios';




function App() {

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName,setLastName] = useState("");


  const handleLogin = (e) => {
    e.preventDefault();
    axios.post("http://localhost:8000/login", {
      email: email,
      password: password
    }, {
      withCredentials: true
    })
      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      })
  }

  const handleRegister = (e) => {
    e.preventDefault();
    axios.post("http://localhost:8000/register", {
      name: firstName+" "+lastName,
      email: email,
      password: password
    }, {
      withCredentials: true
    })

      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      })
  }


    


  return (
    <div className="App">
      <h1>Log in</h1>
      <form>
        <label>
          Email:
          <input type="text" name="email" 
          onChange={(e)=> setEmail(e.target.value)} />
        </label>
        <label>
          Password:
          <input type="password" name="password"
            onChange={(e)=> setPassword(e.target.value)}
          
          />
        </label>
        <input type="submit" value="Log in" />
        <button onClick={handleLogin}>Log in</button>
      </form>

      <h1>Register</h1>
      <form>
        <label>
          First Name:
          <input className=' text-lg  ' type="text" name="firstName"
          onChange={(e)=> setFirstName(e.target.value)}
          
          />
        </label>
        <label>
          Last Name:
          <input type="text" name="lastName"
          onChange={(e)=> setLastName(e.target.value)}
          
          />
        </label>
        <label>
          Email:
          <input type="text" name="email"
          onChange={(e)=> setEmail(e.target.value)}
          />
        </label>
        <label>
          Password:
          <input type="password" name="password"
          onChange={(e)=> setPassword(e.target.value)}
          
          />
        </label>

        <button onClick={handleRegister}>Register</button>
        </form>


    </div>
  );
}

export default App;
