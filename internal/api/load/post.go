package load

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scalecraft/skiguard/internal/config"
	"github.com/scalecraft/skiguard/internal/copyhistory"
	"github.com/scalecraft/skiguard/internal/duckdb"
	"github.com/scalecraft/skiguard/internal/loginhistory"
	"github.com/scalecraft/skiguard/internal/snowflake"
	"github.com/scalecraft/skiguard/internal/users"
)

type PostHandlerRequest struct {
	Cfg *config.Config
}

type PostHandlerResponse struct {
}

func PostHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := PostHandlerRequest{Cfg: cfg}

		response, err := DataLoad(req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"internal service error loading data.": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func DataLoad(req PostHandlerRequest) (*PostHandlerResponse, error) {
	snowflakeDB, err := snowflake.Connect(req.Cfg.SnowflakeDSN)
	if err != nil {
		return nil, err
	}

	err = loginhistory.Load(snowflakeDB)
	if err != nil {
		return nil, err
	}

	err = copyhistory.Load(snowflakeDB)
	if err != nil {
		return nil, err
	}

	err = users.Load(snowflakeDB)
	if err != nil {
		return nil, err
	}

	err = duckdb.Execute("checkpoint;")
	if err != nil {
		return nil, err
	}

	return &PostHandlerResponse{}, nil
}
