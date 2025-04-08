import React from 'react';
import styles from "./LoginPage.module.css"
import LogForm from "./LogForm";
import Button from "../../components/Button";
import {useNavigate} from "react-router-dom";


const LoginPage = () => {
    const navigate = useNavigate();

    return (
        <div className={styles.container}>
            <h1 className={styles.text}>Authentication</h1>
            <LogForm>
                <Button title="Back" variant="type3" onClick={() => navigate("/")} />
            </LogForm>
        </div>
    )
}

export default LoginPage;