import Table from "../../components/Table";
import Button from "../../components/Button";
import {handleErrorMessage} from "../../Utils/ErrorHandler";
import apiClient from "../../Utils/apiClient";


function RequestTable({onClose}) {
    const columns = ["ID", "Name", "", ""]
    const apiUrl = "/api/v1/companies/join-request"

    const handleApprove = async (userId) => {
        try {
            const response = await apiClient.post("/api/v1/companies/join-request/approve", {user_id: userId}, {headers: {"Content-Type": "application/json", Accept: "application/json"}});
            console.log(response.data);
            onClose();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    const handleReject = async (userId) => {
        try {
            const response = await apiClient.post("/api/v1/companies/join-request/reject", {user_id: userId}, {headers: {"Content-Type": "application/json", Accept: "application/json"}});
            console.log(response.data);
            onClose();
        } catch (error) {
            alert(handleErrorMessage(error));
        }
    }

    const renderRow = (item) => (
        <tr key={item.ID}>
            <td>{item.user_id}</td>
            <td>{item.Email}</td>
            <td><Button title="Approve" variant="type1" onClick={() => handleApprove(item.user_id)} /></td>
            <td><Button title="Reject" variant="type4" onClick={() => handleReject(item.user_id)} /></td>
        </tr>
    )

    return (
        <div>
            <Table apiUrl={apiUrl} columns={columns} renderRow={renderRow} emptyRows={3} />
        </div>
    )
}

export default RequestTable;