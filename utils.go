package unitable

import (
	"fmt"

	"github.com/imdario/mergo"
)

// merge merges the different styles.
func merge(styles []Style) *Style {
	var base Style

	for _, style := range styles {
		if err := mergo.Merge(&base, style, mergo.WithOverride); err != nil {
			panic(err)
		}
	}
	return &base
}

func getRowWidth(cells []Renderer) int {
	var totalWidth int

	for _, element := range cells {
		totalWidth += element.ColSpan()
	}
	return totalWidth
}

func GetMaxRowWidth(tableData [][]Renderer) int {
	var values []int
	for _, row := range tableData {
		values = append(values, getRowWidth(row))
	}

	min, max := findMinAndMax(values)
	if min != max {
		fmt.Println("min != max")
	}

	return max
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
