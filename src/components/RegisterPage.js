import React from 'react';
import './RegisterPage.css';
import { useNavigate } from 'react-router-dom';

const RegisterPage = () => {
  const navigate = useNavigate();

  const handleBack = () => {
    navigate('/');
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Здесь можно обработать логин/пароль
    console.log('Логин отправлен!');
  };

  return (
    <div className="register-page">
      <div className="register-page__card">
        <h1 className="register-page__title">Регистрация</h1>
        <form className="register-page__form" onSubmit={handleSubmit}>
          <div className="register-page__form-group">
            <label htmlFor="login">Логин</label>
            <input type="text" id="login" placeholder="Введите логин" />
          </div>
          <div className="register-page__form-group">
            <label htmlFor="password">Пароль</label>
            <input type="password" id="password" placeholder="Введите пароль" />
          </div>
          <div className="register-page__form-group">
            <label htmlFor="confirm-password">Пароль</label>
            <input type="password" id="confirm-password" placeholder="Повторите пароль" />
          </div>
          <div className="register-page__button-group">
            <button type="button" className="register-page__button register-page__back-button" onClick={handleBack}>Назад</button>
            <button type="submit" className="register-page__button register-page__continue-button">Продолжить</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default RegisterPage;
