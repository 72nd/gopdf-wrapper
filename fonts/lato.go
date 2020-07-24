// fonts package provides the embedded [Lato fonts](https://www.latofonts.com/).
package fonts

import (
	"fmt"
)

func NewLatoFamily() (*FontFamily, error) {
	heavy, err := LatoHeavy()
	if err != nil {
		return nil, err
	}
	normal, err := LatoRegular()
	if err != nil {
		return nil, err
	}
	return &FontFamily{
		Name:   "lato",
		Normal: normal,
		Bold:   heavy}, nil
}

// LatoHeavy returns the heavy style of the Lato font.
func LatoHeavy() ([]byte, error) {
	box, err := riceBox()
	if err != nil {
		return nil, err
	}
	data, err := box.Bytes("Lato-Heavy.ttf")
	if err != nil {
		return nil, fmt.Errorf("could not load lato heavy: %s", err)
	}
	return data, nil
}

// LatoRegular returns the regular style of the Lato font.
func LatoRegular() ([]byte, error) {
	box, err := riceBox()
	if err != nil {
		return nil, err
	}
	data, err := box.Bytes("Lato-Regular.ttf")
	if err != nil {
		return nil, fmt.Errorf("could not load lato regular: %s", err)
	}
	return data, nil
}
