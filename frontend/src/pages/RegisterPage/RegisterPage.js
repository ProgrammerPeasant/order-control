import React from "react";
import styles from "./RegisterPage.module.css"
import RegForm from "./RegForm";
import {useNavigate} from "react-router-dom";
import Button from "../../components/Button";


const RegisterPage = () => {
    const navigate = useNavigate();

    return (
        <div className={styles.container}>
            <h1 className={styles.text}>Registration</h1>
            <RegForm>
                <Button title="Back" variant="type3" onClick={() => navigate("/")} />
            </RegForm>
        </div>
    )
}

export default RegisterPage;