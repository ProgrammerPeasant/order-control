import Modal from "../../components/Modal";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";


const ModalCreateCompany = ({isOpen, onClose}) => {
    const fields = [
        {id: "name", type: "text", placeholder: "Name"},
        {id: "address", type: "text", placeholder: "Address"},
        {id: "desc", type: "text", placeholder: "Description", required: false},
        {id: "logo_url", type: "text", placeholder: "Logo Url", required: false},
        {id: "color_primary", type: "color", placeholder: "Primary color", value: "#4a7dff"},
        {id: "color_secondary", type: "color", placeholder: "Secondary color", value: "#003b62"},
        {id: "color_accent", type: "color", placeholder: "Accent color", value: "#ff6cb4"},
    ];

    const handleSubmit = async (e, formData) => {
        e.preventDefault();
        try {
            const response = await apiClient.post("/api/v1/companies", formData, {
                headers: { "Content-Type": "application/json", "Accept": "application/json" },
            });
            console.log(response.data);
            onClose();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    return (
        <Modal title="Create company" variant="type2" isOpen={isOpen} onClose={onClose}>
            <Form fields={fields} handleSubmit={handleSubmit} />
        </Modal>
    )
}

export default ModalCreateCompany;