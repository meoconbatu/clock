package clock

import "strconv"

const testVersion = 4

type Clock struct {
	Hour, Minute int
}

func New(hour, minute int) Clock {
	var time Clock
	time.Hour = hour
	time.Minute = minute
	return time.CorrectClock()
}
func (time Clock) CorrectClock() Clock {
	time.Hour = (time.Hour + time.Minute/60) % 24
	time.Minute = time.Minute % 60
	if time.Minute < 0 {
		time.Minute = time.Minute + 60
		time.Hour = time.Hour - 1
	}
	if time.Hour < 0 {
		time.Hour = time.Hour + 24
	}
	return time
}

func (time Clock) String() string {
	hourStr := FormatClockElement(time.Hour)
	minuteStr := FormatClockElement(time.Minute)
	return hourStr + ":" + minuteStr
}
func FormatClockElement(element int) string {
	var elementStr string
	if element < 10 {
		elementStr = "0"
	}
	return elementStr + strconv.Itoa(element)
}
func (time Clock) Add(minutes int) Clock {
	time.Minute = time.Minute + minutes
	return time.CorrectClock()
}
