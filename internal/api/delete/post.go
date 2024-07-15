package delete

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scalecraft/snowguard/internal/copyhistory"
	"github.com/scalecraft/snowguard/internal/loginhistory"
	"github.com/scalecraft/snowguard/internal/users"
)

type DeleteHandlerRequest struct {
}

type DeleteHandlerResponse struct {
}

func DeleteHandler() gin.HandlerFunc {
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

func Delete() (*DeleteHandlerResponse, error) {
	err := loginhistory.Delete()
	if err != nil {
		return nil, err
	}

	err = copyhistory.Delete()
	if err != nil {
		return nil, err
	}

	err = users.Delete()
	if err != nil {
		return nil, err
	}

	return &DeleteHandlerResponse{}, nil
}
