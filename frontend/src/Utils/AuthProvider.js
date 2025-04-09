import React, {createContext, useState} from "react";


export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(JSON.parse(localStorage.getItem("user")));

    const login = (token, username, role, userId, companyId) => {
        localStorage.setItem("authToken", token);
        localStorage.setItem("user", JSON.stringify({ username, role, userId, companyId }));
        setUser({ username, role, userId, companyId });
    };

    const logout = () => {
        localStorage.removeItem("authToken");
        localStorage.removeItem("user");
        setUser(null);
    };

    return (
        <AuthContext.Provider value={{ user, login, logout }}>
            {children}
        </AuthContext.Provider>
    )
}