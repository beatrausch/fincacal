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

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPersonsPerNight(t *testing.T) {
	table := map[string]struct {
		Attendees []Attendee
		Expected  map[Night][]AttendeeName
	}{
		"Mix": {
			Attendees: []Attendee{
				{
					Name:   "Max",
					Nights: []Night{"Mon"},
				},
				{
					Name:   "Jon",
					Nights: []Night{"Mon", "Tue"},
				},
				{
					Name:   "Michael",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
			},
			Expected: map[Night][]AttendeeName{
				"Mon": {"Max", "Jon", "Michael"},
				"Tue": {"Jon", "Michael"},
				"Wed": {"Michael"},
			},
		},
		"Even": {
			Attendees: []Attendee{
				{
					Name:   "Max",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
				{
					Name:   "Jon",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
				{
					Name:   "Michael",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
			},
			Expected: map[Night][]AttendeeName{
				"Mon": {"Max", "Jon", "Michael"},
				"Tue": {"Max", "Jon", "Michael"},
				"Wed": {"Max", "Jon", "Michael"},
			},
		},
	}

	for name, tc := range table {
		t.Run(name, func(t *testing.T) {
			actual := personsPerNight(tc.Attendees)
			if diff := cmp.Diff(tc.Expected, actual); diff != "" {
				t.Errorf("personsPerNight: (-want +got):\n%s\n", diff)
			}
		})
	}
}

func TestPricePerNightAndPerson(t *testing.T) {
	table := map[string]struct {
		PricePerNight     float64
		AttendeesPerNight map[Night][]AttendeeName
		Expected          map[Night]float64
	}{
		"Mixed": {
			PricePerNight: 100,
			AttendeesPerNight: map[Night][]AttendeeName{
				"Mon": {"Max", "Jon", "Michael", "Eva"},
				"Tue": {"Max", "Jon"},
				"Wed": {"Michael"},
			},
			Expected: map[Night]float64{
				"Mon": 25.0,
				"Tue": 50.0,
				"Wed": 100.0,
			},
		},
	}

	for name, tc := range table {
		t.Run(name, func(t *testing.T) {
			actual := pricePerNightAndPerson(tc.PricePerNight, tc.AttendeesPerNight)
			if diff := cmp.Diff(tc.Expected, actual); diff != "" {
				t.Errorf("pricePerNightAndPerson: (-want +got):\n%s\n", diff)
			}
		})
	}
}

func TestPricePerPerson(t *testing.T) {
	table := map[string]struct {
		Attendees     []Attendee
		PricePerNight map[Night]float64
		Expected      map[AttendeeName]Price
	}{
		"Even": {
			Attendees: []Attendee{
				{
					Name:   "Max",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
				{
					Name:   "Jon",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
				{
					Name:   "Michael",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
				{
					Name:   "Eva",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
			},
			PricePerNight: map[Night]float64{
				"Mon": 25.0,
				"Tue": 25.0,
				"Wed": 25.0,
			},
			Expected: map[AttendeeName]Price{
				"Max": {
					Overall: 75.0,
					OvernightPrices: []OvernightPrice{
						{
							Night: "Mon",
							Price: 25.0,
						},
						{
							Night: "Tue",
							Price: 25.0,
						},
						{
							Night: "Wed",
							Price: 25.0,
						},
					},
				},
				"Jon": {
					Overall: 75.0,
					OvernightPrices: []OvernightPrice{
						{
							Night: "Mon",
							Price: 25.0,
						},
						{
							Night: "Tue",
							Price: 25.0,
						},
						{
							Night: "Wed",
							Price: 25.0,
						},
					},
				},
				"Michael": {
					Overall: 75.0,
					OvernightPrices: []OvernightPrice{
						{
							Night: "Mon",
							Price: 25.0,
						},
						{
							Night: "Tue",
							Price: 25.0,
						},
						{
							Night: "Wed",
							Price: 25.0,
						},
					},
				},
				"Eva": {
					Overall: 75.0,
					OvernightPrices: []OvernightPrice{
						{
							Night: "Mon",
							Price: 25.0,
						},
						{
							Night: "Tue",
							Price: 25.0,
						},
						{
							Night: "Wed",
							Price: 25.0,
						},
					},
				},
			},
		},
		"Mixed": {
			Attendees: []Attendee{
				{
					Name:   "Max",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
				{
					Name:   "Jon",
					Nights: []Night{"Mon"},
				},
				{
					Name:   "Michael",
					Nights: []Night{"Mon"},
				},
				{
					Name:   "Eva",
					Nights: []Night{"Mon", "Tue", "Wed"},
				},
			},
			PricePerNight: map[Night]float64{
				"Mon": 25.0,
				"Tue": 50.0,
				"Wed": 50.0,
			},
			Expected: map[AttendeeName]Price{
				"Max": {
					Overall: 125.0,
					OvernightPrices: []OvernightPrice{
						{
							Night: "Mon",
							Price: 25.0,
						},
						{
							Night: "Tue",
							Price: 50.0,
						},
						{
							Night: "Wed",
							Price: 50.0,
						},
					},
				},
				"Jon": {
					Overall: 25.0,
					OvernightPrices: []OvernightPrice{
						{
							Night: "Mon",
							Price: 25.0,
						},
					},
				},
				"Michael": {
					Overall: 25.0,
					OvernightPrices: []OvernightPrice{
						{
							Night: "Mon",
							Price: 25.0,
						},
					},
				},
				"Eva": {
					Overall: 125.0,
					OvernightPrices: []OvernightPrice{
						{
							Night: "Mon",
							Price: 25.0,
						},
						{
							Night: "Tue",
							Price: 50.0,
						},
						{
							Night: "Wed",
							Price: 50.0,
						},
					},
				},
			},
		},
	}

	for name, tc := range table {
		t.Run(name, func(t *testing.T) {
			actual := pricePerPerson(tc.Attendees, tc.PricePerNight)
			if diff := cmp.Diff(tc.Expected, actual); diff != "" {
				t.Errorf("pricePerNightAndPerson: (-want +got):\n%s\n", diff)
			}
		})
	}
}
