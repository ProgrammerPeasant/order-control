import Modal from "../../components/Modal";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import Form from "../../components/Form";
import apiClient from "../../Utils/apiClient";

const ModalCreateEstimate = ({estimateId, isOpen, onClose, handleUpdate}) => {
    const fields = [
        {id: "title", type: "text", placeholder: "Title", required: true},
        {id: "overall_discount_percent", type: "number", placeholder: "Overall discount", value: "0", required: true},
    ]

    const handleSubmit = async (e, formData) => {
        e.preventDefault();
        const updatedData = {
            ...formData,
            overall_discount_percent: parseFloat(formData.overall_discount_percent),
            // ATTENTION company_id
        }
        try {
            const response = await apiClient.post("/api/v1/estimates", updatedData, {
                headers: {"Content-Type": "application/json", Accept: "application/json"},
            });
            console.log(response.data);
            onClose();
            handleUpdate(estimateId)
        } catch (error) {
            handleErrorMessage(error)
        }
    }

    return (
        <Modal title="Create Estimate" variant="type2" isOpen={isOpen} onClose={onClose}>
            <Form fields={fields} handleSubmit={handleSubmit} />
        </Modal>
    )
}

export default ModalCreateEstimate;