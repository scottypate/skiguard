package update

import (
	"database/sql"
	"fmt"
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

		response, err := update(req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"internal service error loading data.": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func update(req PostHandlerRequest) (*PostHandlerResponse, error) {
	err := loginhistory.Update(req.db)
	if err != nil {
		return nil, err
	}

	err = copyhistory.Update(req.db)
	if err != nil {
		return nil, err
	}

	fmt.Println("here")
	err = users.Update(req.db)
	if err != nil {
		return nil, err
	}
	fmt.Println("here2")
	return &PostHandlerResponse{}, nil
}
