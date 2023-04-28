package rgeohex_test

import (
	"errors"
	"testing"

	"github.com/ypm-llc/go-geohex-convert-address-jp/rgeohex"
)

func TestCodeToAddressName(t *testing.T) {
	// Test case 1: valid geohex code
	expectedName := "東京都目黒区駒場一丁目"
	actualName, err := rgeohex.CodeToAddressName("XM488542518")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if actualName != expectedName {
		t.Errorf("expected %s but got %s", expectedName, actualName)
	}

	// Test case 2: invalid geohex code
	expectedErr := errors.New("expected a digit, got '%!s(int32=118)'")
	_, err = rgeohex.CodeToAddressName("invalid code")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("expected %s but got %s", expectedErr.Error(), err.Error())
	}
}
