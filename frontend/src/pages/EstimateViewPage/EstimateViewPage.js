import styles from "./EstimateViewPage.module.css"
import ButtonPanel from "./ButtonPanel";
import {useParams} from "react-router-dom";
import {useCallback, useEffect, useState} from "react";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";


const EstimateViewPage = () => {
    const { estimateId } = useParams();
    const [data, setData] = useState([]);

    const InfoRow = ({ label, value }) => (
        <p className={styles.text}><strong>{label}:</strong> {value || "Not provided"}</p>
    );

    const fetchData = useCallback(async () => {
        try {
            const response = await apiClient.get(`/api/v1/estimates/${estimateId}`, {headers: {Accept : "application/json"}});
            console.log(response.data);
            setData(response.data);
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }, [estimateId]);

    useEffect(() => {
        fetchData();
    }, [fetchData]);

    return (
        <div className={styles.container}>
            <div className={styles.header}>
                <div className={styles.info}>
                    <InfoRow label="Title" value={data?.title} />
                    <InfoRow label="Estimate ID" value={data?.ID} />
                    <InfoRow label="Total" value={data?.total_amount.toLocaleString()} />
                    <InfoRow label="Overall discount" value={data?.overall_discount_percent.toLocaleString()} />
                </div>
                <div className={styles.manager}>
                    <InfoRow label="Manager" value={data?.created_by_id} /> {/* ATTENTION manager contact*/}
                </div>
                <div className={styles.logo}>{/* ATTENTION company logo */}</div>
            </div>
            <div className={styles.body}>
                <div className={styles.table}>{/* ATTENTION table */}</div>
                <div className={styles.panel}>
                    <ButtonPanel estimateId={estimateId} data={data} fetchData={fetchData} />
                </div>
            </div>
        </div>
    )
}

export default EstimateViewPage;