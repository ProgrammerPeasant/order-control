import {useContext} from "react";
import {AuthContext} from "./AuthProvider";
import {Navigate} from "react-router-dom";

const RoleRoute = ({children, allowedRoles}) => {
    const {user} = useContext(AuthContext);

    if (!user || !allowedRoles.includes(user.role)) {
        return <Navigate to="/accessdenied" />
    }

    return children;
}

export default RoleRoute;