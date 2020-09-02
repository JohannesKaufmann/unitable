package unitable

import (
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type Renderer interface {
	Render(c *creator.Creator, table *creator.Table, fonts map[string]*model.PdfFont)

	ColSpan() int
}

// NewTable calculates how many columns are needed and
// adds all the cells from `tableData`.
// It returns a `creator.Table` which you still have to draw.
func NewTable(c *creator.Creator, fonts map[string]*model.PdfFont, tableData [][]Renderer) *creator.Table {
	maxRowWidth := GetMaxRowWidth(tableData)

	// Create table.
	table := c.NewTable(maxRowWidth)

	for _, column := range tableData {
		for _, element := range column {
			element.Render(c, table, fonts)
		}

		rowWidth := getRowWidth(column)

		// if there are cells missing, fill the area with empty ones
		table.SkipCells(maxRowWidth - rowWidth)
	}

	return table
}
