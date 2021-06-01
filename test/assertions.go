package test

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func FailAssertion(t *testing.T, got, want interface{}, args ...string) {
	t.Helper()

	t.Errorf("assertion failed: %s\n\n\t\tgot:\n\t\t\t%v\n\n\t\twant:\n\t\t\t%v", strings.Join(args, " "), got, want)
}

func AssertEqual(t *testing.T, got, want interface{}, args ...string) {
	t.Helper()

	if got != want {
		FailAssertion(t, got, want, append(args, "AssertEqual")...)
	}
}

func AssertTypeEqual(t *testing.T, got, want interface{}, args ...string) {
	t.Helper()
	gotType := reflect.TypeOf(got)
	wantType := reflect.TypeOf(want)
	if gotType != wantType {
		FailAssertion(t, got, want, append(args, "AssertTypeEqual")...)
	}
}

func AssertErrorEqual(t *testing.T, got, want error, args ...string) {
	t.Helper()
	if got == nil || got.Error() != want.Error() {
		FailAssertion(t, got, want, append(args, "AssertErrorEqual")...)
	}
}

func AssertErrorIs(t *testing.T, got, want error, args ...string) {
	t.Helper()
	if got == nil || !errors.Is(got, want) {
		FailAssertion(t, got, want, append(args, "AssertErrorIs")...)
	}
}

func AssertNil(t *testing.T, got interface{}, args ...string) {
	t.Helper()
	if got != nil {
		FailAssertion(t, got, nil, append(args, "AssertNil")...)
	}
}

func AssertTrue(t *testing.T, got bool, args ...string) {
	t.Helper()
	if !got {
		FailAssertion(t, got, true, append(args, "AssertOk")...)
	}
}

func AssertFalse(t *testing.T, got bool, args ...string) {
	t.Helper()
	if got {
		FailAssertion(t, got, false, append(args, "AssertOk")...)
	}
}

func AssertDeepEqual(t *testing.T, got, want interface{}, args ...string) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		FailAssertion(t, got, want, append(args, "AssertDeepEqual")...)
	}
}
