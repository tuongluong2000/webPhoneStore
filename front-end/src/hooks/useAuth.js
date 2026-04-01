import { loginApi } from "../api/authApi";

export default function useAuth() {
  const login = async (email, password) => {
    const data = await loginApi(email, password);
    // backend feedback token
    localStorage.setItem("token", data.token);

    return data;
  };

  return { login };
}