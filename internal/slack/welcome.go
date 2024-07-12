package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func SendWelcomeMessage(channelId string, token string) error {
	client := slack.New(token)
	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf(
		":information_source: Snowguard has been installed to post alerts to this channel.\n This app will monitor the `%s` Snowflake account for anomalies in login activity and data copying.", os.Getenv("SNOWFLAKE_ACCOUNT"),
	), false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	_, _, err := client.PostMessage(channelId, slack.MsgOptionBlocks(
		headerSection,
	))

	return err
}
