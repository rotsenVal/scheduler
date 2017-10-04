package types

import "time"

type Schedule struct {
	StartDate time.Time        `json:"startDate"`
	EndDate   time.Time        `json:"endDate"`
	Interval  ScheduleInterval `json:"interval"`
	URL       string           `json:"url"`
}

type ScheduleInterval struct {
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}
