import React from 'react';
import Table from "../../components/Table";
import Button from "../../components/Button";


function CompanyTable() {
    const handleInfo = () => {
        console.log("info");
    }

    const handleCreate = () => {
        console.log("create");
    }

    const handleDelete = () => {
        console.log("delete");
    }

    const columns = ["Logo", "Name", "Type", "Level", "", "", ""]
    const apiUrl = "https://67ed9d094387d9117bbe2a38.mockapi.io/api/skkk/companies";

    const renderRow = (company) => (
        <tr key={company.id}>
            <td><img src={company.logo} alt="" style={{display: "flex", width: "50px", height: "50px"}}/></td>
            <td>{company.name}</td>
            <td>{company.type}</td>
            <td>{company.level}</td>
            <td><Button title="See info" variant="type3" onClick={handleInfo} /></td>
            <td><Button title="Create Estimate" variant="type2" onClick={handleCreate} /> </td>
            <td><Button title="Delete" variant="type4" onClick={handleDelete} /></td>
        </tr>
    );

    return <Table apiUrl={apiUrl} columns={columns} renderRow={renderRow} />;
}

export default CompanyTable;