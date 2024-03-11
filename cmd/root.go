package cmd

import (
	"flag"
	"github.com/beatrausch/fincacal/internal/config"
	"github.com/beatrausch/fincacal/internal/csv"
	"github.com/beatrausch/fincacal/internal/display"
	"github.com/beatrausch/fincacal/internal/finca"
	"log"
)

var (
	configuration string
	output        string
)

func Execute() {
	flag.StringVar(&configuration, "config", "config.yaml", "Sets the config file.")
	flag.StringVar(&output, "csv", "", "Output CSV filename")

	flag.Parse()

	accommodation, err := config.ReadConfig(configuration)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	summary := finca.CalculateAccommodationPriceSummary(accommodation)
	table := finca.PriceTable(accommodation.Stay, summary.PricePerAttendee)

	display.RenderOverview(summary)
	display.RenderPriceTable(table)

	if output != "" {
		if err := csv.WriteCSV(output, table); err != nil {
			log.Fatalf("Failed to write csv: %v", err)
		}
	}

	//pterm.DefaultTable.WithHasHeader().WithData(priceTable(pricePerPerson, order)).Render()
	//pterm.DefaultBasicText.Printfln("Overall price: %.2f€ %s", overall, check)
}
