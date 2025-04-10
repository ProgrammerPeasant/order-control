import React, {useEffect, useState} from "react";
import styles from "./EditableDataTable.module.css"
import styles2 from "../../components/Table.module.css"


const DataTable = ({data, children, color_secondary, color_accent}) => {
    const [items, setItems] = useState([]);

    useEffect(() => {
        if (data && data.items) {
            setItems(data.items);
        }
    }, [data]);

    return (
        <div className={styles.container}>
            <div className={styles.scrollable}>
                <table className={styles2.table}>
                    <thead style={{backgroundColor: color_secondary}}>
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
                            <td colSpan={8}>No Data</td>
                        </tr>
                    )}
                    </tbody>
                </table>
            </div>
            <div className={styles.panel} style={{backgroundColor: color_accent}}>
                {children}
            </div>
        </div>
    )
};

export default DataTable;