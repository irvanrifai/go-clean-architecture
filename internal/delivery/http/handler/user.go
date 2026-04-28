package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/irvanrifai/go-clean-architecture/internal/domain"
	"net/http"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(usecase domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.usecase.Register(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

// Router.go dalam paket yang sama
func NewRouter(handler *UserHandler) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", handler.Register)
	}
	return r
}