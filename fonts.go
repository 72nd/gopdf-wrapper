package gopdf_wrapper

import "github.com/signintech/gopdf"

// FontStyle wraps the gopdf font style constants into a type.
type FontStyle int

const (
	Bold FontStyle = gopdf.Bold
)

// FontFunction is a function which returns a font as bytes and an optional error.
type FontFunction func() ([]byte, error)

