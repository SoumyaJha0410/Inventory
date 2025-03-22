import React from 'react';
import { BrowserRouter, Route, Routes } from "react-router-dom";

import './App.css';
import { Root } from './components/Root.tsx';
import { SignUp } from './components/SignUp.tsx';
import { Login } from './components/Login.tsx';
import Dashboard from './components/Dasshboard.tsx';
import AddProduct from './components/AddProduct.tsx';
import {UpdateProduct} from './components/UpdateProduct.tsx';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Root />} />
        <Route path='/login' element={<Login />} />
        <Route path='/signup' element={<SignUp />} />
        <Route path='/dashboard' element={<Dashboard />} />
        <Route path='/addProduct' element={<AddProduct />} />
        <Route path='/updateProduct' element={<UpdateProduct />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
