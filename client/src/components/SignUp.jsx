import React, { useState } from 'react';
import axios from 'axios';

const Signup = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSignup = async (e) => {
    e.preventDefault();
    try {
      await axios.post('http://localhost:8080/api/register', { username, password });
      alert('Signup successful!');
    } catch (error) {
      console.error('Signup failed:', error);
      alert('Signup failed: ' + error.response.data);
    }
  };

  return (
    <form onSubmit={handleSignup}>
      <input type="text" onChange={(e) => setUsername(e.target.value)} placeholder="Username" required />
      <input type="password" onChange={(e) => setPassword(e.target.value)} placeholder="Password" required />
      <button type="submit">Signup</button>
    </form>
  );
};

export default Signup;
