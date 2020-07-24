// fonts contains the embedded fonts and utility functions.
package fonts

import (
	"fmt"

	rice "github.com/GeertJohan/go.rice"
)

// FontFamily encapsulates a font family for easy adding to the document
type FontFamily struct {
	Name   string
	Normal []byte
	Bold   []byte
}

func riceBox() (*rice.Box, error) {
	box, err := rice.FindBox("utils")
	if err != nil {
		return nil, fmt.Errorf("rice find box failed: %s", err)
	}
	return box, nil
}
