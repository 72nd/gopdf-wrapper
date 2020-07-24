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

func loadFromRice(filename, description string) ([]byte, error) {
	box, err := riceBox()
	if err != nil {
		return nil, err
	}
	data, err := box.Bytes(filename)
	if err != nil {
		return nil, fmt.Errorf("could not load %s: %s", description, err)
	}
	return data, nil

}

func riceBox() (*rice.Box, error) {
	box, err := rice.FindBox("utils")
	if err != nil {
		return nil, fmt.Errorf("rice find box failed: %s", err)
	}
	return box, nil
}

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
	return loadFromRice("Lato-Heavy.ttf", "lato heavy")
}

// LatoRegular returns the regular style of the Lato font.
func LatoRegular() ([]byte, error) {
	return loadFromRice("Lato-Regular.ttf", "lato heavy")
}
