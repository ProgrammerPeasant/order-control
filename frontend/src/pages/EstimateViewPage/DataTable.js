import React from "react";
import styles from "../../components/Table.module.css"


const DataTable = ({items}) => {
    if (!items) {
        return <p className={styles.text}>No data</p> ;
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