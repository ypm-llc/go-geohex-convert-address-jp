/*
Japan address convert library to geohex code
*/
package addrgeohex

import (
	"fmt"

	"github.com/bsm/go-geohex/v3"
	"github.com/ypm-llc/go-msearch-gsi-jp/msearch"
)

// convert address name to geohex code
func AddressNameToCode(address string, level int) (string, error) {
	features, err := msearch.ContainsSearchAddress(address)
	if err != nil {
		return "", err
	}

	if len(features) == 0 {
		return "", fmt.Errorf("no features found")
	}

	lat := features[0].Geometry.Coordinates[1]
	lon := features[0].Geometry.Coordinates[0]
	gh, err := geohex.Encode(lat, lon, level)
	if err != nil {
		return "", err
	}

	return gh.Code, nil
}
