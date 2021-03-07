package offset

import (
	"time"
)

// Schedule is a collection of events and also provides ways to sort and align them.
type Schedule struct {
	Events []Event
	Holidays []Event
	AnchorEvent string
	AnchorDate time.Time
}

// AddEvent currently does not do any sort of validation on added events.
// However it exists to support that functionality in the future. 
func (schedule *Schedule) AddEvent(event Event) {
	schedule.Events = append(schedule.Events, event)
}

func (schedule *Schedule) AddHoliday(holiday Event) {
	schedule.Holidays = append(schedule.Holidays, holiday)
}

// AnchorTo attempts to set Schedule.AnchorEvent and Schedule.AnchorDate to an Event.Name and Event.PreferedDate
// of an Event in Schedule.Events that has an Event.Name that matches anchor_name.
// If a match is found the function returns true.
// If no match is found the function returns false then Schedule.AnchorEvent and Schedule.AnchorDate are set to
// the coresponding Event.Name and Event.PreferedDate of the Event in Schedule.Events with the earliest Event.PreferedDate.
func (schedule *Schedule) AnchorTo(anchor_name string) (anchor_exists bool) {
	// Create some default values to comapre to
	schedule.AnchorEvent = schedule.Events[0].Name
	schedule.AnchorDate = schedule.Events[0].PreferedDate
	// Return false by default
	anchor_exists = false

	for _, event := range schedule.Events {
		if event.PreferedDate.Before(schedule.AnchorDate) {
			schedule.AnchorEvent = event.Name
			schedule.AnchorDate = event.PreferedDate
		}
		
		if anchor_name == event.Name {
			schedule.AnchorEvent = event.Name
			schedule.AnchorDate = event.PreferedDate
			anchor_exists = true
			return
		}
	}
	return
}

func (schedule *Schedule) BuildSchedule() {

}