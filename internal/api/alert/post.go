package alert

import (
	"bytes"
	"database/sql"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/scalecraft/skiguard/internal/config"
	"github.com/scalecraft/skiguard/internal/duckdb"
	"github.com/scalecraft/skiguard/internal/slack"
)

type Alert struct {
	Hour        string
	MetricValue int64
	RowNumber   int64
	P100        int64
}

type PostHandlerRequest struct {
	cfg *config.Config
}

type PostHandlerResponse struct {
}

func PostHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := PostHandlerRequest{cfg: cfg}
		req.cfg.NowUTC = time.Now().UTC().Format("2006-01-02 15:04:05")

		response, err := execute(req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"internal service error sending slack alert.": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func execute(req PostHandlerRequest) (*PostHandlerResponse, error) {
	alertCategories := []string{
		"failed_logins",
		"total_logins",
		"rows_copied",
		"users_created",
		"users_deleted",
		"total_copies",
	}
	dir := "internal/api/alert/templates/"
	for _, category := range alertCategories {
		tmplFile := category + ".tmpl.sql"
		var sql bytes.Buffer

		tmpl, err := template.New(tmplFile).ParseFiles(dir + tmplFile)
		if err != nil {
			return nil, err
		}

		err = tmpl.Execute(&sql, req.cfg)
		if err != nil {
			return nil, err
		}
		rows, err := duckdb.Query(sql.String())
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		err = alert(req, rows, category)
	}

	return &PostHandlerResponse{}, nil
}

func alert(req PostHandlerRequest, rows *sql.Rows, category string) error {
	alerts := []Alert{}
	for rows.Next() {
		var alert Alert
		err := rows.Scan(
			&alert.Hour,
			&alert.MetricValue,
			&alert.RowNumber,
			&alert.P100,
		)
		if err != nil {
			return err
		}
		alerts = append(alerts, alert)
	}

	if len(alerts) == 0 {
		return nil
	}

	tableString := writeTable(category, alerts)

	slackAlert := slack.SlackAlertMessage{
		SnowflakeActivity: category,
		Timestamp:         req.cfg.NowUTC,
		SnowflakeAccount:  req.cfg.SnowflakeAccount,
		TableString:       tableString,
		Token:             req.cfg.SlackToken,
		ChannelID:         req.cfg.SlackChannelId,
		AlertThreshold:    req.cfg.AlertThreshold,
	}
	slackAlert.Send()

	return nil
}

func writeTable(category string, alerts []Alert) string {
	var tableString bytes.Buffer
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(&tableString)
	t.AppendHeader(table.Row{"Metric Name", "Hour", "Metric Value", "p100 (30 days)"})
	for _, alert := range alerts {
		t.AppendRow(table.Row{category, alert.Hour, alert.MetricValue, alert.P100})
	}

	t.Render()

	return tableString.String()
}
