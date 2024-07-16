package truncate

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scalecraft/snowguard/internal/copyhistory"
	"github.com/scalecraft/snowguard/internal/loginhistory"
)

type PostHandlerRequest struct {
}

type PostHandlerResponse struct {
}

func PostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := Delete()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"internal service error loading data.": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func Delete() (*PostHandlerResponse, error) {
	err := loginhistory.Truncate()
	if err != nil {
		return nil, err
	}

	err = copyhistory.Truncate()
	if err != nil {
		return nil, err
	}

	return &PostHandlerResponse{}, nil
}
