package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackAlertMessage struct {
	SnowflakeActivity string
	Timestamp         string
	SnowflakeAccount  string
	TableString       string
	Token             string
	ChannelID         string
	AlertThreshold    float64
}

func (s *SlackAlertMessage) Send() error {
	client := slack.New(s.Token)
	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf(
		":warning: Snowguard flagged the following activity as being abnormal.\n The following activity exceeded the alert threshold of `p%d` for the most recent hour.", int(s.AlertThreshold*100)),
		false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Fields
	typeField := slack.NewTextBlockObject("mrkdwn", "*Category:*\n"+s.SnowflakeActivity, false, false)
	whenField := slack.NewTextBlockObject("mrkdwn", "*When:*\n"+s.Timestamp, false, false)
	lastUpdateField := slack.NewTextBlockObject("mrkdwn", "*Snowflake Account:*\n"+s.SnowflakeAccount, false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, typeField)
	fieldSlice = append(fieldSlice, whenField)
	fieldSlice = append(fieldSlice, lastUpdateField)

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	// Table
	tableText := slack.NewTextBlockObject("mrkdwn", "```"+s.TableString+"```", false, false)
	tableSection := slack.NewSectionBlock(tableText, nil, nil)

	// Approve and Deny Buttons
	// approveBtnTxt := slack.NewTextBlockObject("plain_text", "Authorized", false, false)
	// approveBtn := slack.NewButtonBlockElement("", "click_me_123", approveBtnTxt)

	// denyBtnTxt := slack.NewTextBlockObject("plain_text", "Unauthorized", false, false)
	// denyBtn := slack.NewButtonBlockElement("", "click_me_123", denyBtnTxt)

	// actionBlock := slack.NewActionBlock("", approveBtn, denyBtn)

	// Build Message with blocks created above
	_, _, err := client.PostMessage(s.ChannelID, slack.MsgOptionBlocks(
		headerSection, fieldsSection, tableSection,
	))

	return err
}

func New(token string) *slack.Client {
	return slack.New(token)
}
