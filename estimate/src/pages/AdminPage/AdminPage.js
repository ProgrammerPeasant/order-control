import React from "react";
import styles from "./AdminPage.module.css";
import CompanyTable from "../AdminPage/CompanyTable";
import Button from "../../components/Button";


const AdminPage = () => {
    const handleRegManager = () => {
        console.log("handleRegManager");
    }

    const handleRegCompany = () => {
        console.log("handleRegCompany");
    }

    const handleUpdateCompany = () => {
        console.log("handleUpdateCompany");
    }

    return (
        <div className={styles.page}>
            <div className={styles.container}>
                <CompanyTable />
            </div>
            <div className={styles.buttons}>
                <Button title="Reg Manager" variant="type2" onClick={handleRegManager} />
                <Button title="Reg Company" variant="type2" onClick={handleRegCompany} />
                <Button title="Update Company Info" variant="type2" onClick={handleUpdateCompany} />
            </div>
        </div>
    )
}

export default AdminPage;