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

type Night string
type AttendeeName string

type Attendee struct {
	Name   AttendeeName `json:"name"`
	Nights []Night      `json:"nights"`
}

type Accommodation struct {
	Price     float64    `json:"price"`
	Deposit   float64    `json:"deposit"`
	Stay      []Night    `json:"stay"`
	Attendees []Attendee `json:"attendees"`
}

func (ac Accommodation) GetOverallNights() int {
	return len(ac.Stay)
}

func (ac Accommodation) Remainder() float64 {
	return ac.Price - ac.Deposit
}

type OvernightPrice struct {
	Night Night
	Price float64
}

type Price struct {
	Overall         float64
	OvernightPrices []OvernightPrice
}

type Summary struct {
	Attendees     int
	OverallNights int

	Deposit          float64
	Remainder        float64
	OverallPrice     float64
	PricePerNight    float64
	PricePerAttendee map[AttendeeName]Price
}

func (s Summary) ValidatePrice() bool {
	overall := 0.0
	for _, price := range s.PricePerAttendee {
		overall = overall + price.Overall
	}
	return overall == s.Remainder
}
