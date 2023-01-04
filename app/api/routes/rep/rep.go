package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Method to add a new rep
func CreateRep(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
