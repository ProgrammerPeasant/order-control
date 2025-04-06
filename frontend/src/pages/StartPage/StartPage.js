import React from "react"
import Button from "../../components/Button"
import styles from "./StartPage.module.css"
import {useNavigate} from "react-router-dom"


const StartPage = () => {
    const navigate = useNavigate()

    const handleLogin = () => {
        console.log("Login")
        navigate("/login")
    }

    const handleRegister = () => {
        console.log("Register")
        navigate("/register")
    }

    return (
        <div className={styles.container}>
            <h1 className={styles.text}>Estimate-control</h1>
            <div className={styles.buttons}>
                <Button title="Sign in" variant="type1" onClick={handleLogin} />
                <Button title="Sign up" variant="type2" onClick={handleRegister} />
            </div>
        </div>
    )
}

export default StartPage;