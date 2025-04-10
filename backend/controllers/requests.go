package controllers

// LoginRequest ...
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// StandardRegisterRequest для обычной регистрации
type StandardRegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	CompanyID uint   `json:"company_id" binding:"required"`
}

// AdminRegisterRequest для административной регистрации
type AdminRegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	Role      string `json:"role" binding:"required,oneof=ADMIN MANAGER USER"`
	CompanyID uint   `json:"company_id" binding:"required"`
}

// ApproveRejectRequest для одобрения/отклонения запросов
type ApproveRejectRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}
