import React from 'react';
import './LoginPage.css';
import { useNavigate } from 'react-router-dom';

const LoginPage = () => {
  const navigate = useNavigate();

  const handleBack = () => {
    navigate('/'); // Переход на StartPage
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Здесь можно обработать логин/пароль
    console.log('Логин отправлен!');
  };

  return (
    <div className="login-page">
      <div className="login-card">
        <h1 className="login-title">Вход</h1>
        <form className="login-form" onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="login">Логин</label>
            <input type="text" id="login" placeholder="Введите логин" />
          </div>
          <div className="form-group">
            <label htmlFor="password">Пароль</label>
            <input type="password" id="password" placeholder="Введите пароль" />
          </div>
          <div className="button-group">
            <button type="button" className="back-button" onClick={handleBack}>
              Назад
            </button>
            <button type="submit" className="submit-button">
              Продолжить
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default LoginPage;
