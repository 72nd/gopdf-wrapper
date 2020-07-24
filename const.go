package gopdf_wrapper

import "github.com/signintech/gopdf"

// FontStyle wraps the gopdf font style constants into a type.
type FontStyle int

const (
	NormalStyle    FontStyle = gopdf.Regular
	ItalicStyle              = gopdf.Italic
	BoldStyle                = gopdf.Bold
	UnderlineStyle           = gopdf.Underline
)

// fontFunction is a function which returns a font as bytes and an optional error.
type fontFunction func() ([]byte, error)

// LineStyle wraps the gopdf line styles into a type.
type LineStyle string

const (
	SolidLine  LineStyle = ""
	DashedLine           = "dashed"
	DottedLine           = "dotted"
)
