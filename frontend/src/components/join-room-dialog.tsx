import { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useJoinRoom } from '@/api/room';
import { useNavigate } from 'react-router-dom';
import { CustomAlert } from './custom-alert';

export function JoinRoomDialog() {
  const [open, setOpen] = useState(false);
  const [formData, setFormData] = useState({
    roomId: '',
    password: '',
  });
  const { joinRoom, loading, error, data } = useJoinRoom();
  const [showSuccessToast, setShowSuccessToast] = useState(false);
  const [showErrorToast, setShowErrorToast] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await joinRoom(formData);
  };

  useEffect(() => {
    if (data?.status === 'success' && data.room) {
      setShowSuccessToast(true);
      setTimeout(() => {
        navigate(`/room/${data.room?.id}`);
      }, 2000);
    }
  }, [data, navigate]);

  useEffect(() => {
    if (error) {
      setShowErrorToast(true);
    }
  }, [error]);

  return (
    <>
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogTrigger asChild>
          <Button variant="outline" size="lg" className="w-full sm:w-auto px-10 py-4 text-lg font-semibold border-indigo-600 text-indigo-600 hover:bg-indigo-600 hover:text-white transition-colors">
            Join a Room
          </Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Join Room</DialogTitle>
          </DialogHeader>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="roomId">Room ID</Label>
              <Input
                id="roomId"
                placeholder="Enter room ID"
                value={formData.roomId}
                onChange={(e) => setFormData({ ...formData, roomId: e.target.value })}
                required
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="password">Password (if private)</Label>
              <Input
                id="password"
                type="password"
                placeholder="Enter room password"
                value={formData.password}
                onChange={(e) => setFormData({ ...formData, password: e.target.value })}
              />
            </div>
            <div className="flex justify-end space-x-2">
              <Button type="button" variant="outline" onClick={() => setOpen(false)}>
                Cancel
              </Button>
              <Button type="submit" disabled={loading}>
                {loading ? 'Joining...' : 'Join Room'}
              </Button>
            </div>
          </form>
        </DialogContent>
      </Dialog>
      <CustomAlert
        message="Joined room successfully!"
        isVisible={showSuccessToast}
        onClose={() => setShowSuccessToast(false)}
        duration={3000}
        type="success"
      />
      <CustomAlert
        message={error || "Failed to join room"}
        isVisible={showErrorToast}
        onClose={() => setShowErrorToast(false)}
        duration={3000}
        type="error"
      />
    </>
  );
}