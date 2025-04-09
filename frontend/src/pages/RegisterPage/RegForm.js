import React from "react";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import {useNavigate} from "react-router-dom";


function RegForm({children}) {
    const navigate = useNavigate();

    const fields = [
        {id: "username", type: "text", placeholder: "Username", required: true},
        {id: "company_id", type: "number", placeholder: "Company ID", required: true},
        {id: "email", type: "email", placeholder: "Email", required: true},
        {id: "password", type: "password", placeholder: "Password", required: true},
        {id: "confirmPassword", type: "password", placeholder: "Password confirmation", required: true},
    ];

    const handleSubmit = async (e, formData) => {
        e.preventDefault()
        if (formData.password && formData.confirmPassword && formData.password !== formData.confirmPassword) {
            alert("Passwords don't match")
            return;
        }

        const updatedData = {
            ...formData,
            company_id: parseInt(formData.company_id, 10),
        }

        try {
            const response = await apiClient.post("/api/register", updatedData, {
                headers: {"Content-Type": "application/json", "Accept": "application/json"},
            });

            console.log(response.data);
            navigate("/login");
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    };

    return (
        <Form fields={fields} handleSubmit={handleSubmit}>
            {children}
        </Form>
    )
}

export default RegForm