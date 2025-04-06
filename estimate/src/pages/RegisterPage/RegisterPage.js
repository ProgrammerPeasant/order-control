import React from "react";
import styles from "./RegisterPage.module.css"
import RegForm from "./RegForm";


const RegisterPage = () => {
    return (
        <div className={styles.container}>
            <h1 className={styles.text}>Registration</h1>
            <RegForm />
        </div>
    )
}

export default RegisterPage;