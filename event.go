package offset

import (
	"time"
)

// Event defines a "milestone" that is intended to be tracked in relation to other Events.
// It is often associated with another event (ParentEvent).
// The day an Event occurs on is based on the ParentOffset relative to it's ParentEvent.
// So an Event with a ParentOffset of 1 and a ParentEvent with a date of 1/3/2020 would 
// have a calculated date of 1/4/2020.
// Negative values of ParentOffset are acceptable. So using the previous example, 
// a ParentOffset of -1 would result in a calculated date of 1/2/2020
type Event struct {
	Name string
	Length int // In days
	ParentEvent string
	ParentOffset int // In days
	PreferedDate time.Time
	BlackoutDays BlackoutDays
}

// NewEvent is a generic constructor to generate an Event struct.
// Defaults to a 1 day ParrentOffset
func NewEvent(name string, length int, pref_date time.Time) (event Event) {
	event.Name = name
	event.Length = length
	event.PreferedDate = pref_date

	event.ParentOffset = 1
	return
}

// NewWeekdendBlackoutEvent provides functionality to easily generate an Event 
// that already has Saturdays and Sundays completely blacked out (cannot start or overlap).
// An event that cannot take place on weekends is likely to be very common,
// so it makes sense to create a "shortcut" to do that.
func NewWeekdendBlackoutEvent(name string, length int, pref_date time.Time) (event Event) {
	event = NewEvent(name, length, pref_date)
	event.BlackoutDays.Saturday = Blackout{Start:true, Overlap:true}
	event.BlackoutDays.Sunday = Blackout{Start:true, Overlap:true}	
	return
}

// NewWeekdendOverlapOnlyEvent provides functionality to easily generate an Event 
// that cannot start (but can overlap) the weekend.
// An event that cannot take start (but can overlap) weekends is likely to be very common,
// so it makes sense to create a "shortcut" to do that.
func NewWeekdendOverlapOnlyEvent(name string, length int, pref_date time.Time) (event Event) {
	event = NewEvent(name, length, pref_date)
	event.BlackoutDays.Saturday = Blackout{Start:true}
	event.BlackoutDays.Sunday = Blackout{Start:true}	
	return
}

// BlackoutDays is used by the Event type to determine what days an Event cannot take place on.
type BlackoutDays struct {
	Monday Blackout
	Tuesday Blackout
	Wednesday Blackout
	Thursday Blackout
	Friday Blackout
	Saturday Blackout
	Sunday Blackout
}

// Blackout determines exactly what kind of blackout applies to a specific day.
// An event with a BlackoutDays.Saturday.Start = true and BlackoutDays.Saturday.Overlap = false
// means that Event cannot start on a Saturday but can take place on Saturday (assuming it started on a previous day).
type Blackout struct {
	Start bool
	Overlap bool
}