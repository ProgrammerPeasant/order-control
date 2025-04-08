import React, {useState} from "react";
import styles from "./Form.module.css"
import Button from "./Button";


function Form({fields, handleSubmit, children}) {
    const [formData, setFormData] = useState(
        fields.reduce((acc, field) => {
            acc[field.id] = field.value || "";
            return acc;
        }, {})
    );

    const handleChange = (e) => {
        setFormData({...formData, [e.target.id]: e.target.value});
    };

    return (
        <form className={styles.form} onSubmit={(e) => handleSubmit(e, formData)}>
            {fields.map((field) => (
                <input
                    key={field.id}
                    type={field.type}
                    id={field.id}
                    placeholder={field.placeholder}
                    value={formData[field.id]}
                    onChange={handleChange}
                    required={field.required}
                />
            ))}
            <div className={styles.buttons}>
                {children}
                <Button title="Continue" variant="type2" type="submit" />
            </div>
        </form>
    );
}

export default Form;