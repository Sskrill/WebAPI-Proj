package webAPIUsers

import (
	"github.com/gin-gonic/gin"
)

func NewRouting(h *Handler) {
	rout := gin.Default()

	rout.GET("/users/:id", h.GetUser)
	rout.PUT("/users/:id", h.UpdateUser)
	rout.POST("/users", h.CreateUser)
	rout.DELETE("/users/:id", h.DeletUser)
	rout.GET("/users", h.GetAllUsers)
	rout.Run()
}
