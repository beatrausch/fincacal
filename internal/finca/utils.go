package finca

func personsPerNight(attendees []Attendee) map[Night][]AttendeeName {
	result := make(map[Night][]AttendeeName)

	for _, attendee := range attendees {
		for _, night := range attendee.Nights {
			result[night] = append(result[night], attendee.Name)
		}
	}

	return result
}

func pricePerNightAndPerson(pricePerNight float64, personsPerNight map[Night][]AttendeeName) map[Night]float64 {
	price := make(map[Night]float64)

	for night, participants := range personsPerNight {
		price[night] = pricePerNight / float64(len(participants))
	}

	return price
}

func pricePerPerson(attendees []Attendee, pericePerPersonAndNight map[Night]float64) map[AttendeeName]Price {
	result := make(map[AttendeeName]Price)

	for _, attendee := range attendees {
		var pricePerson Price
		for _, night := range attendee.Nights {
			nightPrice := pericePerPersonAndNight[night]
			pricePerson.Overall = pricePerson.Overall + nightPrice
			pricePerson.OvernightPrices = append(pricePerson.OvernightPrices, OvernightPrice{
				Night: night, Price: nightPrice,
			})
		}
		result[attendee.Name] = pricePerson
	}

	return result
}
