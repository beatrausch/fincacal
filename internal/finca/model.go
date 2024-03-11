package finca

type Night string
type AttendeeName string

type Attendee struct {
	Name   AttendeeName `json:"name"`
	Nights []Night      `json:"nights"`
}

type Accommodation struct {
	Price     float64    `json:"price"`
	Stay      []Night    `json:"stay"`
	Attendees []Attendee `json:"attendees"`
}

func (ac Accommodation) GetOverallNights() int {
	return len(ac.Stay)
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

	OverallPrice     float64
	PricePerNight    float64
	PricePerAttendee map[AttendeeName]Price
}
