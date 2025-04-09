import React, { useState, useEffect } from "react";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import styles from "../../components/Table.module.css"


const DataTable = ({estimateId}) => {
    const [items, setItems] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            setLoading(true);
            setError(null);
            try {
                const response = await apiClient.get(`/api/v1/estimates/${estimateId}`, {headers: {Accept : "application/json"}});
                setItems(response.data.items);
            } catch (error) {
                if (error.response) {
                    const {status, data} = error.response;
                    setError({status, message: data.message || "An error occurred"});
                } else {
                    setError({status: null, message: error.message || "Network error"});
                }
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, [estimateId]);

    if (loading) {
        return <p className={styles.text}>Loading...</p>;
    }

    if (error) {
        const message = handleErrorMessage(error);
        return <p className={styles.text}>{message}</p>;
    }

    return (
        <table className={styles.table}>
            <thead>
            <tr>
                <th>ID</th>
                <th>Product Name</th>
                <th>Quantity</th>
                <th>Unit Price</th>
                <th>Total Price</th>
                <th>Discount (%)</th>
                <th>Created At</th>
                <th>Updated At</th>
            </tr>
            </thead>
            <tbody>
            {items.length > 0 ? (
                items.map((item) => (
                    <tr key={item.ID}>
                        <td>{item.ID}</td>
                        <td>{item.product_name}</td>
                        <td>{item.quantity}</td>
                        <td>{item.unit_price}</td>
                        <td>{item.total_price}</td>
                        <td>{item.discount_percent}</td>
                        <td>{new Date(item.CreatedAt).toLocaleString()}</td>
                        <td>{new Date(item.UpdatedAt).toLocaleString()}</td>
                    </tr>
                ))
            ) : (
                <tr>
                    <td colSpan="8">No data</td>
                </tr>
            )}
            </tbody>
        </table>
    );
};

export default DataTable;