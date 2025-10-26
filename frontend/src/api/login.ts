import { useState } from "react";

interface LoginData {
  username: string;
  password: string;
}

interface User {
  id: string;
  email: string;
  username: string;
}

interface LoginResponse {
  status: string;
  message?: string;
  user?: User;
}

export function UseLogin() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<LoginResponse | null>(null);

  const login = async (LoginData: LoginData) => {

    const apiUrl = import.meta.env.VITE_API_URL

    setLoading(true);
    setError(null);

    try {
      const response = await fetch(`${apiUrl}/api/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          username: LoginData.username,
          password: LoginData.password,
        }),
      });

      const result: LoginResponse = await response.json();

      if (!response.ok) {
        throw new Error(result.message || 'Authentication failed');
      }

      setData(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setLoading(false);
    }
  };

  return { login, loading, error, data };
}
