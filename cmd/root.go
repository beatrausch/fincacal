/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
