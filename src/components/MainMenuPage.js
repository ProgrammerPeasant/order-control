import React from 'react';
import './MainMenuPage.css';
import { useNavigate } from 'react-router-dom';

const MainMenuPage = () => {
  const navigate = useNavigate();

  const handleToDash = () => {
    navigate('/dashboard');
  };

  const handleOptions = () => {
    navigate('/');
  };

  const handleExit = () => {
    navigate('/');
  };

  return (
    <div className="main-menu-page">
      <div className="main-menu-page__card">
        <div className="main-menu-page__button-group">
          <button className="main-menu-page__button main-menu-page__open-upload-button" onClick={handleToDash}>Мои сметы</button>
          <button className="main-menu-page__button main-menu-page__options-button" onClick={handleOptions}>Настройки</button>
          <button className="main-menu-page__button main-menu-page__exit-button" onClick={handleExit}>Выйти</button>
        </div>
      </div>
    </div>
  );
};

export default MainMenuPage;


