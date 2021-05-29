package test

import "github.com/rs/xid"

func NewRandomID() string {
	return xid.New().String()
}
