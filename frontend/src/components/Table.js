import React, {useState, useEffect} from "react";
import apiClient from "../Utils/apiClient";
import styles from "./Table.module.css"


function Table({apiUrl, columns, renderRow, emptyRows = 10}) {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            setLoading(true);
            setError(null);
            try {
                const response = await apiClient.get(apiUrl);
                console.log(response.data);
                setData(response.data);
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
    }, [apiUrl, data.message]);

    if (loading) {
        return <p className={styles.text}>Loading...</p>;
    }

    if (!((Array.isArray(data) && data.length > 0) || (!Array.isArray(data) && data.ID))) {
        return <p className={styles.text}>No data</p>
    }

    if (error) {
        let message;
        if (error.status === 400) {
            message = "Wrong format";
        } else if (error.status === 401) {
            message = "Unauthorized";
        } else if (error.status === 404) {
            message = "Not Found";
        } else if (error.status === 500) {
            message = "Internal Server Error";
        }
        return <p className={styles.text}>{message}</p>;
    }

    return (
        <table className={styles.table}>
            <thead>
            <tr>{columns.map((column, index) => (<th key={index}>{column}</th>))}</tr>
            </thead>
            <tbody>
            {Array.isArray(data) && data.length > 0
                ? data.map((item) => renderRow(item))
                : !Array.isArray(data) && data ? renderRow(data)
                    : null}

                {[...Array(emptyRows)].map((_, i) => (
                    <tr key={`empty-${i}`}>{columns.map((_, index) => (<td key={index}>&nbsp;</td>))}</tr>
                ))}
            </tbody>
        </table>
    )
}

export default Table;