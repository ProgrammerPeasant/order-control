import React from "react";
import styles from "./Modal.module.css"
import Button from "./Button";


const Modal = ({isOpen, onClose, children, ...props}) => {
    if (!isOpen) return null;

    return (
        <div className={styles.overlay} onClick={onClose}>
            <div className={styles.container} onClick={(e) => e.stopPropagation()}>
                <h1 className={styles.text}>{props.title}</h1>
                {children}
                <Button title="Close" variant="type3" onClick={onClose} />
            </div>
        </div>
    )
}

export default Modal;