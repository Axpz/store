package types

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// User 用户信息
type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Plan      string `json:"plan"`
	Created   int64  `json:"created"`
	Updated   int64  `json:"updated"`
	LastLogin int64  `json:"last_login"`
	Verified  *bool  `json:"verified,omitempty"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Plan     string `json:"plan" binding:"required,oneof=free premium enterprise"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
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

// HashPassword 对密码进行加密
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword 检查密码是否正确
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type verificationClaims struct {
	User string `json:"user"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func (u *User) GenVerificationJWTToken(jwtSecret string, duration time.Duration) (string, error) {
	secret := []byte(jwtSecret)
	userJson, err := json.Marshal(u)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, verificationClaims{
		User: string(userJson),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *User) VerifyAndParseVerificationJWTToken(jwtSecret string, tokenString string) (User, error) {
	secret := []byte(jwtSecret)
	claims := &verificationClaims{}
	user := User{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return User{}, err
	}

	if !token.Valid {
		return User{}, errors.New("invalid token")
	}

	err = json.Unmarshal([]byte(claims.User), &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
