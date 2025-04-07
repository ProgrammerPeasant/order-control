import React, {useState, useEffect, useContext} from "react";
import apiClient from "../../Utils/apiClient";
import Modal from "../../components/Modal";
import styles from "./ModalCompanyInfo.module.css";
import Button from "../../components/Button";
import Form from "../../components/Form";
import {AuthContext} from "../../Utils/AuthProvider";


const ModalCompanyInfo = ({companyId, isOpen, onClose}) => {
    const {user} = useContext(AuthContext);
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [isCreateMode, setIsCreateMode] = useState(false);

    useEffect(() => {
        const fetchData = async () => {
            if (!isOpen) return;

            setIsCreateMode(false);
            setLoading(true);
            setError(null);
            try {
                const response = await apiClient.get(`/api/v1/companies/${companyId}`, {headers: { Accept: "application/json" },});
                console.log(response.data);
                setData(response.data);
            } catch (error) {
                if (error.response) {
                    const { status, data } = error.response;
                    setError({ status, message: data.message || "An error occurred" });
                } else {
                    setError({ status: null, message: error.message || "Network error" });
                }
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
        let message;
        if (error.status === 400) {
            message = "Wrong format";
        } else if (error.status === 401) {
            message = "Unauthorized";
        } else if (error.status === 404) {
            message = "Company does not exist";
        } else if (error.status === 500) {
            message = "Internal Server Error";
        }
        return (
            <Modal title="Company Info" variant="type1" isOpen={isOpen} onClose={onClose}>
                <p className={styles.text}>{message}</p>
            </Modal>
        )
    }

    const InfoRow = ({ label, value }) => (
        <p className={styles.text}><strong>{label}:</strong> {value || "Not provided"}</p>
    );

    const fields = [
        {id: "title", type: "text", placeholder: "Title", required: true},
        {id: "overall_discount_percent", type: "number", placeholder: "Overall discount", required: true},
    ]

    const handleSubmit = async (e, formData) => {
        e.preventDefault();

        const updatedData = {
            ...formData,
            overall_discount_percent: parseFloat(formData.overall_discount_percent),
            createdAt: new Date(), // ATTENTION
            updatedAt: new Date(), // ATTENTION
            deletedAt: null, // ATTENTION
            company_id: parseInt(companyId, 10), // ATTENTION
            created_by_id: user.id, // ATTENTION
        }

        try {
            const response = await apiClient.post("/api/v1/estimates", updatedData, {
                headers: { "Content-Type": "application/json", "Accept": "application/json" },
            });

            console.log(response.data);
            setIsCreateMode(false);
        } catch (error) {
            if (error.response) {
                const {status, data} = error.response;
                if (status === 400) {
                    alert(`Invalid data: ${data.message || "Check your input fields"}`);
                } else if (status === 401) {
                    alert(`Unauthorized: ${data.message}`);
                } else if (status === 403) {
                    alert(`Access denied: ${data.message || "Admin and manager only"}`);
                } else if (status === 500) {
                    alert(`Server error: ${data.message || "Please try again later"}`);
                } else {
                    alert(`Error: ${data.message || "Something went wrong"}`);
                }
            } else if (error.request) {
                alert("No response from server. Please check your internet connection.");
            } else {
                alert(`Request error: ${error.message}`);
            }
        }
    }

    const handleCreateEstimate = () => {
        console.log("Create Estimate");
        setIsCreateMode(true);
    }

    if (isCreateMode) {
        return (
            <Modal title="Company Info" variant="type2" isOpen={isOpen} onClose={onClose}>
                <Form fields={fields} handleSubmit={handleSubmit} />
            </Modal>
        )
    }

    return (
        <Modal title="Company Info" variant="type1" isOpen={isOpen} onClose={onClose}>
            <div>
                <InfoRow label="ID" value={data?.ID} />
                <InfoRow label="Name" value={data?.name} />
                <InfoRow label="Description" value={data?.desc} />
                <InfoRow label="Address" value={data?.address} />
                <InfoRow label="Created At" value={new Date(data?.CreatedAt).toLocaleString()} />
                <InfoRow label="Updated At" value={new Date(data?.UpdatedAt).toLocaleString()} />
                <InfoRow label="Deleted At" value={data?.DeletedAt ?new Date(data?.DeletedAt).toLocaleString() : "Not deleted"} />
            </div>
            <Button title="Create Estimate" variant="type2" onClick={handleCreateEstimate} />
            <Button title="Edit Company Info" variant="type2" onClick={() => console.log("")} />
            <Button title="Delete Company" variant="type4" onClick={() => console.log("")} />
        </Modal>
    );
};

export default ModalCompanyInfo;