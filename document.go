// gopdf-wrapper is a simple wrapper around the [gopdf](https://github.com/signintech/gopdf) library
// aiming to simplify recurring tasks of creating PDF's with gopdf.
package gopdf_wrapper

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/signintech/gopdf"
	"github.com/signintech/gopdf/fontmaker/core"
)

// Doc is the basic structure for a PDF file.
type Doc struct {
	gopdf.GoPdf
	fontSize        int
	defaultFontSize int
	lineSpread      float64
	capValue        float64
	fontStyle       string
	currentX        float64
	currentY        float64
}

// NewDoc returns a new Doc.
func NewDoc(fontSize int, lineSpread float64) (*Doc, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4, Unit: gopdf.Unit_MM})

	doc := Doc{
		GoPdf:           pdf,
		fontSize:        fontSize,
		defaultFontSize: fontSize,
		lineSpread:      lineSpread,
		fontStyle:       "",
		currentX:        0,
		currentY:        0,
	}

	latoRegular, err := LatoRegular()
	if err != nil {
		return nil, err
	}
	doc.AddFont(latoRegular, "regular", NormalStyle, true)

	latoHeavy, err := LatoHeavy()
	if err != nil {
		return nil, err
	}
	doc.AddFont(latoHeavy, "bold", BoldStyle, false)

	var parser core.TTFParser
	if err := parser.ParseByReader(bytes.NewBuffer(latoRegular)); err != nil {
		return nil, fmt.Errorf("error while parsing font for height calculation: %s", err)
	}
	doc.capValue = float64(parser.CapHeight()) * 1000.0 / float64(parser.UnitsPerEm())
	doc.SetFontSize(fontSize)
	return &doc, nil
}

// AddFont adds a font to a document.
func (d *Doc) AddFont(font []byte, name string, fontStyle FontStyle, useKerning bool) error {
	style := int(fontStyle)
	err := d.GoPdf.AddTTFFontByReaderWithOption(name, bytes.NewBuffer(font), gopdf.TtfOption{Style: style, UseKerning: useKerning})
	if err != nil {
		return err
	}
	return nil
}

// SetPosition encapsulates the SetX and SetY methods of gopdf. New elements will be added
// at the currently set position.
func (d *Doc) SetPosition(x, y float64) {
	d.GoPdf.SetX(x)
	d.currentX = x
	d.GoPdf.SetY(y)
	d.currentY = y
}

// SetFontSize sets the font size for all elements added after.
func (d *Doc) SetFontSize(size int) error {
	d.fontSize = size
	if err := d.GoPdf.SetFont("lato", "", size); err != nil {
		return fmt.Errorf("error while changing PDF font size: %s", err)
	}
	return nil
}

// DefaultFontSize resets the font size to the initial default.
func (d *Doc) DefaultFontSize() {
	d.SetFontSize(d.defaultFontSize)
}

// SetFontStyle changes the font style (italic, bold...) for elements added afterwards.
func (d *Doc) SetFontStyle(style string) error {
	if err := d.GoPdf.SetFont("lato", style, d.fontSize); err != nil {
		return fmt.Errorf("error while changing PDF font style: %s", err)
	}
	d.fontStyle = style
	return nil
}

// DefaultFontStyle resets the font style to normal style.
func (d *Doc) DefaultFontStyle() {
	d.SetFontStyle("")
}

// AddText adds a text field at the given position.
func (d *Doc) AddText(x, y float64, content string) error {
	d.SetPosition(x, y)
	if err := d.GoPdf.Cell(nil, content); err != nil {
		return fmt.Errorf("error adding text to PDF: %s", err)
	}
	return nil
}

// AddFormattedText adds a text field at the given position with individual size and style.
func (d *Doc) AddFormattedText(x, y float64, content string, size int, style string) {
	d.SetFontSize(size)
	d.SetFontStyle(style)
	d.AddText(x, y, content)
	d.DefaultFontSize()
	d.DefaultFontStyle()
}

// AddMultilineText adds a text field with multiple lines of text. The lines are
// divided by the new-line character `\n`.
func (d *Doc) AddMultilineText(x, y float64, content string) {
	data := strings.Split(content, "\n")
	for i := range data {
		d.AddText(x, y, data[i])
		y += d.LineHeight()
	}
}

// AddFormattedMultilineText has the same functionality as AddMultilineText but
// with a individual font-size and style.
func (d *Doc) AddFormattedMultilineText(x, y float64, content string, size int, style string) {
	d.SetFontSize(size)
	d.SetFontStyle(style)
	data := strings.Split(content, "\n")
	for i := range data {
		d.AddText(x, y, data[i])
		y += d.LineHeight()
	}
	d.DefaultFontSize()
	d.DefaultFontStyle()
}

// AddLine adds a line to the Document.
func (d *Doc) AddLine(x1, y1, x2, y2, width float64, lineStyle LineStyle) {
	d.GoPdf.SetLineWidth(width)
	style := string(lineStyle)
	d.GoPdf.SetLineType(style)
	d.GoPdf.Line(x1, y1, x2, y2)
}

// LineHeight calculates and returns the line height.
func (d Doc) LineHeight() float64 {
	return d.capValue * float64(d.fontSize) / 2000.0 * d.lineSpread
}

func (d Doc) textLineWidth() float64 {
	return gopdf.PageSizeA4.W - d.GoPdf.MarginLeft() - d.GoPdf.MarginRight()
}
