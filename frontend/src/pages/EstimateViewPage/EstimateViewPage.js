import styles from "./EstimateViewPage.module.css"
import ButtonPanel from "./ButtonPanel";
import {useParams} from "react-router-dom";
import React, {useCallback, useContext, useEffect, useState} from "react";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import DataTable from "./DataTable";
import {AuthContext} from "../../Utils/AuthProvider";
import EditableDataTable from "./EditableDataTable";


const EstimateViewPage = () => {
    const { estimateId } = useParams();
    const [data, setData] = useState([]);
    const [companyInfo, setCompanyInfo] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const {user} = useContext(AuthContext)

    const InfoRow = ({ label, value }) => (
        <p className={styles.text}><strong>{label}:</strong> {value || "Not provided"}</p>
    );

    const fetchData = useCallback(async () => {
        setLoading(true);
        setError(null);
        try {
            const response = await apiClient.get(`/api/v1/estimates/${estimateId}`, {headers: {Accept : "application/json"}});
            console.log(response.data);
            setData(response.data);
        } catch (error) {
            setError(handleErrorMessage(error));
        } finally {
            setLoading(false);
        }
    }, [estimateId]);

    useEffect(() => {
        const fetchInfo = async () => {
            setLoading(true);
            setError(null);
            try {
                const response = await apiClient.get(`/api/v1/companies/${user.companyId}`, {headers: {Accept : "application/json"}});
                console.log(response.data);
                setCompanyInfo(response.data);
            } catch (error) {
                if (user.role !== "ADMIN") {
                    setError(handleErrorMessage(error));
                }
            } finally {
                setLoading(false);
            }
        };

        fetchData();
        fetchInfo();
    }, [fetchData, user.companyId, user.role]);

    if (loading) {
        return (
            <div className={styles.container}>
                {loading && <p className={styles.text}>Loading...</p>}
            </div>
        )
    }

    if (error) {
        return (
            <div className={styles.container}>
                <p className={styles.text}>{error}</p>
            </div>
        )
    }

    return (
        <div className={styles.container}>
            <div className={styles.header}>
                <div className={styles.info}>
                    <InfoRow label="Title" value={data?.title} />
                    <InfoRow label="Estimate ID" value={data?.ID} />
                    <InfoRow label="Total" value={data?.total_amount?.toLocaleString()} />
                    <InfoRow label="Overall discount" value={data?.overall_discount_percent?.toLocaleString()} />
                </div>
                <div className={styles.manager}>
                    <InfoRow label="Manager" value={data?.created_by_id} />
                </div>
                <div className={styles.logo}>
                    <img src={companyInfo?.logo_url || "/defaultpic.jpg"} alt="Company Logo" style={{ maxWidth: '80%', maxHeight: '80%', borderRadius: '10px' }} />
                </div>
            </div>
            <div className={styles.body}>
                {user.role === "USER" && <DataTable data={data}>
                    <ButtonPanel estimateId={estimateId} data={data} fetchData={fetchData} />
                </DataTable>}
                {user.role !== "USER" && <EditableDataTable data={data} setData={setData} fetchData={fetchData}>
                    <ButtonPanel estimateId={estimateId} data={data} fetchData={fetchData} />
                </EditableDataTable>}
            </div>
        </div>
    )
}

export default EstimateViewPage;