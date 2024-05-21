package models

import "time"

type User struct {
	ID            int     `db:"user_id"`
	First_name    *string `db:"first_name"`
	Last_name     *string `db:"last_name"`
	Password      *string
	Email         *string `db:"email"`
	Phone         *string
	Token         *string
	User_type     *string
	Refresh_token *string
	Created_at    time.Time
	Updated_at    time.Time
	User_id       string
}
