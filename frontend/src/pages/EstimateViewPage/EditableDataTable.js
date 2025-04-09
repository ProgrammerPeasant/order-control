import React, {useEffect, useState} from "react";
import apiClient from "../../Utils/apiClient";
import styles from "./EditableDataTable.module.css"
import styles2 from "../../components/Table.module.css"
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import Button from "../../components/Button";


const EditableDataTable = ({data, setData, fetchData, children}) => {
    const [items, setItems] = useState([]);

    useEffect(() => {
        if (data && data.items) {
            setItems(data.items);
        }
    }, [data]);

    const handleCellChange = (id, field, value) => {
        setItems((prevItems) =>
            prevItems.map((item) =>
                item.ID === id ? { ...item, [field]: value } : item
            )
        );
    };

    const handleAdd = () => {
        const newItem = {
            ID: Date.now() % 100000,
            product_name: "",
            quantity: 0,
            unit_price: 0,
            total_price: 0,
            discount_percent: 0,
            CreatedAt: new Date().toISOString(),
            UpdatedAt: new Date().toISOString(),
            estimate_id: data.ID,
            DeletedAt: null
        };
        setItems((prevItems) => [...prevItems, newItem]);
    };

    const handleSave = async () => {
        try {
            const updatedData = {
                ...data,
                items: items,
            };
            setData(updatedData);
            const response = await apiClient.put(`/api/v1/estimates/${data.ID}`, updatedData, {headers: {"Content-Type": "application/json", "Accept": "application/json"}});
            console.log(response.data);
            fetchData();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    };

    return (
        <div className={styles.container}>
            <div className={styles.scrollable}>
            <table className={styles2.table}>
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
                            <td><input className={styles.input} type="text" value={item.product_name} onChange={(e) => handleCellChange(item.ID, "product_name", e.target.value)}/></td>
                            <td><input className={styles.input} type="number" value={item.quantity} onChange={(e) => handleCellChange(item.ID, "quantity", +e.target.value)}/></td>
                            <td><input className={styles.input} type="number" value={item.unit_price} onChange={(e) => handleCellChange(item.ID, "unit_price", +e.target.value)}/></td>
                            <td>{item.total_price}</td>
                            <td><input className={styles.input} type="number" value={item.discount_percent} onChange={(e) => handleCellChange(item.ID, "discount_percent", +e.target.value)}/></td>
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
            </div>
            <div className={styles.panel}>
                <div className={styles.buttons}>
                    <Button title="Add" variant="type2" onClick={handleAdd} />
                    <Button title="Save changes" variant="type3" onClick={handleSave} />
                </div>
                {children}
            </div>
        </div>
    );
};

export default EditableDataTable;