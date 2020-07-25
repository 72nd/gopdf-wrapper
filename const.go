package gopdfwrapper

import "github.com/signintech/gopdf"

// FontStyle wraps the gopdf font style constants into a type.
type FontStyle int

const (
	// NormalStyle is the normal font style aka. regular style.
	NormalStyle FontStyle = gopdf.Regular
	// ItalicStyle represents the italic font style.
	ItalicStyle = gopdf.Italic
	// BoldStyle represents the bold font style.
	BoldStyle = gopdf.Bold
	// UnderlineStyle represents the underline font style.
	UnderlineStyle = gopdf.Underline
)

// fontFunction is a function which returns a font as bytes and an optional error.
type fontFunction func() ([]byte, error)

// LineStyle wraps the gopdf line styles into a type.
type LineStyle string

const (
	// SolidLine is a solid line.
	SolidLine LineStyle = ""
	// DashedLine is a dashed line.
	DashedLine = "dashed"
	// DottedLine is a dotted line.
	DottedLine = "dotted"
)
