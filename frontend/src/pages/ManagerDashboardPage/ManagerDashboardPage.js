import styles from "./ManagerDashboardPage.module.css"
import EstimateTable from "./EstimateTable";
import Button from "../../components/Button";
import React, {useState} from "react";
import {useNavigate} from "react-router-dom";
import useDebounce from "../../Utils/useDebounce";
import ModalCreateEstimate from "./ModalCreateEstimate";
import ModalJoinRequests from "./ModalJoinRequests";


const ManagerDashboardPage = () => {
    const [activeModal, setActiveModal] = useState(null);
    const openModal = (modalId) => setActiveModal(modalId);
    const closeModal = () => setActiveModal(null);

    const [estimateId, setEstimateId] = useState("");
    const delay = 100;
    const debouncedValue = useDebounce(estimateId, delay);

    const navigate = useNavigate();

    const handleUpdate = (estimateId) => {
        const prevId = estimateId
        setEstimateId("        ");
        setTimeout(() => {setEstimateId(prevId)}, delay * 2)
    }

    return (
        <div className={styles.page}>
            <div className={styles.container}>
                <EstimateTable estimateId={debouncedValue} handleUpdate={handleUpdate} />
            </div>
            <div className={styles.buttons}>
                <Button title="Settings" variant="type3" onClick={() => navigate("/settings")}/>
                <input className={styles.input} id="estimate_id" type="text" value={estimateId} onChange={(e) => setEstimateId(e.target.value)} placeholder="Estimate ID" />
                <Button title="Create Estimate" variant="type2" onClick={() => openModal("modalCreateEstimate")} />
                <Button title="Show Requests" variant="type2" onClick={() => openModal("modalJoinRequests")} />
            </div>
            <ModalCreateEstimate isOpen={activeModal === "modalCreateEstimate"} onClose={closeModal} estimateId={debouncedValue} handleUpdate={handleUpdate}/>
            <ModalJoinRequests isOpen={activeModal === "modalJoinRequests"} onClose={closeModal} />
        </div>
    )
}

export default ManagerDashboardPage;