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
        if (user.role === "CLIENT") {
            navigate("/clientdashboard");
        } else if (user.role === "ADMIN") {
            navigate("/admin");
        } else if (user.role === "MANAGER") {
            // navigate("/managerdashboard"); // ATTENTION
        }
    }

    const handleLogout = () => {
        console.log("Logout")
        openModal("modalLogout")
    }

    const handleLogoutFinally = () => {
        logout();
        navigate("/");
    }

    const handleReset = () => {
        console.log("Reset")
        openModal("modalReset")
    }

    const handleResetFinally = () => {
        console.log("ResetFinally") // ATTENTION
    }

    return (
        <div className={styles.container}>
            <h1 className={styles.text}>Settings</h1>
            <div className={styles.buttons}>
                <Button title="Log out" variant="type2" onClick={handleLogout} />
                <Button title="Reset password" variant="type2" onClick={handleReset} />
                <Button title="Back" variant="type3" onClick={handleBack} />
            </div>

            <Modal title="Log out?" variant="type1" isOpen={activeModal === "modalLogout"} onClose={closeModal}>
                <Button title="Log out" variant="type4" onClick={handleLogoutFinally} />
            </Modal>
            <Modal title="Reset password?" variant="type1" isOpen={activeModal === "modalReset"} onClose={closeModal}>
                <Button title="Reset" variant="type2" onClick={handleResetFinally} />
            </Modal>
        </div>
    )
}

export default SettingsPage;