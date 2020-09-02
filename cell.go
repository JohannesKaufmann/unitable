package unitable

import (
	"github.com/unidoc/unipdf/v3/contentstream/draw"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type Cell struct {
	TextContent string

	Content creator.VectorDrawable
	Style   *Style
}

func (element *Cell) ColSpan() int {
	if element.Style.ColSpan == 0 {
		return 1
	}
	return element.Style.ColSpan
}

func (element *Cell) getContent(c *creator.Creator, table *creator.Table, fonts map[string]*model.PdfFont) creator.VectorDrawable {
	// if the user supplied content already, use it
	if element.Content != nil {
		return element.Content
	}

	style := element.Style

	p := c.NewParagraph(element.TextContent)
	p.SetFont(fonts[style.Font])
	if style.FontSize != 0 {
		p.SetFontSize(style.FontSize)
	}
	if style.Color != nil {
		p.SetColor(style.Color)
	}

	return p
}
func (element *Cell) Render(c *creator.Creator, table *creator.Table, fonts map[string]*model.PdfFont) {
	style := element.Style

	cell := table.MultiColCell(style.ColSpan)
	if style.BackgroundColor != nil {
		cell.SetBackgroundColor(style.BackgroundColor)
	}

	for _, border := range style.Border {
		cell.SetBorder(border.Side, border.Style, border.Width)
		cell.SetBorderColor(border.Color)
		cell.SetBorderLineStyle(border.LineStyle)
	}

	cell.SetHorizontalAlignment(style.HorizontalAlign)
	cell.SetVerticalAlignment(style.VerticalAlign)

	cell.SetContent(element.getContent(c, table, fonts))
}

type BorderStyle struct {
	Side  creator.CellBorderSide
	Style creator.CellBorderStyle
	Width float64

	// Unfortunately the `Color` and `LineStyle` can't be applied
	// to individual sides. Instead you can only change the color
	// for all sides at once.
	Color     creator.Color
	LineStyle draw.LineStyle
}
type Style struct {
	Border []BorderStyle

	Color           creator.Color
	BackgroundColor creator.Color

	ColSpan int

	Font     string
	FontSize float64

	HorizontalAlign creator.CellHorizontalAlignment
	VerticalAlign   creator.CellVerticalAlignment
}
