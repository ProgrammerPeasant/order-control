import React from 'react';
import './CreatePage.css';
import { useNavigate } from 'react-router-dom'

const CreatePage = () => {
	const navigate = useNavigate();

	const handleBack = () => {
		navigate('/dashboard')
	};

	const handleSubmit = (e) => {
		e.preventDefault();
		navigate('/')
		// обработка создания новой сметы
	};

	return (
		<div className="create-page">
			<div className="create-page__card">
				<h1 className="create-page__title">Введите название новой сметы</h1>
				<form className="create-page__form" onSubmit={handleSubmit}>
					<div className="create-page__form-group">
						<label>Название</label>
						<input placeholder="Введите название" />
					</div>
					<div className="create-page__button-group">
						<button type="button" className="create-page__button create-page__back-button" onClick={handleBack}>Назад</button>
						<button type="submit" className="create-page__button create-page__submit-button">Продолжить</button>
					</div>
				</form>
			</div>
		</div>
	);
};

export default CreatePage;