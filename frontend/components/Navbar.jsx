import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => (
  <nav style={{ padding: '1rem', background: '#222', color: '#fff' }}>
    <Link to="/" style={{ marginRight: '1rem', color: '#fff', textDecoration: 'none' }}>Home</Link>
    <Link to="/auth/signup" style={{ marginRight: '1rem', color: '#fff', textDecoration: 'none' }}>Sign Up</Link>
    <Link to="/auth/signin" style={{ color: '#fff', textDecoration: 'none' }}>Sign In</Link>
  </nav>
);

export default Navbar;
