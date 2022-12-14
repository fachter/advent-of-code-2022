package adventOfCode

import "testing"

func AssertExpectedOrFail(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("FAILED. Expected %s, got %s", expected, actual)
	}
}
