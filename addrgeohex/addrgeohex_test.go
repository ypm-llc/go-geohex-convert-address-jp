package addrgeohex_test

import (
	"errors"
	"testing"

	"github.com/ypm-llc/go-geohex-convert-address-jp/addrgeohex"
)

func TestAddressNameToCode(t *testing.T) {
	// Test case 1: valid address name
	expectedCode := "XM4885425"
	actualCode, err := addrgeohex.AddressNameToCode("東京都目黒区駒場一丁目", 7)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
	if actualCode != expectedCode {
		t.Errorf("expected %s but got %s", expectedCode, actualCode)
	}

	// Test case 2: invalid address name
	expectedErr := errors.New("no features found")
	_, err = addrgeohex.AddressNameToCode("invalid address name", 6)
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("expected %s but got %s", expectedErr.Error(), err.Error())
	}
}
