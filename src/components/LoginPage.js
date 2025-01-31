import React from 'react';
import './LoginPage.css';
import { useNavigate } from 'react-router-dom';

const LoginPage = () => {
  const navigate = useNavigate();

  const handleBack = () => {
    navigate('/');
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    
    // const formData = new FormData(e.target);
    // const data = Object.fromEntries(formData);

    navigate('/mainmenu');
    // Здесь можно обработать логин/пароль
    console.log('Логин отправлен!');
  };

  // const form = document.querySelector('form');
  // form.addEventListener('submit', handleSubmit);

  return (
    <div className="login-page">
      <div className="login-page__card">
        <h1 className="login-page__title">Вход</h1>
        <form className="login-page__form" onSubmit={handleSubmit}>
          <div className="login-page__form-group">
            <label htmlFor="login">Логин</label>
            <input type="text" id="login" placeholder="Введите логин" />
          </div>
          <div className="login-page__form-group">
            <label htmlFor="password">Пароль</label>
            <input type="password" id="password" placeholder="Введите пароль" />
          </div>
          <div className="login-page__button-group">
            <button type="button" className="login-page__button login-page__back-button" onClick={handleBack}>Назад</button>
            <button type="submit" className="login-page__button login-page__continue-button">Продолжить</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default LoginPage;
