import Modal from "../../components/Modal";
import RegForm from "./RegForm";


const ModalRegisterUser = ({isOpen, onClose}) => {
    return (
        <Modal title="Create User" variant="type2" isOpen={isOpen} onClose={onClose}>
            <RegForm onClose={onClose} />
        </Modal>
    )
}

export default ModalRegisterUser;