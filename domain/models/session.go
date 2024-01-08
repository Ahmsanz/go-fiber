package models

import "time"

type Session struct {
	Username   string        `json:"username"`
	Valid      bool          `json:"isValid"`
	Token      string        `json:"token"`
	Expiration time.Duration `json:"exp"`
}

func BuildSession(username string, token string, expiration time.Duration) *Session {
	return &Session{Username: username, Valid: true, Token: token, Expiration: expiration}
}
