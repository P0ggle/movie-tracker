import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { signup } from "../services/api";
import "./SignupPage.css";

const SignupPage: React.FC = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const navigate = useNavigate();

  const handleSignup = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      await signup(username, password, email);
      navigate("/login");
    } catch (error) {
      console.error("Error signing up:", error);
      alert("Failed to sign up. Please try again.");
    }
  };

  return (
    <div className="signup-page">
      <Link to="/" className="button-style top-right-link">Home</Link>
      <div className="signup-form-container">
        <h1>Signup</h1>
        <form onSubmit={handleSignup}>
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
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <button type="submit" className="button-style">Signup</button>
        </form>
      </div>
    </div>
  );
};

export default SignupPage;

