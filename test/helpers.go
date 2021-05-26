package test

import (
	"reflect"
	"testing"

	"github.com/rs/xid"
)

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()

	gotType := reflect.TypeOf(got)
	wantType := reflect.TypeOf(want)
	if gotType != wantType {
		t.Errorf("got type: %v, want type: %v", gotType, wantType)
	}
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func NewRandomID() string {
	return xid.New().String()
}
