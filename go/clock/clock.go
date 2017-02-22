package clock

import (
	"fmt"
)

const testVersion = 4
const minutesPerHour = 60
const hoursPerDay = 24
const minutesPerDay = minutesPerHour * hoursPerDay

type Clock struct {
	minutes int
}

func New(hour, minute int) Clock {
	var c Clock
	return c.Add(hour*minutesPerHour + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d",
		(c.minutes/60)%24,
		c.minutes%60)
}

func (c Clock) Add(minutes int) Clock {
	c.minutes = (c.minutes + minutes) % minutesPerDay

	c.minutes %= minutesPerDay
	if c.minutes < 0 {
		c.minutes += minutesPerDay
	}
	return c
}
