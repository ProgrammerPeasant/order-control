import React from "react";
import styles from "./RegisterPage.module.css"
import RegForm from "./RegForm";
import {useNavigate} from "react-router-dom";
import Button from "../../components/Button";


const RegisterPage = () => {
    const navigate = useNavigate();

    const handleBack = () => {
        console.log("Back to start")
        navigate("/");
    }
    return (
        <div className={styles.container}>
            <h1 className={styles.text}>Registration</h1>
            <RegForm>
                <Button title="Back" variant="type3" onClick={handleBack} />
            </RegForm>
        </div>
    )
}

export default RegisterPage;