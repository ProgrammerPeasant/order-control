import React from "react";
import {useNavigate} from "react-router-dom";
import styles from "./CompaniesPage.module.css";
import Button from "../../components/Button";
import CompanyTable from "./CompanyTable";


const CompaniesPage = () => {
    const navigate = useNavigate();

    const handleBack = () => {
        console.log("back");
        navigate("/clientdashboard");
    }

    return (
        <div className={styles.page}>
            <div className={styles.container}>
                <CompanyTable />
            </div>
            <div className={styles.buttons}>
                <Button title="back" variant="type3" onClick={handleBack} />
            </div>
        </div>
    )
}

export default CompaniesPage