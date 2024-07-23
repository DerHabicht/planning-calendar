package doomsday

import (
	"time"
)

var doomsdaysByCenturyMod4 = []time.Weekday{
	time.Tuesday,
	time.Sunday,
	time.Friday,
	time.Wednesday,
}

// ComputeDoomsday implements part of Conrad's "Doomsday Algorithm" for calculating the day of the week by year.
func ComputeDoomsday(yr int) time.Weekday {
	century := yr / 100
	years := yr % 100

	dd := doomsdaysByCenturyMod4[century%4]

	dozens := years / 12
	years -= dozens * 12
	leapYears := years / 4

	doomsday := time.Weekday((int(dd) + dozens + years + leapYears) % 7)
	return doomsday
}
