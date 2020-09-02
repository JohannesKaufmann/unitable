package main

import (
	"github.com/JohannesKaufmann/unitable"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	filepath := "simple.pdf"

	if err := run(filepath); err != nil {
		panic(err)
	}
}

func run(filepath string) error {
	creatr := creator.New()
	creatr.SetPageMargins(50, 50, 50, 50)

	// our colors
	lightYellow := creator.ColorRGBFromHex("#faefc4")
	lightBlue := creator.ColorRGBFromHex("#e8fbfe")

	// the default style (also used by header and price)
	defaultStyle := unitable.Style{
		Font:     "regular",
		FontSize: 12,

		Border: []unitable.BorderStyle{
			{
				Side:  creator.CellBorderSideAll,
				Style: creator.CellBorderStyleSingle,
				Width: 1.5,
				Color: lightBlue,
			},
		},
	}
	cell := unitable.NewCellBase(defaultStyle)

	// the header style
	headerStyle := unitable.Style{
		Font:            "bold",
		HorizontalAlign: creator.CellHorizontalAlignmentCenter,
		BackgroundColor: lightYellow,
	}
	header := unitable.NewCellBase(defaultStyle, headerStyle)

	// the price style
	priceStyle := unitable.Style{
		Font:            "regular",
		HorizontalAlign: creator.CellHorizontalAlignmentRight,
	}
	price := unitable.NewCellBase(defaultStyle, priceStyle)

	// and now assemble the data with our new helper functions
	data := [][]unitable.Renderer{
		{
			header("Name"),
			header("Count"),
			// override the header style later on
			header(
				"Price",
				unitable.Style{
					BackgroundColor: lightBlue,
				},
			),
		},
		{
			cell("Item A"),
			cell("2"),
			price("16,00 €"),
		},
		{
			cell("Item B"),
			cell("1"),
			price("8,00 €"),
		},
		{
			price("24,00 €", unitable.Style{ColSpan: 3}),
		},
	}

	fonts, err := getFonts()
	if err != nil {
		return err
	}

	table := unitable.NewTable(creatr, fonts, data)

	// Do your work on the table & draw it
	table.SetMargins(0, 0, 20, 0)
	if err := creatr.Draw(table); err != nil {
		return err
	}

	// Write to output file.
	return creatr.WriteToFile(filepath)
}

func getFonts() (map[string]*model.PdfFont, error) {
	// Create fonts.
	fontRegular, err := model.NewStandard14Font("Helvetica")
	if err != nil {
		return nil, err
	}

	fontBold, err := model.NewStandard14Font("Helvetica-Bold")
	if err != nil {
		return nil, err
	}

	fonts := map[string]*model.PdfFont{
		"regular": fontRegular,
		"bold":    fontBold,
	}

	return fonts, nil
}
