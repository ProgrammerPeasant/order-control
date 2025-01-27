import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import StartPage from './components/StartPage';
import RegisterPage from './components/RegisterPage';
import LoginPage from './components/LoginPage';
import MainMenuPage from './components/MainMenuPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<StartPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/mainmenu" element={<MainMenuPage />} />
      </Routes>
    </Router>
  );
}

export default App;
