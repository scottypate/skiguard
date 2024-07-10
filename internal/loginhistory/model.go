package loginhistory

import (
	"time"
)

type LoginHistory struct {
	EventId                    int64     `json:"event_id"`
	EventTimestamp             time.Time `json:"event_timestamp"`
	EventType                  string    `json:"event_type"`
	UserName                   string    `json:"user_name"`
	ClientIp                   string    `json:"client_ip"`
	ReportedClientType         string    `json:"reported_client_type"`
	ReportedClientVersion      string    `json:"reported_client_version"`
	FirstAuthenticationFactor  string    `json:"first_authentication_factor"`
	SecondAuthenticationFactor string    `json:"second_authentication_factor"`
	IsSuccess                  string    `json:"is_success"`
	ErrorCode                  int64     `json:"error_code"`
	ErrorMessage               string    `json:"error_message"`
	RelatedEventId             int64     `json:"related_event_id"`
	Connection                 string    `json:"connection"`
}
