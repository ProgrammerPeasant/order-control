import React, {useState, useEffect} from "react";
import apiClient from "../apiClient";
import styles from "./Table.module.css"


function Table({apiUrl, columns, renderRow, emptyRows = 10}) {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await apiClient.get(apiUrl);
                console.log(response.data);
                setData(response.data);
            } catch (error) {
                return <p className={styles.text}>Error: {error.message}</p>;
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, [apiUrl]);

    if (loading) {
        return <p className={styles.text}>Loading...</p>;
    }

    if (!Array.isArray(data)) {
        return <p className={styles.text}>No Data</p>;
    }

    return (
        <table className={styles.table}>
            <thead>
            <tr>{columns.map((column, index) => (<th key={index}>{column}</th>))}</tr>
            </thead>
            <tbody>
            {data.map((item) => renderRow(item))}
                {[...Array(emptyRows)].map((_, i) => (
                    <tr key={`empty-${i}`}>{columns.map((_, index) => (<td key={index}>&nbsp;</td>))}</tr>
                ))}
            </tbody>
        </table>
    )
}

export default Table;