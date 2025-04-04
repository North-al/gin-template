package handler

import (
	"net/http"

	"github.com/North-al/gin-template/internal/biz/entity"
	"github.com/North-al/gin-template/internal/biz/service"
	"github.com/North-al/gin-template/internal/pkg"
	"github.com/North-al/gin-template/internal/types/rest"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterRoutes() []*rest.RouteConfig {
	return []*rest.RouteConfig{
		rest.NewRouteConfig("/auth/login", h.Login, http.MethodPost, false),
		rest.NewRouteConfig("/auth/register", h.Register, http.MethodPost, false),
	}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.LoginRequest true "Login Request"
// @Success 200 {object} entity.UserEntity "Login Response"
// @Failure 400 {object} pkg.Response "Bad Request"
// @Failure 500 {object} pkg.Response "Internal Server Error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req entity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.BadRequest(c, err.Error())
	}

	user, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		pkg.InternalError(c, err.Error())
		return
	}

	pkg.Success(c, user)
}

// Register 用户注册
func (h *AuthHandler) Register(c *gin.Context) {
	var req entity.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.BadRequest(c, err.Error())
	}

	user, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		pkg.InternalError(c, err.Error())
		return
	}

	pkg.Success(c, user)
}
