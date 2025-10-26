import { useState } from 'react';

interface RegisterData {
  email: string;
  username: string;
  password: string;
  confirmPassword: string;
}

interface RegisterResponse {
  status: string;
  message?: string;
}

export function useRegister() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<RegisterResponse | null>(null);

  const register = async (registerData: RegisterData) => {
    if (registerData.password !== registerData.confirmPassword) {
      setError('Passwords do not match');
      return;
    }

    const apiUrl = import.meta.env.VITE_API_URL

    setLoading(true);
    setError(null);

    try {
      const response = await fetch(`${apiUrl}/api/auth/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: registerData.email,
          username: registerData.username,
          password: registerData.password,
        }),
      });

      const result: RegisterResponse = await response.json();

      if (!response.ok) {
        throw new Error(result.message || 'Registration failed');
      }

      setData(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setLoading(false);
    }
  };

  return { register, loading, error, data };
}