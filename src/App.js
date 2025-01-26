import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import StartPage from './components/StartPage';
import RegisterPage from './components/RegisterPage';
import LoginPage from './components/LoginPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<StartPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
      </Routes>
    </Router>
  );
}

export default App;
