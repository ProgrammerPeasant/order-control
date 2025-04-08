import React, {useContext} from "react";
import {useNavigate} from "react-router-dom";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";
import {AuthContext} from "../../Utils/AuthProvider";
import {handleErrorMessage} from "../../Utils/ErrorHandler";


function RegForm({children}) {
    const {login} = useContext(AuthContext);

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
        try {
            const response = await apiClient.post("/api/register", formData, {
                headers: {"Content-Type": "application/json", "Accept": "application/json"},
            });

            console.log(response.data);
        } catch (error) {
            alert(handleErrorMessage(error));
        }
        try {
            const response = await apiClient.post("/api/login", formData, {
                headers: {"Content-Type": "application/json", "Accept": "application/json"},
            });

            console.log(response.data);
            const {token, username, role, userId} = response.data;
            login(token, username, role, userId);
            if (role === "ADMIN") {
                navigate("/admin");
            } else if (role === "CLIENT") {
                navigate("/clientdashboard");
            } else if (role === "MANAGER") {
                navigate("/managerdashboard");
            }
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