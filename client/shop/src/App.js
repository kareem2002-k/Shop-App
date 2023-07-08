import logo from './logo.svg';
import './App.css';
import { Route, Routes } from 'react-router-dom';
import Login from './Pages/Login';
import Register from './Pages/Register';
import Home from './Pages/Home';
import { BrowserRouter } from 'react-router-dom';




function App() {




    


  return (
    <BrowserRouter>
       <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/Home" element={<Home />} />

        </Routes>
    </BrowserRouter>
   
  );
}

export default App;
