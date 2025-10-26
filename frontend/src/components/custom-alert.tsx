import { useEffect } from 'react';
import { capitalizeFirstWords } from '@/lib/utils';

interface CustomAlertProps {
  message: string;
  isVisible: boolean;
  onClose: () => void;
  duration?: number;
  type?: 'success' | 'error' | 'info';
}

export function CustomAlert({
  message,
  isVisible,
  onClose,
  duration = 3000,
  type = 'success'
}: CustomAlertProps) {
  useEffect(() => {
    if (isVisible) {
      const timer = setTimeout(() => {
        onClose();
      }, duration);
      return () => clearTimeout(timer);
    }
  }, [isVisible, onClose, duration]);

  const typeStyles = {
    success: 'bg-green-500 text-white',
    error: 'bg-red-500 text-white',
    info: 'bg-blue-500 text-white'
  };

  return (
    <div className={`toast ${isVisible ? 'show' : 'hide'} px-4 py-2 rounded-lg shadow-lg ${typeStyles[type]}`}>
      {capitalizeFirstWords(message)}
    </div>
  );
}