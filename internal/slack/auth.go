package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

func AuthVerity(token string) error {
	client := slack.New(token)
	// Header Section
	_, err := client.AuthTest()

	if err != nil {
		return fmt.Errorf("error verifying slack token: %v", err)
	}

	return nil
}
