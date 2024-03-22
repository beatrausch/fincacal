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
		{"Deposit", fmt.Sprintf("%.2f€", overall.Deposit)},
		{"Overall Price", fmt.Sprintf("%.2f€", overall.OverallPrice)},
		{"Remainder", fmt.Sprintf("%.2f€", overall.Remainder)},
		{"Price per Night", fmt.Sprintf("%.2f€", overall.PricePerNight)},
		{"Check", status(overall.ValidatePrice())},
	}).Render()
}

func RenderNewLine() {
	pterm.DefaultBasicText.Println()
}

func status(ok bool) string {
	if ok {
		return "\u2714"
	} else {
		return "\u2717"
	}
}
