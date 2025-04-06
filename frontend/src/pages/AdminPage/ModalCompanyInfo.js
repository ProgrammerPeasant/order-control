import React, { useState, useEffect } from "react";
import apiClient from "../../apiClient";
import Modal from "../../components/Modal";
import styles from "./ModalCompanyInfo.module.css";


const ModalCompanyInfo = ({companyId, isOpen, onClose}) => {
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            if (!isOpen) return;

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

    return (
        <Modal title="Company Info" isOpen={isOpen} onClose={onClose}>
            {loading && <p className={styles.text}>Loading...</p>}
            {error && (
                <div>
                    {error.status === 404 ? (
                        <p className={styles.text}>Company does not exist</p>
                    ) : (
                        <p className={styles.text}>{error.message}</p>
                    )}
                </div>
            )}
            {data && !error && (
                <div className={styles.companyInfo}>
                    {/* ID */}
                    <p className={styles.text}>
                        <strong>ID:</strong> {data.ID || "Not provided"}
                    </p>

                    {/* Name */}
                    <p className={styles.text}>
                        <strong>Name:</strong> {data.name || "Not provided"}
                    </p>

                    {/* Description */}
                    <p className={styles.text}>
                        <strong>Description:</strong> {data.desc || "Not provided"}
                    </p>

                    {/* Address */}
                    <p className={styles.text}>
                        <strong>Address:</strong> {data.address || "Not provided"}
                    </p>

                    {/* Created At */}
                    <p className={styles.text}>
                        <strong>Created At:</strong>{" "}
                        {data.CreatedAt
                            ? new Date(data.CreatedAt).toLocaleString() // Форматируем дату
                            : "Not provided"}
                    </p>

                    {/* Updated At */}
                    <p className={styles.text}>
                        <strong>Updated At:</strong>{" "}
                        {data.UpdatedAt
                            ? new Date(data.UpdatedAt).toLocaleString() // Форматируем дату
                            : "Not provided"}
                    </p>

                    {/* Deleted At */}
                    <p className={styles.text}>
                        <strong>Deleted At:</strong>{" "}
                        {data.DeletedAt
                            ? new Date(data.DeletedAt).toLocaleString() // Форматируем дату
                            : "Not deleted"}
                    </p>
                </div>
            )}
        </Modal>
    );
};

export default ModalCompanyInfo;