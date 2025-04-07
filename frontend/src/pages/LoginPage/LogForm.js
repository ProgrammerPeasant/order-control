import React, {useContext} from "react";
import {useNavigate} from "react-router-dom";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";
import {AuthContext} from "../../Utils/AuthProvider";


function LogForm({children}) {
    const {login} = useContext(AuthContext);

    const navigate = useNavigate();

    const fields = [
        {id: "username", type: "text", placeholder: "Username", required: true},
        {id: "password", type: "password", placeholder: "Password", required: true},
    ];

    const handleSubmit = async (e, formData) => {
        e.preventDefault();
        try {
            const response = await apiClient.post("/api/login", formData, {
                headers: {"Content-Type": "application/json", "Accept": "application/json"},
            });

            const {token, username, role, userId} = response.data;
            login(token, username, role, userId);
            if (role === "ADMIN") {
                navigate("/admin");
            } else if (role === "CLIENT") {
                navigate("/clientdashboard");
            } else if (role === "MANAGER") {
                console.log("managerdash")
                // navigate("/managerdashboard"); // ATTENTION
            }
            console.log(response.data);
        } catch (error) {
            if (error.response) {
                const { status, data } = error.response;
                if (status === 400) {
                    alert(`Invalid data: ${data.message || "Check your input fields"}`);
                } else if (status === 401) {
                    alert(`Authentication error: ${data.message || "Invalid credentials"}`);
                } else if (status === 500) {
                    alert(`Server error: ${data.message || "Please try again later"}`);
                } else {
                    alert(`Error: ${data.message || "Something went wrong"}`);
                }
            } else if (error.request) {
                alert("No response from server. Please check your internet connection.");
            } else {
                alert(`Request error: ${error.message}`);
            }
        }
    }

    return (
        <Form fields={fields} handleSubmit={handleSubmit}>
            {children}
        </Form>
    )
}

export default LogForm;
