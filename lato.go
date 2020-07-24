// fonts package provides the embedded [Lato fonts](https://www.latofonts.com/).
package gopdf_wrapper

import (
	"fmt"

	rice "github.com/GeertJohan/go.rice"
)

func riceBox() (*rice.Box, error) {
	box, err := rice.FindBox("utils")
	if err != nil {
		return nil, fmt.Errorf("rice find box failed: %s", err)
	}
	return box, nil
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
