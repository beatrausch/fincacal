package finca

import "fmt"

func CalculateAccommodationPriceSummary(accommodation *Accommodation) Summary {
	pricePerNight := accommodation.Price / float64(accommodation.GetOverallNights())

	pricePerAttendee := pricePerPerson(accommodation.Attendees,
		pricePerNightAndPerson(pricePerNight,
			personsPerNight(accommodation.Attendees),
		),
	)

	return Summary{
		Attendees:        len(accommodation.Attendees),
		OverallNights:    len(accommodation.Stay),
		OverallPrice:     accommodation.Price,
		PricePerNight:    pricePerNight,
		PricePerAttendee: pricePerAttendee,
	}
}

func PriceTable(stay []Night, pricePerAttendee map[AttendeeName]Price) [][]string {
	var table [][]string

	header := buildHeader([]string{"Attendee", "Overall"}, stay)
	table = append(table, header)

	for attendee, price := range pricePerAttendee {
		row := make([]string, len(header))
		row[0] = string(attendee)
		row[1] = fmt.Sprintf("%.2f€", price.Overall)

		for _, nightPrice := range price.OvernightPrices {
			for idx, wd := range stay {
				if wd == nightPrice.Night {
					row[idx+2] = fmt.Sprintf("%.2f€", nightPrice.Price)
				}
			}
			row = append(row)

		}
		table = append(table, append(row))
	}
	return table
}

func buildHeader(base []string, stay []Night) []string {
	var nights []string
	for _, night := range stay {
		nights = append(nights, string(night))
	}
	return append(base, nights...)
}
