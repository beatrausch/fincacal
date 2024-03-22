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
