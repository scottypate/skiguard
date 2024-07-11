package users

import "time"

type Users struct {
	LoginName           string    `json:"login_name"`
	CreatedOn           time.Time `json:"created_on"`
	Email               string    `json:"email"`
	HasPassword         bool      `json:"has_password"`
	Disabled            bool      `json:"disabled"`
	LastSuccessLogin    time.Time `json:"last_success_login"`
	PasswordLastSetTime time.Time `json:"password_last_set_time"`
}
