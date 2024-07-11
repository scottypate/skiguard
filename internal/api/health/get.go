package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetHandlerRequest struct {
}

type GetHandlerResponse struct {
}

func GetHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}
