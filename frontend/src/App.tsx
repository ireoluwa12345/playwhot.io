import "./App.css";
import Register from "@/pages/Register";
import Login from "@/pages/Login"
import Home from "@/pages/Home"
import { Routes, Route } from 'react-router-dom';

function App() {
  return (
    <>
      <Routes>
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/home" element={<Home />} />
        <Route path="/" element={<Home />} />
      </Routes>
    </>
  );
}

export default App;
