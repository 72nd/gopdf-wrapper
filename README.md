# gopdf-wrapper

The [gopdf](https://github.com/signintech/gopdf) library is a great library for creating PDFs in golang. This wrapper provides some convince abstractions for recurring tasks. Like:

- Maintain a default font size
- Add formatted text
- Add multiline text
- gopdf-wrapper embeds the [Lato font](https://www.latofonts.com/) and [Liberation Sans](https://github.com/liberationfonts/liberation-fonts) which is probably enough for simple reports

The code of this library is licensed under the [MIT License]() the fonts (Lato and Liberation Sans) on the other hand are licensed under the [SIL Open Font License (OFL)](http://scripts.sil.org/cms/scripts/page.php?site_id=nrsi&id=OFL). 

## Example

```golang
// New document with font size 12 and line spread 1. 
doc, err := NewDoc(12, 1)
if err != nil {
	t.Error(err)
}
liberation, err := fonts.NewLiberationSansFamily()
if err != nil {
	t.Error(err)
}
doc.SetFontFamily(*liberation)
doc.AddPage()

// Text
doc.AddFormattedText(10, 20, "This is a document", 25, "bold")
doc.AddSizedText(10, 20 + doc.LineHeight(25), "Some subtitle", 18)
doc.AddText(10, 40, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor.")
doc.AddMultilineText(10, 50, "Ut enim ad minim veniam,\nquis nostrud exercitation ullamco laboris\nnisi ut aliquip ex ea commodo consequat.\nDuis aute irure dolor in reprehenderit\nin voluptate velit esse cillum dolore\neu fugiat nulla pariatur.")
doc.AddWrapText(10, 80, 140, "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem.")

// Lines
doc.AddLine(10, 140, 30, 160, 0.1, SolidLine)
doc.AddLine(10, 160, 30, 140, 0.1, SolidLine)
doc.AddLine(10, 150, 30, 150, 0.1, SolidLine)
doc.AddLine(20, 140, 20, 160, 0.1, SolidLine)

// Write document to PDF.
doc.WritePdf("document.pdf")
```
