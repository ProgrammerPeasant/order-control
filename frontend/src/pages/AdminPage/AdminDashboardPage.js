import React, {useState} from "react";
import styles from "./AdminDashboardPage.module.css";
import CompanyTable from "../AdminPage/CompanyTable";
import Button from "../../components/Button";
import {useNavigate} from "react-router-dom";
import useDebounce from "../../useDebounce";
import ModalCompanyInfo from "./ModalCompanyInfo";


const AdminDashboardPage = () => {
    const [activeModal, setActiveModal] = useState(null);
    const openModal = (modalId) => setActiveModal(modalId);
    const closeModal = () => setActiveModal(null);

    const [companyId, setCompanyId] = useState("");
    const debouncedValue = useDebounce(companyId, 500);

    const navigate = useNavigate();

    const handleSettings = () => {
        console.log("settings");
        navigate("/settings");
    }

    const handleCreateManager = () => {
        console.log("handleRegManager"); // ATTENTION
    }

    const handleCreateCompany = () => {
        console.log("handleRegCompany"); // ATTENTION
    }

    // const handleUpdateCompanyInfo = () => {
    //     console.log("handleUpdateCompany"); // ATTENTION
    // }

    const handleGetCompanyInfo = () => {
        console.log("handleGetCompanyInfo");
        openModal("modalCompanyInfo")
    }

    return (
        <div className={styles.page}>
            <div className={styles.container}>
                <CompanyTable companyId={debouncedValue} />
            </div>
            <div className={styles.buttons}>
                <Button title="Settings" variant="type3" onClick={handleSettings}/>
                <input className={styles.input} id="company_id" type="text" value={companyId} onChange={(e) => setCompanyId(e.target.value)} placeholder="Company ID" />
                <Button title="Register Manager" variant="type2" onClick={handleCreateManager} />
                <Button title="Create Company" variant="type2" onClick={handleCreateCompany} />
                <Button title="Get Company Info" variant="type2" onClick={handleGetCompanyInfo} />
            </div>
            <ModalCompanyInfo isOpen={activeModal === "modalCompanyInfo"} onClose={closeModal} companyId={debouncedValue}></ModalCompanyInfo>
        </div>
    )
}

export default AdminDashboardPage;