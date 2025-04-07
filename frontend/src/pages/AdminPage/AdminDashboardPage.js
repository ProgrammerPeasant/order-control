import React, {useState} from "react";
import styles from "./AdminDashboardPage.module.css";
import CompanyTable from "../AdminPage/CompanyTable";
import Button from "../../components/Button";
import {useNavigate} from "react-router-dom";
import useDebounce from "../../Utils/useDebounce";
import ModalCompanyInfo from "./ModalCompanyInfo";
import ModalCreateCompany from "./ModalCreateCompany";


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

    const handleGetCompanyInfo = () => {
        console.log("handleGetCompanyInfo");
        openModal("modalCompanyInfo")
    }

    const handleCreateCompany = () => {
        console.log("handleCreateCompany");
        openModal("modalCreateCompany");
    }

    return (
        <div className={styles.page}>
            <div className={styles.container}>
                <CompanyTable companyId={debouncedValue} />
            </div>
            <div className={styles.buttons}>
                <Button title="Settings" variant="type3" onClick={handleSettings}/>
                <input className={styles.input} id="company_id" type="text" value={companyId} onChange={(e) => setCompanyId(e.target.value)} placeholder="Company ID" />
                <Button title="Get Company Info" variant="type2" onClick={handleGetCompanyInfo} />
                <Button title="Register Manager" variant="type2" onClick={() => console.log("")} />
                <Button title="Create Company" variant="type2" onClick={handleCreateCompany} />
            </div>
            <ModalCompanyInfo isOpen={activeModal === "modalCompanyInfo"} onClose={closeModal} companyId={debouncedValue} />
            <ModalCreateCompany isOpen={activeModal === "modalCreateCompany"} onClose={closeModal} />
        </div>
    )
}

export default AdminDashboardPage;