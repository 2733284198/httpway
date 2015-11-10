package httpway

import "net/http"

//session interface
type Session interface {
	IsAuth() bool

	Set(key string, val interface{})

	Get(key string) interface{}
	GetInt(key string) int
	GetString(key string) string
}

//sessions manager interface
type SessionManager interface {
	Get(w http.ResponseWriter, r *http.Request) Session
	Set(w http.ResponseWriter, r *http.Request, session Session)
}