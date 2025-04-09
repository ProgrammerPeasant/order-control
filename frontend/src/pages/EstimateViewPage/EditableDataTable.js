import React, { useState, useEffect } from "react";
import apiClient from "../../Utils/apiClient";
import styles from "./EditableDataTable.module.css"


const EditableDataTable = ({ estimateId }) => {
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

    const handleCellChange = (id, field, value) => {
        setItems((prevItems) =>
            prevItems.map((item) =>
                item.ID === id ? { ...item, [field]: value } : item
            )
        );
    };

    const addNewItem = () => {
        const newItem = {
            ID: 22222, // ATTENTION
            product_name: "",
            quantity: 0,
            unit_price: 0,
            total_price: 0,
            discount_percent: 0,
            CreatedAt: new Date().toISOString(),
            UpdatedAt: new Date().toISOString(),
        };
        setItems((prevItems) => [...prevItems, newItem]);
    };

    // Сохранение изменений на сервере
    const saveChanges = async () => {
        try {
            await apiClient.put(`/api/v1/estimates/${estimateId}`, { items }); // ATTENTION
            alert("Changes saved successfully!");
        } catch (error) {
            console.error("Failed to save changes:", error);
            alert("Failed to save changes. Please try again.");
        }
    };

    return (
        <div>
            <h2>Estimate Table</h2>
            <button onClick={addNewItem} className={styles.addButton}>
                Add New Item
            </button>
            <button onClick={saveChanges} className={styles.saveButton}>
                Save Changes
            </button>
            {loading && <p>Loading...</p>}
            {error && <p>Error: {error.message}</p>}
            {!loading && !error && (
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
                                <td>
                                    <input
                                        type="text"
                                        value={item.product_name}
                                        onChange={(e) =>
                                            handleCellChange(item.ID, "product_name", e.target.value)
                                        }
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        value={item.quantity}
                                        onChange={(e) =>
                                            handleCellChange(item.ID, "quantity", +e.target.value)
                                        }
                                    />
                                </td>
                                <td>
                                    <input
                                        type="number"
                                        value={item.unit_price}
                                        onChange={(e) =>
                                            handleCellChange(item.ID, "unit_price", +e.target.value)
                                        }
                                    />
                                </td>
                                <td>{item.total_price}</td>
                                <td>
                                    <input
                                        type="number"
                                        value={item.discount_percent}
                                        onChange={(e) =>
                                            handleCellChange(item.ID, "discount_percent", +e.target.value)
                                        }
                                    />
                                </td>
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
            )}
        </div>
    );
};

export default EditableDataTable;