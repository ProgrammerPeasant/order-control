import React from "react";
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import {AuthProvider} from "./Utils/AuthProvider";

import StartPage from "./pages/StartPage";
import RegisterPage from "./pages/RegisterPage";
import LoginPage from "./pages/LoginPage";
import ClientDashboardPage from "./pages/ClientDashboardPage";
import SettingsPage from "./pages/SettingsPage";
import AdminDashboardPage from "./pages/AdminDashboardPage/AdminDashboardPage";
import ManagerDashboardPage from "./pages/ManagerDashboardPage";
import EstimateViewPage from "./pages/EstimateViewPage";


function App() {
    return (
        <AuthProvider>
            <Router>
                <Routes>
                    <Route path="/" element={<StartPage />} />
                    <Route path="/register" element={<RegisterPage />} />
                    <Route path="/login" element={<LoginPage />} />
                    <Route path="/clientdashboard" element={<ClientDashboardPage />} />
                    <Route path="/settings" element={<SettingsPage />} />
                    <Route path="/admin" element={<AdminDashboardPage />} />
                    <Route path="/managerdashboard" element={<ManagerDashboardPage />} />
                    <Route path="/estimateview/:estimateId" element={<EstimateViewPage />} />
                </Routes>
            </Router>
        </AuthProvider>
    )
}

export default App;