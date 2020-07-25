// Package gopdfwrapper is a simple wrapper around the [gopdf](https://github.com/signintech/gopdf) library
// aiming to simplify recurring tasks of creating PDF's with gopdf.
package gopdfwrapper

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/72nd/gopdf-wrapper/fonts"
	"github.com/signintech/gopdf"
	"github.com/signintech/gopdf/fontmaker/core"
)

// Doc is the basic structure for a PDF file.
type Doc struct {
	gopdf.GoPdf
	fontSize        int
	fontName        string
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

	latoFamily, err := fonts.NewLatoFamily()
	if err != nil {
		return nil, err
	}
	doc.SetFontFamily(*latoFamily)

	doc.SetFontSize(fontSize)
	return &doc, nil
}

// SetFontFamily sets the used font family.
func (d *Doc) SetFontFamily(family fonts.FontFamily) error {
	d.setFont(family.Bold, family.Name, BoldStyle, true)
	d.setFont(family.Italic, family.Name, ItalicStyle, true)
	d.setFont(family.Normal, family.Name, NormalStyle, true)

	var parser core.TTFParser
	if err := parser.ParseByReader(bytes.NewBuffer(family.Normal)); err != nil {
		return fmt.Errorf("error while parsing font for height calculation: %s", err)
	}
	d.capValue = float64(parser.CapHeight()) * 1000.0 / float64(parser.UnitsPerEm())
	d.fontName = family.Name
	return nil
}

// setFont adds a font to a document.
func (d *Doc) setFont(font []byte, name string, fontStyle FontStyle, useKerning bool) error {
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
	if err := d.GoPdf.SetFont(d.fontName, "", size); err != nil {
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
	if err := d.GoPdf.SetFont(d.fontName, style, d.fontSize); err != nil {
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

// AddSizedText adds a text field at the given position with the given size.
func (d *Doc) AddSizedText(x, y float64, content string, size int) {
	d.SetFontSize(size)
	d.AddText(x, y, content)
	d.DefaultFontSize()
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
		y += d.DefaultLineHeight()
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
		y += d.DefaultLineHeight()
	}
	d.DefaultFontSize()
	d.DefaultFontStyle()
}

// AddWrapText automatically wraps the text when the line reaches x2.
func (d *Doc) AddWrapText(x1, y, x2 float64, content string) {
	width := x2 - x1
	chars := []rune(content)
	lines := 0.0
	var i, j int
	for j = 0; j < len(chars); j++ {
		l, _ := d.GoPdf.MeasureTextWidth(string(chars[i:j]))
		if l >= width {
			d.AddText(x1, y+d.LineHeight(d.defaultFontSize)*lines, string(chars[i:j-1]))
			i = j - 1
			lines++
		}
	}
	// fmt.Println(string(chars[]))
}

// AddFormattedWrapText automatically wraps the formatted text when the line reaches x2.
func (d *Doc) AddFormattedWrapText(x1, y, x2 float64, content string, size int, style string) {
	d.SetFontSize(size)
	d.SetFontStyle(style)
	d.AddWrapText(x1, y, x2, content)
	d.DefaultFontStyle()
	d.DefaultFontSize()
}

// AddLine adds a line to the Document.
func (d *Doc) AddLine(x1, y1, x2, y2, width float64, lineStyle LineStyle) {
	d.GoPdf.SetLineWidth(width)
	style := string(lineStyle)
	d.GoPdf.SetLineType(style)
	d.GoPdf.Line(x1, y1, x2, y2)
}

// DefaultLineHeight calculates and returns the line height.
func (d Doc) DefaultLineHeight() float64 {
	return d.capValue * float64(d.fontSize) / 2000.0 * d.lineSpread
}

// LineHeight returns the line height of a text line with a given height in mm.
func (d Doc) LineHeight(fontSize int) float64 {
	return d.capValue * float64(fontSize) / 2000.0 * d.lineSpread
}

func (d Doc) textLineWidth() float64 {
	return gopdf.PageSizeA4.W - d.GoPdf.MarginLeft() - d.GoPdf.MarginRight()
}
