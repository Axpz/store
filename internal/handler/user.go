package handler

import (
	"net/http"
	"time"

	"github.com/Axpz/store/internal/pkg/jwt"
	"github.com/Axpz/store/internal/service"
	"github.com/Axpz/store/internal/types"
	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService *service.UserService
	jwtSecret   string
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	var logger = utils.LoggerFromContext(c.Request.Context())
	logReq := req
	logReq.Password = "*"
	logger.Info("Login", zap.Any("req", logReq))

	user, err := h.userService.GetUser(c, utils.GetUserIDFromEmail(req.Email))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	h.userService.UpdateUserLastLogin(c, user.ID)

	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 生成token
	token, err := jwt.GenerateToken(user.ID, user.Username, h.jwtSecret, 7*24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败"})
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

// SignUp 处理用户注册
func (h *UserHandler) SignUp(c *gin.Context) {
	h.CreateUser(c)
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req types.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}

	// 保存用户
	if err := h.userService.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userGet, err := h.userService.GetUser(c, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户后获取用户失败"})
		return
	}

	c.JSON(http.StatusCreated, userGet)
}

// GetUser 获取用户信息
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUser(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser 更新用户信息
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user types.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	if err := h.userService.UpdateUser(c, id, user.Username, user.Email, user.Plan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
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

// Verify 验证用户会话
func (h *UserHandler) Verify(c *gin.Context) {
	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.userService.GetUser(c, userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
