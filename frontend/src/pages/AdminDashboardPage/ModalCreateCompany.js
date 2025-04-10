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
        {id: "color1", type: "color", placeholder: "Color 1", value: "#4a7dff", required: false},
        {id: "color2", type: "color", placeholder: "Color 2", value: "#003b62", required: false},
        {id: "color3", type: "color", placeholder: "Color 3", value: "#ff6cb4", required: false},
    ];

    const handleSubmit = async (e, formData) => {
        e.preventDefault();
        const updatedData = {
            ...formData,
            design_colors: [formData?.color1, formData?.color2, formData?.color3]
        };

        try {
            const response = await apiClient.post("/api/v1/companies", updatedData, {
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