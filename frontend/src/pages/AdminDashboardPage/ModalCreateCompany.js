import Modal from "../../components/Modal";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";
import {handleErrorMessage} from "../../Utils/ErrorHandler";


const ModalCreateCompany = ({isOpen, onClose}) => {
    const fields = [
        {id: "name", type: "text", placeholder: "Name", required: true},
        {id: "address", type: "text", placeholder: "Address", required: true},
        {id: "desc", type: "text", placeholder: "Description", required: true}
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