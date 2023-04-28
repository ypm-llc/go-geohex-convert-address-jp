# ypm-llc/go-geohex-convert-address-jp

Japan address mutual convert libraries (among Japan address name and GeoHex code)

## Synopsis

### GeoHex code to Japan address name

    import (
        "fmt"
        "github.com.ypm-llc/go-geohex-convert-address-jp/rgeohex"
    )
    func main () {
        geohexCode = "XM488542518"
	    actualName, err := rgeohex.CodeToAddressName(geohexCode)
        if err != nil {
            panic err
        }
        fmt.printf("actualName = %s\n", actualName) // it says "東京都目黒区駒場一丁目"
    }

### Japan address name to GeoHex code

    import (
        "fmt"
        "github.com.ypm-llc/go-geohex-convert-address-jp/addrgeohex"
    )
    func main () {
        addressName := "東京都目黒区駒場一丁目"
        geohexLevel := 7
        actualCode, err := addrgeohex.AddressNameToCode(addressName, geohexLevel)
        if err != nil {
            panic err
        }
        fmt.printf("actualCode = %s\n", actualCode) // it says "XM4885425"
    }

## See also

- http://www.geohex.org/ GeoHex documentation
- https://memo.appri.me/programming/gsi-geocoding-api 国土地理院APIでお手軽ジオコーディング＆逆ジオコーディング - 芽萌丸