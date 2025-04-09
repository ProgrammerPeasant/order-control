import styles from "./ManagerDashboardPage.module.css"
import Button from "../../components/Button";
import React, {useState} from "react";
import Table from "../../components/Table";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import Modal from "../../components/Modal";
import {useNavigate} from "react-router-dom";


function EstimateTable({estimateId, handleUpdate}) {
    const [activeModal, setActiveModal] = useState(null);
    const [selectedEstimateId, setSelectedEstimateId] = useState(null);
    const openModal = (modalId, estimateId) => {
        setSelectedEstimateId(estimateId);
        setActiveModal(modalId);
    }
    const closeModal = () => {
        setSelectedEstimateId(null);
        setActiveModal(null);
    }
    const navigate = useNavigate();

    const columns = ["ID", "Title", "Total", "Created at", "Created by ID", "", ""]
    const apiUrl = estimateId ? `/api/v1/estimates/${estimateId}` : "/api/v1/estimates/my";

    if (estimateId === "        ") {
        return <p className={styles.text}>Refreshing...</p>;
    }

    const handleView = (estimateId) => {
        navigate(`/estimateview/${estimateId}`);
    }

    const handleDelete = async () => {
        if (!selectedEstimateId) return;
        try {
            const response = await apiClient.delete(`/api/v1/estimates/${selectedEstimateId}`, {headers: {Accept: "application/json"}});
            console.log(response.data);
            handleUpdate(estimateId);
            closeModal();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    const renderRow = (item) => (
        <tr key={item.ID}>
            <td>{item.ID}</td>
            <td>{item.title}</td>
            <td>{item?.total_amount}</td>
            <td>{new Date(item?.CreatedAt).toLocaleString() || "N/A"}</td>
            <td>{item?.created_by_id || "N/A"}</td>
            <td><Button title="View" variant="type3" onClick={() => handleView(item.ID)}/></td>
            <td><Button title="Delete" variant="type4" onClick={() => openModal("modalDeleteEstimate", item.ID)}/></td>
        </tr>
    );

    return (
        <div>
            <Table apiUrl={apiUrl} columns={columns} renderRow={renderRow} />
            <Modal title={"Delete estimate?"} variant="type1" isOpen={activeModal === "modalDeleteEstimate"} onClose={closeModal}>
                <Button title="Delete" variant="type4" onClick={() => handleDelete()}/>
            </Modal>
        </div>
    )
}

export default EstimateTable;