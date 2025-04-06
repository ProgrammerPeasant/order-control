import React, {useContext} from "react";
import Button from "../../components/Button";
import {useNavigate} from "react-router-dom";
import Form from "../../components/Form";
import apiClient from "../../apiClient";
import {AuthContext} from "../../AuthProvider";


function RegForm() {
    const {login} = useContext(AuthContext);

    const navigate = useNavigate();

    const handleBack = () => {
        console.log("Back to start")
        navigate("/");
    }

    const fields = [
        {id: "username", type: "text", placeholder: "Username", required: true},
        {id: "email", type: "email", placeholder: "Email", required: true},
        {id: "password", type: "password", placeholder: "Password", required: true},
        {id: "confirmPassword", type: "password", placeholder: "Password confirmation", required: true},
    ];

    const handleSubmit = async (e, formData) => {
        e.preventDefault()
        if (formData.password && formData.confirmPassword && formData.password !== formData.confirmPassword) {
            alert("Passwords don't match")
        }

        const updatedData = {
            ...formData,
            role: "CLIENT", // ATTENTION
            company_id: "1" // ATTENTION
        }

        try {
            const response = await apiClient.post("/api/register", updatedData, {
                headers: {"Content-Type": "application/json", "Accept": "application/json"},
            });

            console.log(response.data);
        } catch (error) {
            if (error.response) {
                const { status, data } = error.response;
                if (status === 400) {
                    alert(`Invalid data: ${data.message || "Check your input fields"}`);
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
                if (status === 500) {
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
    };

    return (
        <Form fields={fields} handleSubmit={handleSubmit}>
            <Button title="Back" variant="type3" onClick={handleBack} />
        </Form>
    )
}

export default RegForm