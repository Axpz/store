package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Axpz/store/internal/pkg/jwt"
	"github.com/Axpz/store/internal/service"
	"github.com/Axpz/store/internal/types"
	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService  *service.UserService
	emailService *service.EmailService
	jwtSecret    string
}

// NewUserHandler 创建用户处理器
func NewUserHandler(userService *service.UserService, jwtSecret string) *UserHandler {
	return &UserHandler{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		// 登录注册等公开接口
		auth.POST("/login", h.Login)
		auth.POST("/logout", h.Logout)
		auth.POST("/signup", h.SignUp)
		auth.GET("/verify", h.Verify)
	}

	// 需要认证的路由组
	user := router.Group("/api/users")
	// authorized.Use(middleware.Auth(cfg.JWT.Secret))
	{
		// 用户相关路由
		user.POST("", h.CreateUser)
		user.GET("/:id", h.GetUser)
		user.PUT("/:id", h.UpdateUser)
		user.DELETE("/:id", h.DeleteUser)
	}
}

// Login 处理用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var req types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request parameters"})
		return
	}

	var logger = utils.LoggerFromContext(c.Request.Context())
	logReq := req
	logReq.Password = "*"
	logger.Info("Login", zap.Any("req", logReq))

	user, err := h.userService.GetUser(c, utils.GetUserIDFromEmail(req.Email))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if user.Verified != nil && !*user.Verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not verified, please check your email for verification."})
		return
	}

	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	h.userService.UpdateUserLastLogin(c, user.ID)

	// 生成token
	token, err := jwt.GenerateToken(user.ID, user.Username, h.jwtSecret, 7*24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// 设置httponly cookie，有效期7天
	c.SetCookie(
		"token",    // cookie name
		token,      // cookie value
		7*24*60*60, // max age in seconds (7 days)
		"/",        // path
		"",         // domain
		true,       // secure
		true,       // httponly
	)

	user.Password = "*"
	c.JSON(http.StatusOK, types.LoginResponse{
		Token: token,
		User:  *user,
	})

}

// Logout 处理用户登出
func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// SignUp
func (h *UserHandler) SignUp(c *gin.Context) {
	var req types.RegisterRequest
	var logger = utils.LoggerFromContext(c.Request.Context())

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request parameters" + err.Error()})
		return
	}

	logReq := req
	logReq.Password = "*"
	logger.Info("SignUp", zap.Any("req", logReq))

	user := types.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	user.HashPassword()

	userGet, err := h.userService.GetUser(c, utils.GetUserIDFromEmail(user.Email))
	if err != nil && status.Code(err) != codes.NotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user:" + err.Error()})
		return
	}

	if userGet != nil {
		if userGet.Verified == nil || *userGet.Verified {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}
	} else {
		if err := h.userService.CreateUser(c, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user:" + err.Error()})
			return
		}
	}

	tokenString, err := user.GenVerificationJWTToken(h.jwtSecret, 7*24*time.Hour)
	if err != nil {
		logger.Error("failed to generate verification token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate verification token:" + err.Error()})
		return
	}

	verificationLink := fmt.Sprintf("https://www.axpz.org/api/auth/verify?token=%s", tokenString)
	logger.Info("Verification link", zap.String("link", verificationLink))

	if err := h.emailService.SendVerificationEmail(c, verificationLink, user.Email); err != nil {
		logger.Error("failed to send verification email", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email:" + err.Error()})
		return
	}

	logger.Info("Verification email sent", zap.String("email", user.Email))
	c.JSON(http.StatusOK, gin.H{"message": "Verification email sent"})
}

// CreateUser
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req types.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request parameters"})
		return
	}

	var logger = utils.LoggerFromContext(c.Request.Context())
	logReq := req
	logReq.Password = "*"
	logger.Info("CreateUser", zap.Any("req", logReq))

	// 创建用户
	user := types.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	// 加密密码
	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// 保存用户
	if err := h.userService.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userGet, err := h.userService.GetUser(c, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user after creating"})
		return
	}

	c.JSON(http.StatusCreated, userGet)
}

// GetUser 获取用户信息
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUser(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser 更新用户信息
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var userReq types.UpdateUserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request parameters"})
		return
	}

	user := types.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
		Plan:     userReq.Plan,
	}

	if userReq.Password != "" {
		user.Password = userReq.Password
		user.HashPassword()
	}

	logger := utils.LoggerFromContext(c.Request.Context())
	logReq := userReq
	logReq.Password = "*"
	logger.Info("UpdateUser", zap.Any("req", logReq))

	if err := h.userService.UpdateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userGet, err := h.userService.GetUser(c, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user after updating"})
		return
	}

	c.JSON(http.StatusOK, userGet)
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.userService.DeleteUser(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *UserHandler) Verify(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing token"})
		return
	}

	user := types.User{}

	user, err := user.VerifyAndParseVerificationJWTToken(h.jwtSecret, tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or expired token" + err.Error()})
		return
	}

	user.Verified = &[]bool{true}[0]

	h.userService.UpdateUser(c, &user)

	c.Redirect(http.StatusFound, "https://www.axpz.org/login?verify=success")
}
