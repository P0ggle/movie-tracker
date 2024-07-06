import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { login } from "../services/api";
import "./LoginPage.css";

const LoginPage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      await login(username, password);
      navigate("/");
    } catch (error) {
      console.error("Error logging in:", error);
      alert("Failed to log in. Please check your credentials.");
    }
  };

  return (
    <div className="login-page">
      <Link to="/" className="button-style top-right-link">Home</Link>
      <div className="login-form-container">
        <h1>Login</h1>
        <form onSubmit={handleLogin}>
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <button type="submit" className="button-style">Login</button>
        </form>
      </div>
    </div>
  );
};

export default LoginPage;

