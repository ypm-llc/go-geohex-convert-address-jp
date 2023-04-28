/*
Japan reverse geocoding library for geohex
*/
package rgeohex

import (
	"github.com/bsm/go-geohex/v3"
	"github.com/ypm-llc/go-msearch-gsi-jp/mreversegeocode"
)

// convert geohex code to address name
func CodeToAddressName(code string) (string, error) {
	ll, err := geohex.Decode(code)
	if err != nil {
		return "", err
	}

	return mreversegeocode.LatLonToAddressName(ll.Lat, ll.Lon)
}
