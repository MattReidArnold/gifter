package id

import "github.com/rs/xid"

func GenerateId() (string, error) {
	id := xid.New()
	return id.String(), nil
}
