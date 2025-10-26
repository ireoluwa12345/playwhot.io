import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from '@/components/ui/button';

const Navbar: React.FC = () => {
  return (
    <nav className="w-full">
      <div className="w-full px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="flex items-center">
            <Link to="/" className="flex items-center">
              <img src="/logo.jpg" alt="Logo" className="h-8 w-8 mr-2" />
              <span className="text-xl font-bold text-white-800">PlayWhot</span>
            </Link>
          </div>
          <div className="flex space-x-4">
            <Link to="/">
              <Button variant="ghost" className="text-white-800 hover:text-white-600">
                Home
              </Button>
            </Link>
            <Link to="/login">
              <Button variant="ghost" className="text-white-800 hover:text-white-600">
                Login
              </Button>
            </Link>
            <Link to="/register">
              <Button variant="ghost" className="text-white-800 hover:text-white-600">
                Register
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;