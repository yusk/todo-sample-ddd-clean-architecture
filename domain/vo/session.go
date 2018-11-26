package vo

import "encoding/gob"

type Session struct {
	UserID uint `json:"user_id"`
}

func init() {
	gob.Register(Session{})
}
