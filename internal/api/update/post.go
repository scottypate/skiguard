package update

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scalecraft/snowguard/internal/config"
	"github.com/scalecraft/snowguard/internal/copyhistory"
	"github.com/scalecraft/snowguard/internal/duckdb"
	"github.com/scalecraft/snowguard/internal/loginhistory"
	"github.com/scalecraft/snowguard/internal/snowflake"
	"github.com/scalecraft/snowguard/internal/users"
)

type PostHandlerRequest struct {
	cfg *config.Config
}

type PostHandlerResponse struct {
}

func PostHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := PostHandlerRequest{cfg: cfg}

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
	snowflakeDB, err := snowflake.Connect(req.cfg.SnowflakeDSN)
	if err != nil {
		return nil, err
	}

	err = loginhistory.Update(snowflakeDB)
	if err != nil {
		return nil, err
	}

	err = copyhistory.Update(snowflakeDB)
	if err != nil {
		return nil, err
	}

	err = users.Update(snowflakeDB)
	if err != nil {
		return nil, err
	}

	err = duckdb.Execute("checkpoint;")
	if err != nil {
		return nil, err
	}

	return &PostHandlerResponse{}, nil
}
