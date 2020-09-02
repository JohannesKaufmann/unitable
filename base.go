package unitable

type CellFunc func(content string, overrideStyles ...Style) *Cell

// NewCellBase returns a helper function that can be used
// repeatedly to create a new cell with the styles applied.
func NewCellBase(defaultStyles ...Style) CellFunc {
	return func(content string, overrideStyles ...Style) *Cell {
		style := merge(append(defaultStyles, overrideStyles...))

		return &Cell{
			TextContent: content,
			Style:       style,
		}
	}
}
