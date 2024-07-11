package load

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scalecraft/snowguard/internal/copyhistory"
	"github.com/scalecraft/snowguard/internal/loginhistory"
	"github.com/scalecraft/snowguard/internal/users"
)

type PostHandlerRequest struct {
	db *sql.DB
}

type PostHandlerResponse struct {
}

func PostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := PostHandlerRequest{db: db}

		response, err := dataLoad(req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"internal service error loading data.": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func dataLoad(req PostHandlerRequest) (*PostHandlerResponse, error) {
	err := loginhistory.Load(req.db)
	if err != nil {
		return nil, err
	}

	err = copyhistory.Load(req.db)
	if err != nil {
		return nil, err
	}

	err = users.Load(req.db)
	if err != nil {
		return nil, err
	}

	return &PostHandlerResponse{}, nil
}
