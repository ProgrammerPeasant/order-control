import React from "react";
import styles from "./Button.module.css"


const Button = ({variant, onClick, type = "button", ...props}) => {
    return (<button className={`${styles.button} ${styles[`button--${variant}`]}`} onClick={onClick}
                    type={type}>{props.title}</button>)
}

export default Button;