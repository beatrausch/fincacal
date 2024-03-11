package calender

import "fmt"

type IndexedWeekday struct {
	Index   int
	WeekDay WeekDay
}

func IndexedWeekdays(wds []WeekDay) []IndexedWeekday {
	var iwds []IndexedWeekday
	for idx, wd := range wds {
		iwds = append(iwds, IndexedWeekday{
			Index:   idx,
			WeekDay: wd,
		})
	}
	return iwds
}

func (iwd IndexedWeekday) String() string {
	return fmt.Sprintf("%d-%s", iwd.Index, iwd.WeekDay)
}
