package TestingUtils

import "testing"

func CheckEquals(excpected int, actual int, additionalMessge string,  t *testing.T) {
	if excpected != actual {
		t.Errorf("expected = %d, actual = %d:" + additionalMessge, excpected, actual)
	}
}
