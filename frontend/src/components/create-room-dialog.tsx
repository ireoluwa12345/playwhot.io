import { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useCreateRoom } from '@/api/room';
import { useNavigate } from 'react-router-dom';
import { CustomAlert } from './custom-alert';

export function CreateRoomDialog() {
  const [open, setOpen] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    maxPlayers: 4,
    isPrivate: false,
    password: '',
  });
  const { createRoom, loading, error, data } = useCreateRoom();
  const [showSuccessToast, setShowSuccessToast] = useState(false);
  const [showErrorToast, setShowErrorToast] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await createRoom(formData);
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
            Create Room
          </Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Create New Room</DialogTitle>
          </DialogHeader>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="roomName">Room Name</Label>
              <Input
                id="roomName"
                placeholder="Enter room name"
                value={formData.name}
                onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                required
              />
            </div>
            <div className="space-y-2">
              <Label htmlFor="maxPlayers">Max Players</Label>
              <Input
                id="maxPlayers"
                type="number"
                min="2"
                max="8"
                value={formData.maxPlayers}
                onChange={(e) => setFormData({ ...formData, maxPlayers: parseInt(e.target.value) })}
                required
              />
            </div>
            <div className="space-y-2">
              <Label>
                <input
                  type="checkbox"
                  checked={formData.isPrivate}
                  onChange={(e) => setFormData({ ...formData, isPrivate: e.target.checked })}
                  className="mr-2"
                />
                Private Room
              </Label>
            </div>
            {formData.isPrivate && (
              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <Input
                  id="password"
                  type="password"
                  placeholder="Enter room password"
                  value={formData.password}
                  onChange={(e) => setFormData({ ...formData, password: e.target.value })}
                  required={formData.isPrivate}
                />
              </div>
            )}
            <div className="flex justify-end space-x-2">
              <Button type="button" variant="outline" onClick={() => setOpen(false)}>
                Cancel
              </Button>
              <Button type="submit" disabled={loading}>
                {loading ? 'Creating...' : 'Create Room'}
              </Button>
            </div>
          </form>
        </DialogContent>
      </Dialog>
      <CustomAlert
        message="Room created successfully!"
        isVisible={showSuccessToast}
        onClose={() => setShowSuccessToast(false)}
        duration={3000}
        type="success"
      />
      <CustomAlert
        message={error || "Failed to create room"}
        isVisible={showErrorToast}
        onClose={() => setShowErrorToast(false)}
        duration={3000}
        type="error"
      />
    </>
  );
}