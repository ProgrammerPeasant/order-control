import React, {useState} from "react";
import styles from "./ClientDashboardPage.module.css"
import EstimateTable from "./EstimateTable";
import Button from "../../components/Button";
import {useNavigate} from "react-router-dom";
import useDebounce from "../../useDebounce";


const ClientDashboardPage = () => {
    const [companyId, setCompanyId] = useState("");
    const debouncedValue = useDebounce(companyId, 500);

    const navigate = useNavigate();

    const handleSettings = () => {
        console.log("settings");
        navigate("/settings");
    }

    const handleCompanies = () => {
        console.log("companies");
        navigate("/companies");
    }

    const handleUpload = () => {
        console.log("upload");
    }

    return (
        <div className={styles.page}>
            <div className={styles.container}>
                {!debouncedValue && <p className={styles.text}>Enter company ID below</p> || debouncedValue && <EstimateTable companyId={debouncedValue} />}
            </div>
            <div className={styles.buttons}>
                <Button title="Settings" variant="type3" onClick={handleSettings}/>
                <input className={styles.input} id="company_id" type="text" value={companyId} onChange={(e) => setCompanyId(e.target.value)} placeholder="Company ID"/>
                <Button title="Companies" variant="type1" onClick={handleCompanies}/>
                <Button title="Upload" variant="type2" onClick={handleUpload}/>
            </div>
        </div>

    )
}

export default ClientDashboardPage;