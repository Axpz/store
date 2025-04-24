package types

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Plan     string `json:"plan" binding:"required,oneof=free premium enterprise"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Plan     string `json:"plan" binding:"required,oneof=free premium enterprise"`
}

// User 用户信息
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Plan     string `json:"plan"`
	Created  int64  `json:"created"`
	Updated  int64  `json:"updated"`
}

// ErrorResponse 错误响应
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse 成功响应
type SuccessResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse 创建错误响应
func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}

// NewSuccessResponse 创建成功响应
func NewSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{
		Message: message,
	}
}
