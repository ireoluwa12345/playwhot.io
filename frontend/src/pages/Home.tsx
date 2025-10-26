import React from 'react';
import { Button } from '@/components/ui/button';
import Navbar from '@/components/Navbar';

const Home: React.FC = () => {
  const handleCreateRoom = () => {
    // Placeholder for creating a room
    console.log('Create Room clicked');
  };

  const handleJoinRoom = () => {
    // Placeholder for joining a room
    console.log('Join Room clicked');
  };

  return (
    <div className="min-h-screen bg-gradient-to-br w-full">
      <Navbar />
      <section className="flex items-center justify-center min-h-[calc(100vh-4rem)] px-4 w-full">
        <div className="text-center w-full max-w-4xl mx-auto">
          <h1 className="text-4xl md:text-6xl font-bold mb-6 leading-tight">
            Welcome to <span className="text-indigo-600">PlayWhot</span>
          </h1>
          <p className="text-lg md:text-xl mb-10 max-w-2xl mx-auto">
            Play the classic Whot card game online with friends. Create or join a room to start the fun!
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
            <Button onClick={handleCreateRoom} variant="outline" size="lg" className="w-full sm:w-auto px-10 py-4 text-lg font-semibold border-indigo-600 text-indigo-600 hover:bg-indigo-600 hover:text-white transition-colors">
              Create Room
            </Button>
            <Button onClick={handleJoinRoom} variant="outline" size="lg" className="w-full sm:w-auto px-10 py-4 text-lg font-semibold border-indigo-600 text-indigo-600 hover:bg-indigo-600 hover:text-white transition-colors">
              Join a Room
            </Button>
          </div>
        </div>
      </section>
    </div>
  );
};

export default Home;