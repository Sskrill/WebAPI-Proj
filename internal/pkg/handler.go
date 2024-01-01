package webAPIUsers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sskrill/WebAPI-Proj/internal/cache"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}
type Handler struct {
	crud  CRUD
	cache *cache.Cache
}

func NewHandler(crud CRUD, cache *cache.Cache) *Handler {
	return &Handler{crud: crud, cache: cache}
}
func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	user, ok := h.cache.Get(id)
	if ok {
		c.JSON(http.StatusAccepted, user)
		return
	}
	user, err = h.crud.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	h.cache.Set(id, user)
	c.JSON(http.StatusOK, user)
}
func (h *Handler) CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}
	err := h.crud.Insert(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Created User")
}
func (h *Handler) UpdateUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	err = h.crud.Update(id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Updated user")

}
func (h *Handler) DeletUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
	}
	err = h.crud.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Deleted user")

}
func (h *Handler) GetAllUsers(c *gin.Context) {

	users := h.crud.GetAll()
	c.JSON(http.StatusOK, users)
}
