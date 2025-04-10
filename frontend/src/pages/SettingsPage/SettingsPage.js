import React, {useContext, useState} from "react";
import styles from "./SettingsPage.module.css"
import Button from "../../components/Button";
import {useNavigate} from "react-router-dom";
import Modal from "../../components/Modal";
import {AuthContext} from "../../Utils/AuthProvider";


const SettingsPage = () => {
    const [activeModal, setActiveModal] = useState(null);
    const openModal = (modalId) => setActiveModal(modalId);
    const closeModal = () => setActiveModal(null);

    const {logout} = useContext(AuthContext);
    const {user} = useContext(AuthContext);

    const navigate = useNavigate();

    const handleBack = () => {
        if (user.role === "ADMIN") {
            navigate("/admin");
        } else if (user.role === "MANAGER") {
            navigate("/managerdashboard");
        } else if (user.role === "USER") {
            navigate("/clientdashboard");
        }
    }

    const handleLogout = () => {
        logout();
        navigate("/");
    }

    return (
        <div className={styles.container}>
            <h1 className={styles.text}>Settings</h1>
            <div className={styles.buttons}>
                <Button title="Log out" variant="type2" onClick={() => openModal("modalLogout")}/>
                <Button title="Back" variant="type3" onClick={handleBack}/>
            </div>
            <Modal title="Log out?" variant="type1" isOpen={activeModal === "modalLogout"} onClose={closeModal}>
                <Button title="Log out" variant="type4" onClick={handleLogout}/>
            </Modal>
        </div>
    )
}

export default SettingsPage;