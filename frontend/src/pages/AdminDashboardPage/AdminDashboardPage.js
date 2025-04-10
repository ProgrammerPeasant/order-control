import React, {useState} from "react";
import styles from "./AdminDashboardPage.module.css";
import CompanyTable from ".//CompanyTable";
import Button from "../../components/Button";
import {useNavigate} from "react-router-dom";
import useDebounce from "../../Utils/useDebounce";
import ModalCompanyInfo from "./ModalCompanyInfo";
import ModalCreateCompany from "./ModalCreateCompany";
import ModalRegisterUser from "./ModalRegisterUser";


const AdminDashboardPage = () => {
    const [activeModal, setActiveModal] = useState(null);
    const openModal = (modalId) => setActiveModal(modalId);
    const closeModal = () => setActiveModal(null);

    const [companyId, setCompanyId] = useState("");
    const delay = 100;
    const debouncedValue = useDebounce(companyId, delay);

    const navigate = useNavigate();

    const handleUpdate = (companyId) => {
        const prevId = companyId
        setCompanyId("        ");
        setTimeout(() => {
            setCompanyId(prevId)
        }, delay * 2)
    }

    return (
        <div className={styles.page}>
            <div className={styles.container}>
                <CompanyTable companyId={debouncedValue} handleUpdate={handleUpdate}/>
            </div>
            <div className={styles.buttons}>
                <Button title="Settings" variant="type3" onClick={() => navigate("/settings")}/>
                <input className={styles.input} id="company_id" type="text" value={companyId}
                       onChange={(e) => setCompanyId(e.target.value)} placeholder="Company ID"/>
                <Button title="Get Company Info" variant="type2" onClick={() => openModal("modalCompanyInfo")}/>
                <Button title="Register User" variant="type2" onClick={() => openModal("registerUser")}/>
                <Button title="Create Company" variant="type2" onClick={() => openModal("modalCreateCompany")}/>
                <Button title="Refresh Table" variant="type1" onClick={() => handleUpdate(companyId)}/>
            </div>
            <ModalCompanyInfo isOpen={activeModal === "modalCompanyInfo"} onClose={closeModal}
                              companyId={debouncedValue} handleUpdate={handleUpdate}/>
            <ModalRegisterUser isOpen={activeModal === "registerUser"} onClose={closeModal}/>
            <ModalCreateCompany isOpen={activeModal === "modalCreateCompany"} onClose={closeModal}/>
        </div>
    )
}

export default AdminDashboardPage;