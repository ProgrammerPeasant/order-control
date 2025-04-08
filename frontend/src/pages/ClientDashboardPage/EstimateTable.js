import React from "react";
import Table from "../../components/Table";
import Button from "../../components/Button";


function EstimateTable({estimateId}) {
    const columns = ["ID", "Title", "Total", "Created at", "Created by ID", ""];
    const apiUrl = estimateId ? `/api/v1/estimates/${estimateId}` : "/api/v1/estimates/my";

    const handleView = (estimateId) => { // ATTENTION
        console.log(estimateId);
    }

    const renderRow = (item) => (
        <tr key={item.ID}>
            <td>{item.ID}</td>
            <td>{item.title}</td>
            <td>{item?.total_amount}</td>
            <td>{new Date(item?.CreatedAt).toLocaleString() || "N/A"}</td>
            <td>{item?.created_by_id || "N/A"}</td>
            <td><Button title="View" variant="type3" onClick={() => handleView(item.ID)} /></td>
        </tr>
    );

    return <Table apiUrl={apiUrl} columns={columns} renderRow={renderRow} />;
}

export default EstimateTable;