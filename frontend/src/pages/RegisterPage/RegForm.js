import React, {useContext} from "react";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import {useNavigate} from "react-router-dom";
import {AuthContext} from "../../Utils/AuthProvider";


function RegForm({children}) {
    const {login} = useContext(AuthContext);
    const navigate = useNavigate();

    const fields = [
        {id: "username", type: "text", placeholder: "Username"},
        {id: "company_id", type: "number", placeholder: "Company ID"},
        {id: "email", type: "email", placeholder: "Email"},
        {id: "password", type: "password", placeholder: "Password"},
        {id: "confirmPassword", type: "password", placeholder: "Password confirmation"},
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
        } catch (error) {
            alert(handleErrorMessage(error));
        }
        try {
            const response = await apiClient.post("/api/login", updatedData, {
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "application/json"
                },
            });
            console.log(response.data);
            const {token, username, role, userId, companyId} = response.data;
            login(token, username, role, userId, companyId);
            navigate("/clientdashboard");
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