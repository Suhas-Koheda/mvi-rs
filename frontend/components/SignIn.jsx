import React, { useState } from 'react';

const SignIn = () => {
  const [form, setForm] = useState({ Email: '', Password: '' });
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async e => {
    e.preventDefault();
    setError('');
    setSuccess('');
    try {
      const res = await fetch('/auth/signin', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(form)
      });
      if (!res.ok) throw new Error('Sign in failed');
      setSuccess('Sign in successful!');
    } catch (err) {
      setError('Sign in error');
    }
  };

  return (
    <div style={{ maxWidth: 400, margin: '2rem auto' }}>
      <h2>Sign In</h2>
      <form onSubmit={handleSubmit}>
        <input name="Email" type="email" placeholder="Email" value={form.Email} onChange={handleChange} required style={{ width: '100%', marginBottom: 8 }} />
        <input name="Password" type="password" placeholder="Password" value={form.Password} onChange={handleChange} required style={{ width: '100%', marginBottom: 8 }} />
        <button type="submit" style={{ width: '100%' }}>Sign In</button>
        {error && <div style={{ color: 'red', marginTop: 8 }}>{error}</div>}
        {success && <div style={{ color: 'green', marginTop: 8 }}>{success}</div>}
      </form>
    </div>
  );
};

export default SignIn;
