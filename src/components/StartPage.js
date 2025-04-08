import React from 'react';
import { useNavigate } from 'react-router-dom';
import './StartPage.css';

const StartPage = () => {
  const navigate = useNavigate();

  const goToRegister = () => {
    navigate('/register');
  };

  const goToLogin = () => {
    navigate('/login');
  };

  return (
    <div className="start-page">
      <div className="start-page__card">
        <h1 className="start-page__title">Название</h1>
        <p className="start-page__subtitle">Дополнительный текст</p>
      </div>
      <button className="start-page__button start-page__button--register" onClick={goToRegister}>Зарегистрироваться</button>
      <button className="start-page__button start-page__button--signin" onClick={goToLogin}>Войти</button>
    </div>
  );
};

export default StartPage;
