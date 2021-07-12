package api

import "time"

type Api struct {
	Url string
	Query string
	Timeout time.Duration
	Lang string
	Addr string
	Password string
	Db int
}

