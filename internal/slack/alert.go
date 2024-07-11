package slack

import "github.com/slack-go/slack"

func Send(client *slack.Client) error {
	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", "Snowguard flagged the following activity associated with your Snowflake account.", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Fields
	typeField := slack.NewTextBlockObject("mrkdwn", "*Category:*\n${SNOWFLAKE_ACTIVITY}", false, false)
	whenField := slack.NewTextBlockObject("mrkdwn", "*When:*\n${TIMESTAMP}", false, false)
	lastUpdateField := slack.NewTextBlockObject("mrkdwn", "*Snowflake Account:*\n${SNOWFLAKE_ACCOUNT}", false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, typeField)
	fieldSlice = append(fieldSlice, whenField)
	fieldSlice = append(fieldSlice, lastUpdateField)

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	// Approve and Deny Buttons
	// approveBtnTxt := slack.NewTextBlockObject("plain_text", "Authorized", false, false)
	// approveBtn := slack.NewButtonBlockElement("", "click_me_123", approveBtnTxt)

	// denyBtnTxt := slack.NewTextBlockObject("plain_text", "Unauthorized", false, false)
	// denyBtn := slack.NewButtonBlockElement("", "click_me_123", denyBtnTxt)

	// actionBlock := slack.NewActionBlock("", approveBtn, denyBtn)

	// Build Message with blocks created above
	_, _, err := client.PostMessage("C07BEK4BDGE", slack.MsgOptionBlocks(
		headerSection, fieldsSection,
	))

	return err
}

func New(token string) *slack.Client {
	return slack.New(token)
}
