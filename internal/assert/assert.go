package assert

import "testing"

func Equal[T comparable](t *testing.T, a, b T) {
	t.Helper()

	if a != b {
		t.Errorf("expected `%v` to equal `%v`", a, b)
	}
}

func True(t *testing.T, b bool) {
	t.Helper()

	if !b {
		t.Error("expected `false` to be `true`")
	}
}

func False(t *testing.T, b bool) {
	t.Helper()

	if b {
		t.Error("expected `true` to be `false`")
	}
}

func Nil(t *testing.T, v interface{}) {
	t.Helper()

	if v != nil {
		t.Errorf("expected `%v` to equal `nil`", v)
	}
}

func NotNil(t *testing.T, v interface{}) {
	t.Helper()

	if v == nil {
		t.Error("expected `nil` not to equal `nil`")
	}
}
