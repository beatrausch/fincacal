package calender

import (
	"encoding/json"
	"fmt"
)

type WeekDay int

const (
	Mon WeekDay = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

var str2wd = map[string]WeekDay{
	"Mon": Mon,
	"Tue": Tue,
	"Wed": Wed,
	"Thu": Thu,
	"Fri": Fri,
	"Sat": Sat,
	"Sun": Sun,
}

var wd2str = []string{
	"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun",
}

func (wd *WeekDay) String() string {
	return wd2str[*wd]
}

func (wd *WeekDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(wd.String())
}

func (wd *WeekDay) UnmarshalJSON(data []byte) (err error) {
	var weekday string
	if err := json.Unmarshal(data, &weekday); err != nil {
		return err
	}
	w, ok := str2wd[weekday]
	if !ok {
		return fmt.Errorf("unknown weekday '%s'", weekday)
	}
	*wd = w
	return nil
}
