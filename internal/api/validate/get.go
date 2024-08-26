package validate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scalecraft/skiguard/internal/config"
)

type GetHandlerRequest struct {
	cfg *config.Config
}

type GetHandlerResponse struct {
}

func GetHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := GetHandlerRequest{cfg: cfg}

		response, err := validate(&req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"internal service error loading data.": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func validate(req *GetHandlerRequest) (*GetHandlerResponse, error) {

	return &GetHandlerResponse{}, nil
}
