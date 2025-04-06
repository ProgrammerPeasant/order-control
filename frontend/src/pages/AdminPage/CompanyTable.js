import React from 'react';
import Table from "../../components/Table";
import Button from "../../components/Button";
import styles from "./AdminDashboardPage.module.css"


function CompanyTable({companyId}) {
    const handleCreate = () => {
        console.log("create"); // ATTENTION
    }

    const handleView = () => {
        console.log("view"); // ATTENTION
    }

    const handleEdit = () => {
        console.log("edit"); // ATTENTION
    }

    const handleDelete = () => {
        console.log("delete"); // ATTENTION
    }

    const columns = ["ID", "Title", "Total", "Created at", "Created by", "", "", ""];
    const apiUrl = companyId ? `/api/v1/estimates/company?company_id=${companyId}` : null;

    if (apiUrl === null) {
        return <p className={styles.text}>Type Company ID below</p>;
    }

    const renderRow = (item) => (
        <tr key={item.ID}>
            <td>{item.ID}</td>
            <td>{item.title}</td>
            <td>{item?.total_amount}</td>
            <td>{new Date(item?.CreatedAt).toLocaleString()}</td>
            <td>{item?.CreatedBy?.Username || "N/A"}</td>
            <td><Button title="View" variant="type3" onClick={handleView} /></td>
            <td><Button title="Edit" variant="type2" onClick={handleEdit} /> </td>
            <td><Button title="Delete" variant="type4" onClick={handleDelete} /></td>
        </tr>
    );

    return <Table apiUrl={apiUrl} columns={columns} renderRow={renderRow} />;
}

export default CompanyTable;