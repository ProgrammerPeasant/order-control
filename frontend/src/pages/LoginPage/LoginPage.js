import React from 'react';
import styles from "./LoginPage.module.css"
import LogForm from "./LogForm";


const LoginPage = () => {
    return (
        <div className={styles.container}>
            <h1 className={styles.text}>Authentication</h1>
            <LogForm />
        </div>
    )
}

export default LoginPage;