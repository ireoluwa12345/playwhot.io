import { useState } from 'react';

interface RoomData {
  name: string;
  maxPlayers: number;
  isPrivate: boolean;
  password?: string;
}

interface Room {
  id: string;
  name: string;
  hostId: string;
  players: string[];
  maxPlayers: number;
  isPrivate: boolean;
  createdAt: string;
}

interface RoomResponse {
  status: string;
  message?: string;
  room?: Room;
}

interface JoinRoomData {
  roomId: string;
  password?: string;
}

export function useCreateRoom() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<RoomResponse | null>(null);

  const createRoom = async (roomData: RoomData) => {
    const apiUrl = import.meta.env.VITE_API_URL;

    setLoading(true);
    setError(null);

    try {
      const response = await fetch(`${apiUrl}/api/room`, {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(roomData),
      });

      const result: RoomResponse = await response.json();

      if (!response.ok) {
        throw new Error(result.message || 'Failed to create room');
      }

      setData(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setLoading(false);
    }
  };

  return { createRoom, loading, error, data };
}

export function useJoinRoom() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<RoomResponse | null>(null);

  const joinRoom = async (joinData: JoinRoomData) => {
    const apiUrl = import.meta.env.VITE_API_URL;

    setLoading(true);
    setError(null);

    try {
      const response = await fetch(`${apiUrl}/api/rooms/${joinData.roomId}/join`, {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ password: joinData.password }),
      });

      const result: RoomResponse = await response.json();

      if (!response.ok) {
        throw new Error(result.message || 'Failed to join room');
      }

      setData(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setLoading(false);
    }
  };

  return { joinRoom, loading, error, data };
}

export function useGetRooms() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [data, setData] = useState<Room[] | null>(null);

  const getRooms = async () => {
    const apiUrl = import.meta.env.VITE_API_URL;

    setLoading(true);
    setError(null);

    try {
      const response = await fetch(`${apiUrl}/api/rooms`, {
        method: 'GET',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      const result: Room[] = await response.json();

      if (!response.ok) {
        throw new Error('Failed to fetch rooms');
      }

      setData(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setLoading(false);
    }
  };

  return { getRooms, loading, error, data };
}