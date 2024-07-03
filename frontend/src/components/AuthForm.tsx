import React, { useState } from 'react';

interface AuthFormProps {
  onSubmit: (email: string, password: string, copyPassword?: string) => void;
  buttonText: string;
}

const AuthForm: React.FC<AuthFormProps> = ({ onSubmit, buttonText }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [copyPassword, setCopyPassword] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (buttonText === "Register") {
      onSubmit(email, password, copyPassword);
    } else {
      onSubmit(email, password);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Email</label>
        <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
      </div>
      <div>
        <label>Password</label>
        <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
      </div>
      {buttonText === "Register" && (
        <div>
          <label>Confirm Password</label>
          <input type="password" value={copyPassword} onChange={(e) => setCopyPassword(e.target.value)} required />
        </div>
      )}
      <button type="submit">{buttonText}</button>
    </form>
  );
};

export default AuthForm;
