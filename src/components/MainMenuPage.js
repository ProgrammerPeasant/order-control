import React from 'react';
import './MainMenuPage.css';
import { useNavigate } from 'react-router-dom';

const MainMenuPage = () => {
  const navigate = useNavigate();

  const handleCreate = () => {
    navigate('/');
  };

  const handleOpen = () => {
    navigate('/');
  };

  const handleUpload = () => {
    navigate('/');
  };

  const handleExit = () => {
    navigate('/');
  };

  return (
    <div className="main-menu-page">
      <div className="main-menu-card">
        <div className="button-group">
          <button type="button" className="create-button" onClick={handleCreate}>
            Создать
          </button>
          <button type="button" className="open-button" onClick={handleOpen}>
            Открыть
          </button>
          <button type="button" className="upload-button" onClick={handleUpload}>
            Загрузить файл
          </button>
          <button type="button" className="exit-button" onClick={handleExit}>
            Выйти
          </button>
        </div>
      </div>
    </div>
  );
};

export default MainMenuPage;


