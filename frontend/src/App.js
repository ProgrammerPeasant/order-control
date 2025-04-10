import React from "react";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import {AuthProvider} from "./Utils/AuthProvider";
import PrivateRoute from "./Utils/PrivateRoute";
import RoleRoute from "./Utils/RoleRoute";

import StartPage from "./pages/StartPage";
import RegisterPage from "./pages/RegisterPage";
import LoginPage from "./pages/LoginPage";
import ClientDashboardPage from "./pages/ClientDashboardPage";
import SettingsPage from "./pages/SettingsPage";
import AdminDashboardPage from "./pages/AdminDashboardPage/AdminDashboardPage";
import ManagerDashboardPage from "./pages/ManagerDashboardPage";
import EstimateViewPage from "./pages/EstimateViewPage";
import NotFoundPage from "./pages/NotFoundPage";
import AccessDeniedPage from "./pages/AccessDeniedPage";


function App() {
    return (
        <AuthProvider>
            <Router>
                <Routes>
                    <Route path="/" element={<StartPage/>}/>
                    <Route path="/register" element={<RegisterPage/>}/>
                    <Route path="/login" element={<LoginPage/>}/>

                    <Route path="/clientdashboard" element={<PrivateRoute><RoleRoute
                        allowedRoles={["USER", "ADMIN"]}><ClientDashboardPage/></RoleRoute></PrivateRoute>}/>
                    <Route path="/settings"
                           element={<PrivateRoute><RoleRoute allowedRoles={["USER", "MANAGER", "ADMIN"]}><SettingsPage/></RoleRoute></PrivateRoute>}/>
                    <Route path="/admin"
                           element={<PrivateRoute><RoleRoute allowedRoles={["ADMIN"]}><AdminDashboardPage/></RoleRoute></PrivateRoute>}/>
                    <Route path="/managerdashboard" element={<PrivateRoute><RoleRoute
                        allowedRoles={["MANAGER", "ADMIN"]}></RoleRoute><ManagerDashboardPage/></PrivateRoute>}/>
                    <Route path="/estimateview/:estimateId" element={<PrivateRoute><RoleRoute
                        allowedRoles={["USER", "MANAGER", "ADMIN"]}><EstimateViewPage/></RoleRoute></PrivateRoute>}/>
                    <Route path="*" element={<NotFoundPage/>}/>
                    <Route path="/accessdenied" element={<AccessDeniedPage/>}/>
                </Routes>
            </Router>
        </AuthProvider>
    )
}

export default App;