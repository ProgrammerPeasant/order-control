import Modal from "../../components/Modal";
import RequestTable from "./RequestTable";


const ModalJoinRequests = ({isOpen, onClose}) => {
    return (
        <Modal title="Join Requests" variant="type2" isOpen={isOpen} onClose={onClose}>
            <RequestTable onClose={onClose}/>
        </Modal>
    )
}

export default ModalJoinRequests;