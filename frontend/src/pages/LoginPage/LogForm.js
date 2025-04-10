import React, {useContext} from "react";
import {useNavigate} from "react-router-dom";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";
import {AuthContext} from "../../Utils/AuthProvider";
import {handleErrorMessage} from "../../Utils/ErrorHandler";


function LogForm({children}) {
    const {login} = useContext(AuthContext);

    const navigate = useNavigate();

    const fields = [
        {id: "username", type: "text", placeholder: "Username"},
        {id: "password", type: "password", placeholder: "Password"},
    ];

    const handleSubmit = async (e, formData) => {
        e.preventDefault();
        try {
            const response = await apiClient.post("/api/login", formData, {
                headers: {"Content-Type": "application/json", "Accept": "application/json"},
            });

            console.log(response.data);
            const {token, username, role, userId, companyId} = response.data;
            login(token, username, role, userId, companyId);
            if (role === "ADMIN") {
                navigate("/admin");
            } else if (role === "MANAGER") {
                navigate("/managerdashboard");
            } else if (role === "USER") {
                navigate("/clientdashboard");
            }
        } catch (error) {
            alert(handleErrorMessage(error, true));
        }
    }

    return (
        <Form fields={fields} handleSubmit={handleSubmit}>
            {children}
        </Form>
    )
}

export default LogForm;
