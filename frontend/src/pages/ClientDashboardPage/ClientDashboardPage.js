import React, {useState} from "react";
import styles from "./ClientDashboardPage.module.css"
import EstimateTable from "./EstimateTable";
import Button from "../../components/Button";
import {useNavigate} from "react-router-dom";
import useDebounce from "../../Utils/useDebounce";


const ClientDashboardPage = () => {
    const [estimateId, setEstimateId] = useState("");
    const debouncedValue = useDebounce(estimateId, 500);

    const navigate = useNavigate();

    const handleSettings = () => {
        console.log("settings");
        navigate("/settings");
    }

    return (
        <div className={styles.page}>
            <div className={styles.container}>
                {<EstimateTable estimateId={debouncedValue} />}
            </div>
            <div className={styles.buttons}>
                <Button title="Settings" variant="type3" onClick={handleSettings}/>
                <input className={styles.input} id="estimate_id" type="text" value={estimateId} onChange={(e) => setEstimateId(e.target.value)} placeholder="Estimate ID"/>
            </div>
        </div>
    )
}

export default ClientDashboardPage;