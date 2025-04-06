import React, {createContext, useState} from "react";


export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(JSON.parse(localStorage.getItem("user")));

    const login = (token, username, role, userId) => {
        localStorage.setItem("authToken", token);
        localStorage.setItem("user", JSON.stringify({ username, role, userId }));
        setUser({ username, role, userId });
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