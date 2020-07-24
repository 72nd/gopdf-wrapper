// fonts contains the embedded fonts and utility functions.
package fonts

import (
	"fmt"

	rice "github.com/GeertJohan/go.rice"
)

// FontFamily encapsulates a font family for easy adding to the document
type FontFamily struct {
	Name   string
	Bold   []byte
	Italic []byte
	Normal []byte
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

// NewLatoFamily returns a new FontFamily of the Lato font.
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
		Bold:   heavy,
		Normal: normal,
	}, nil
}

// LatoHeavy returns the heavy style of the Lato font.
func LatoHeavy() ([]byte, error) {
	return loadFromRice("Lato-Heavy.ttf", "lato heavy")
}

// LatoRegular returns the regular style of the Lato font.
func LatoRegular() ([]byte, error) {
	return loadFromRice("Lato-Regular.ttf", "lato heavy")
}

// NewLiberationSansFamily returns a new FontFamily of the Liberation Sans font.
func NewLiberationSansFamily() (*FontFamily, error) {
	bold, err := LiberationSansBold()
	if err != nil {
		return nil, err
	}
	italic, err := LiberationSansItalic()
	if err != nil {
		return nil, err
	}
	normal, err := LiberationSansRegular()
	if err != nil {
		return nil, err
	}
	return &FontFamily{
		Name:   "liberation-sans",
		Bold:   bold,
		Italic: italic,
		Normal: normal,
	}, nil
}

// LiberationSansBold returns the bold style of the Liberation Sans font.
func LiberationSansBold() ([]byte, error) {
	return loadFromRice("LiberationSans-Bold.ttf", "liberations sans bold")
}

// LiberationSansItalic returns the italic style of the Liberation Sans font.
func LiberationSansItalic() ([]byte, error) {
	return loadFromRice("LiberationSans-Italic.ttf", "liberations sans italic")
}

// LiberationSansRegular returns the normal style of the Liberation Sans font.
func LiberationSansRegular() ([]byte, error) {
	return loadFromRice("LiberationSans-Bold.ttf", "liberations sans regular")
}
