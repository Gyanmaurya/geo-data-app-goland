import React, { useState } from 'react';
import axios from 'axios';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/api/login', { username, password });
      localStorage.setItem('token', response.data.token); // Store JWT
      alert('Login successful!');
    } catch (error) {
      console.error('Login failed:', error);
      alert('Login failed: ' + error.response.data);
    }
  };

  return (
    <form onSubmit={handleLogin}>
      <input type="text" onChange={(e) => setUsername(e.target.value)} placeholder="Username" required />
      <input type="password" onChange={(e) => setPassword(e.target.value)} placeholder="Password" required />
      <button type="submit">Login</button>
    </form>
  );
};

export default Login;
