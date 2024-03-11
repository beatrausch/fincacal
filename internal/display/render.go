package display

import (
	"fmt"
	"github.com/beatrausch/fincacal/internal/finca"
	"github.com/pterm/pterm"
)

func RenderPriceTable(table [][]string) {
	pterm.DefaultTable.WithHasHeader().WithData(table).Render()
}

func RenderOverview(overall finca.Summary) {
	pterm.DefaultTable.WithData([][]string{
		{"Attendees", fmt.Sprintf("%d", overall.Attendees)},
		{"Overall Nights", fmt.Sprintf("%d", overall.OverallNights)},
		{"Overall Price", fmt.Sprintf("%.2f€", overall.OverallPrice)},
		{"Price per Night", fmt.Sprintf("%.2f€", overall.PricePerNight)},
	}).Render()
}

func RenderNewLine() {
	pterm.DefaultBasicText.Println()
}
