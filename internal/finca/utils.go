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
