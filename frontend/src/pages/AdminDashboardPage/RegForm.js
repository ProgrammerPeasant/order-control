import apiClient from "../../Utils/apiClient";
import Form from "../../components/Form";
import {handleErrorMessage} from "../../Utils/ErrorHandler";


function RegForm({onClose}) {
    const fields = [
        {id: "role", type: "text", placeholder: "Role (USER, MANAGER or ADMIN)"},
        {id: "email", type: "email", placeholder: "Email"},
        {id: "username", type: "text", placeholder: "Username"},
        {id: "password", type: "password", placeholder: "Password"},
        {id: "confirmPassword", type: "password", placeholder: "Confirm Password"},
        {id: "company_id", type: "number", placeholder: "Company ID"},
    ];

    const handleSubmit = async (e, formData) => {
        e.preventDefault();
        if (formData.password && formData.confirmPassword && formData.password !== formData.confirmPassword) {
            alert("Passwords don't match");
            return;
        }

        const updatedData = {
            ...formData,
            company_id: parseInt(formData.company_id, 10),
        }
        try {
            const response = await apiClient.post("/api/admin/register", updatedData, {
                headers: {"Content-Type": "application/json", "Accept": "application/json"}
            });
            console.log(response.data);
            onClose();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    return (
        <Form fields={fields} handleSubmit={handleSubmit} />
    )
}

export default RegForm