import React from "react";
import Table from "../../components/Table";
import Button from "../../components/Button";


function EstimateTable({companyId}) {
    const handleView = () => {
        console.log("view")
    }

    const handleDelete = () => {
        console.log("delete")
    }

    const columns = ["ID", "Title", "Total", "Created at", "Created by", "", ""];
    const apiUrl = "/api/v1/estimates/company?company_id=" + companyId;

    const renderRow = (item) => (
        <tr key={item.ID}>
            <td>{item.ID}</td>
            <td>{item.title}</td>
            <td>{item?.total_amount}</td>
            <td>{new Date(item?.CreatedAt).toLocaleString() || "N/A"}</td>
            <td>{item?.CreatedBy?.Username || "N/A"}</td>
            <td><Button title="View" variant="type3" onClick={handleView} /></td>
            <td><Button title="Delete" variant="type4" onClick={handleDelete} /></td>
        </tr>
    );

    return <Table apiUrl={apiUrl} columns={columns} renderRow={renderRow} />;
}

export default EstimateTable;