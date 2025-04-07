import Modal from "../../components/Modal";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";


const ModalCreateCompany = ({isOpen, onClose}) => {
    const fields = [
        // {id: "id", type: "number", placeholder: "ID", required: true},
        {id: "name", type: "text", placeholder: "Name", required: true},
        {id: "address", type: "text", placeholder: "Address", required: true},
        {id: "desc", type: "text", placeholder: "Description", required: true},
        // {id: "createdAt", type: "date", placeholder: "CreatedAt", required: true},
        // {id: "updatedAt", type: "date", placeholder: "UpdatedAt", required: true},
        // {id: "deletedAt", type: "date", placeholder: "DeletedAt", required: true},
    ];

    const handleSubmit = async (e, formData) => {
        e.preventDefault();

        const updatedData = {
            ...formData,
            id: "228", // ATTENTION
            createdAt: new Date(), // ATTENTION
            updatedAt: new Date(), // ATTENTION
            deletedAt: null, // ATTENTION
        }

        try {
            const response = await apiClient.post("/api/v1/companies", updatedData, {
                headers: { "Content-Type": "application/json", "Accept": "application/json" },
            });

            console.log(response.data);
            onClose();
        } catch (error) {
            if (error.response) {
                const {status, data} = error.response;
                if (status === 400) {
                    alert(`Invalid data: ${data.message || "Check your input fields"}`);
                } else if (status === 401) {
                    alert(`Unauthorized: ${data.message}`);
                } else if (status === 403) {
                    alert(`Access denied: ${data.message || "Admin only"}`);
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
        <Modal title="Create company" variant="type2" isOpen={isOpen} onClose={onClose}>
            <Form fields={fields} handleSubmit={handleSubmit} />
        </Modal>
    )
}

export default ModalCreateCompany;