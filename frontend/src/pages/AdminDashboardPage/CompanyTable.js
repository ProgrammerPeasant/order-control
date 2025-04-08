import React, {useEffect, useState} from 'react';
import Table from "../../components/Table";
import Button from "../../components/Button";
import styles from "./AdminDashboardPage.module.css"
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import Modal from "../../components/Modal";


function CompanyTable({companyId, handleUpdate}) {
    const [activeModal, setActiveModal] = useState(null);
    const [selectedEstimateId, setSelectedEstimateId] = useState(null);
    const openModal = (modalId, estimateId) => {
        setSelectedEstimateId(estimateId)
        setActiveModal(modalId);
    }
    const closeModal = () => {
        setSelectedEstimateId(null);
        setActiveModal(null);
    }
    const columns = ["Company status", "ID", "Title", "Total", "Created at", "Created by ID", "", ""];
    const apiUrl = companyId ? `/api/v1/estimates/company?company_id=${companyId}` : null;
    const [status, setStatus] = useState("Loading...");

    useEffect(() => {
        const fetchData = async () => {
            try {
                await apiClient.get(`/api/v1/companies/${companyId}`, {
                    headers: { Accept: "application/json" },
                })
                setStatus("Active");
            } catch (error) {
                if (error.response && error.response.status === 404) {
                    setStatus("Deleted");
                } else {
                    setStatus("Error");
                }
            }
        }
        fetchData();
    }, [companyId]);

    if (apiUrl === null) {
        return <p className={styles.text}>Type Company ID below</p>;
    }

    if (companyId === "        ") {
        return <p className={styles.text}>Refreshing...</p>;
    }

    const handleView = (estimateId) => { // ATTENTION
        console.log(estimateId);
    }

    const handleDelete = async () => {
        if (!selectedEstimateId) return;
        try {
            const response = await apiClient.delete(`/api/v1/estimates/${selectedEstimateId}`, {headers: { Accept: "application/json" }});
            console.log(response.data);
            handleUpdate(companyId);
            closeModal();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    const renderRow = (item) => (
        <tr key={item.ID}>
            <td>{status}</td>
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
            <Modal title="Delete estimate?" variant="type1" isOpen={activeModal === "modalDeleteEstimate"} onClose={closeModal}>
                <Button title="Delete" variant="type4" onClick={() => handleDelete()} />
            </Modal>
        </div>
    );
}

export default CompanyTable;