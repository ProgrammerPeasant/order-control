import React from 'react';
import './DashboardPage.css';
import { useNavigate } from 'react-router-dom';
import { ReactComponent as DeleteIcon } from "../assets/deleteicon.svg";
import { ReactComponent as EditIcon } from "../assets/editicon.svg";

const estimates = [
  { id: 1, name: "Строитель", status: "В обработке", date: "21.02.22", notes: "комментарий" },
  { id: 2, name: "Машина", status: "Согласовано", date: "12.09.24", notes: "поставить сабик" },
];

const DashboardPage = () => {
	const navigate = useNavigate();

	const handleBack = () => {
		navigate('/mainmenu');
	};

  const handleCreate = () => {
    navigate("/createpage");
  };

  const handleUpload = () => {
    navigate('/');
  };

	return (
		<div className="dashboard-page">
			<div className="dashboard-page__card">
				<table className="dashboard-page__table">
				  <thead>
					<th>№</th>
					<th>Имя</th>
					<th>Статус</th>
					<th>Дата изменения</th>
					<th>Примечания</th>
					<th colspan="2"></th>
					</thead>
					<tbody>
            {estimates.map((item) => (
              <tr key={item.id}>
                <td>{item.id}</td>
                <td>{item.name}</td>
                <td>{item.status}</td>
                <td>{item.date}</td>
                <td>{item.notes}</td>
                <td><button className="dashboard-page__table-edit-button" onClick={() => alert("редактировать")}><EditIcon className="dashboard-page__table-button-icon"/></button></td>
                <td><button className="dashboard-page__table-delete-button" onClick={() => alert("удалииить")}><DeleteIcon className="dashboard-page__table-button-icon"/></button></td>
              </tr>
            ))}
            {[...Array(10)].map((_, i) => (
              <tr key={`empty-${i}`}>
                <td>&nbsp;</td>
                <td>&nbsp;</td>
                <td>&nbsp;</td>
                <td>&nbsp;</td>
                <td>&nbsp;</td>
                <td>&nbsp;</td>
                <td>&nbsp;</td>
              </tr>
            ))}
					</tbody>
				</table>
			</div>
			<div className="dashboard-page__button-group">
				<button className="dashboard-page__button dashboard-page__back-button" onClick={handleBack}> Назад </button>
				<button className="dashboard-page__button dashboard-page__create-button" onClick={handleCreate}>Создать</button>
				<button className="dashboard-page__button dashboard-page__upload-button" onClick={handleUpload}>Загрузить файл</button>
			</div>
		</div>
	);
};

export default DashboardPage;