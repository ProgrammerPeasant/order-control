import React, {useEffect, useState} from "react";
import apiClient from "../../Utils/apiClient";
import Modal from "../../components/Modal";
import styles from "./ModalCompanyInfo.module.css";
import Button from "../../components/Button";
import Form from "../../components/Form";
import {handleErrorMessage} from "../../Utils/ErrorHandler";


const ModalCompanyInfo = ({companyId, isOpen, onClose, handleUpdate}) => {
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [mode, setMode] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            if (!isOpen) return;

            setMode(null);
            setLoading(true);
            setError(null);
            try {
                const response = await apiClient.get(`/api/v1/companies/${companyId}`, {headers: {Accept: "application/json"},});
                console.log(response.data);
                setData(response.data);
            } catch (error) {
                setError(handleErrorMessage(error));
            } finally {
                setLoading(false);
            }
        };

        if (isOpen && companyId) {
            fetchData();
        }
    }, [isOpen, companyId]);

    if (!isOpen) return null;

    if (loading) {
        return (
            <Modal title="Company Info" variant="type1" isOpen={isOpen} onClose={onClose}>
                <p className={styles.text}>Loading...</p>
            </Modal>
        )
    }

    if (error) {
        return (
            <Modal title="Company Info" variant="type1" isOpen={isOpen} onClose={onClose}>
                <p className={styles.text}>{error}</p>
            </Modal>
        )
    }

    const InfoRow = ({label, value}) => (
        <p className={styles.text}><strong>{label}:</strong> {value || "Not provided"}</p>
    );

    const fieldsUpdate = [
        {id: "name", type: "text", placeholder: "Name", value: data?.name},
        {id: "address", type: "text", placeholder: "Address", value: data?.address},
        {id: "desc", type: "text", placeholder: "Description", value: data?.desc, required: false},
        {id: "logo_url", type: "text", placeholder: "Logo Url", value: data?.logo_url, required: false},
        {id: "color_primary", type: "color", placeholder: "Primary color", value: data?.color_primary, required: false},
        {
            id: "color_secondary",
            type: "color",
            placeholder: "Secondary color",
            value: data?.color_secondary,
            required: false
        },
        {id: "color_accent", type: "color", placeholder: "Accent color", value: data?.color_accent, required: false},
    ]

    const fieldsDelete = [
        {id: "name", type: "text", placeholder: "Name"},
    ]

    const handleSubmitUpdate = async (e, formData) => {
        e.preventDefault();
        try {
            const response = await apiClient.put(`/api/v1/companies/${companyId}`, formData, {
                headers: {"Content-Type": "application/json", "Accept": "application/json"},
            })
            console.log(response.data);
            setMode(null)
            handleUpdate(companyId);
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    const handleSubmitDelete = async (e, formData) => {
        e.preventDefault();
        if (formData.name !== data.name) {
            alert("Names don't match");
            return;
        }
        try {
            const response = await apiClient.delete(`/api/v1/companies/${companyId}`, {
                headers: {"Accept": "application/json"},
            })
            console.log(response.data);
            setMode(null)
            onClose()
            handleUpdate(companyId);
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    if (mode === "update") {
        return (
            <Modal title="Edit Company Info" variant="type2" isOpen={isOpen} onClose={() => setMode(null)}>
                <Form fields={fieldsUpdate} handleSubmit={handleSubmitUpdate}/>
            </Modal>
        )
    }

    if (mode === "delete") {
        return (
            <Modal title="Delete Company" variant="type2" isOpen={isOpen} onClose={() => setMode(null)}>
                <p className={styles.textWarning}>Please enter company name for confirmation</p>
                <Form fields={fieldsDelete} handleSubmit={handleSubmitDelete}/>
            </Modal>
        )
    }

    return (
        <Modal title="Company Info" variant="type1" isOpen={isOpen} onClose={onClose}>
            <div className={styles.container}>
                <div>
                    <InfoRow label="ID" value={data?.ID}/>
                    <InfoRow label="Name" value={data?.name}/>
                    <InfoRow label="Description" value={data?.desc}/>
                    <InfoRow label="Address" value={data?.address}/>
                    <InfoRow label="Created At" value={new Date(data?.CreatedAt).toLocaleString()}/>
                    <InfoRow label="Updated At" value={new Date(data?.UpdatedAt).toLocaleString()}/>
                    <InfoRow label="Deleted At"
                             value={data?.DeletedAt ? new Date(data?.DeletedAt).toLocaleString() : "Not deleted"}/>
                </div>
                <div>
                    <img src={data?.logo_url || "/defaultpic.jpg"} alt="Company Logo"
                         style={{maxWidth: '50%', maxHeight: '50%'}}/>
                </div>
            </div>
            <Button title="Edit Company Info" variant="type2" onClick={() => setMode("update")}/>
            <Button title="Delete Company" variant="type4" onClick={() => setMode("delete")}/>
        </Modal>
    );
};

export default ModalCompanyInfo;