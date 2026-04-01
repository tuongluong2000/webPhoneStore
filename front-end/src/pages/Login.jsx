import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "../styles/login.css";
import useAuth from "../hooks/useAuth";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const {login} = useAuth()

   const handleLogin = async (e) => {
    e.preventDefault(); // block reload page

    try {

      await login(email, password)
      navigate("/home"); // redirect after login
    } catch (err) {
      console.log(err)
      alert("Sai email hoặc mật khẩu");
    }
  };

  return (
    <div className="login-container">
      <div className="login-box">
        <h2>Welcome Back 👋</h2>
        <p>Login to continue</p>

        <form onSubmit={handleLogin}>
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />

          <button type="submit">Login</button>
        </form>

        <span className="hint">
          Demo: admin@gmail.com / 123456
        </span>
      </div>
    </div>
  );
}