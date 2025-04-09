import Button from "../../components/Button";
import styles from "./EstimateViewPage.module.css"
import {useContext, useState} from "react";
import {useNavigate} from "react-router-dom";
import {AuthContext} from "../../Utils/AuthProvider";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import Modal from "../../components/Modal";
import Form from "../../components/Form";


const ButtonPanel = ({estimateId, data, fetchData}) => {
    const navigate = useNavigate();
    const {user} = useContext(AuthContext)

    const [activeModal, setActiveModal] = useState(null);
    const openModal = (modalId) => setActiveModal(modalId);
    const closeModal = () => setActiveModal(null);

    const fieldsUpdate = [
        {id: "title", type: "text", placeholder: "Title", value: data?.title},
        {id: "overall_discount_percent", type: "number", placeholder: "Overall discount", value: data?.overall_discount_percent},
    ]

    const handleExport = async () => {
        try {
            const response = await apiClient.get(`/api/v1/estimates/${estimateId}/export/excel`, {headers: {Accept: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}, responseType: "blob"});
            const url = window.URL.createObjectURL(new Blob([response.data]));
            const link = document.createElement('a');
            link.href = url;
            const contentDisposition = response.headers['content-disposition'];
            let fileName = 'estimate.xlsx';
            if (contentDisposition && contentDisposition.includes('filename=')) {
                fileName = contentDisposition.split('filename=')[1].split(';')[0].replace(/"/g, '');
            }
            link.setAttribute('download', fileName);
            document.body.appendChild(link);
            link.click();
            link.remove();
            window.URL.revokeObjectURL(url);
        } catch (error) {
            alert(handleErrorMessage(error))
        }
    }

    const handleSave = async () => {
        try {
            const response = await apiClient.put(`/api/v1/estimates/${estimateId}`, data, {headers: {"Content-Type": "application/json", "Accept": "application/json"}});
            console.log(response.data);
            fetchData();
            closeModal();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    const handleSubmitUpdate = async (e, formData) => {
        e.preventDefault();
        const updatedData = {
            ...data,
            ...formData,
            overall_discount_percent: parseFloat(formData.overall_discount_percent),
        }
        try {
            const response = await apiClient.put(`/api/v1/estimates/${estimateId}`, updatedData, {headers: {"Content-Type": "application/json", "Accept": "application/json"}});
            console.log(response.data);
            fetchData();
            closeModal();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    const handleExit = async () => {
        if (user.role === "ADMIN") {
            await handleSave();
            navigate("/admin");
        } else if (user.role === "MANAGER") {
            await handleSave();
            navigate("/managerdashboard");
        } else if (user.role === "USER") {
            navigate("/clientdashboard");
        }
    }

    return (
        <div className={styles.buttons}>
            <Button title="Export xlsx" variant="type2" onClick={handleExport}/>
            {user.role !== "USER" && <Button title="Save changes" variant="type2" onClick={handleSave}/>}
            {user.role !== "USER" && <Button title="Edit info" variant="type2" onClick={() => openModal("modalEditInfo")}/>}
            <Button title="Exit" variant="type3" onClick={handleExit}/>
            <Modal title="Edit Estimate Info" variant={"type1"} isOpen={activeModal === "modalEditInfo"} onClose={closeModal}>
                <Form fields={fieldsUpdate} handleSubmit={handleSubmitUpdate} variant="type2"/>
            </Modal>
        </div>
    )
}

export default ButtonPanel