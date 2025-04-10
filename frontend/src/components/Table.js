import React, {useEffect, useState} from "react";
import apiClient from "../Utils/apiClient";
import styles from "./Table.module.css"
import {handleErrorMessage} from "../Utils/ErrorHandler";


function Table({apiUrl, columns, renderRow, emptyRows = 7}) {
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
                setError(handleErrorMessage(error));
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
        return <p className={styles.text}>No Data</p>
    }

    if (error) {
        return <p className={styles.text}>{error}</p>;
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