package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Scheduler struct {
	schedules   map[string][]Schedule
	watching    bool
	ScheduleCh  chan Schedule
	ProcessorCh chan Schedule
}

func NewScheduler() *Scheduler {
	scheduler := Scheduler{
		schedules:   map[string][]Schedule{},
		ScheduleCh:  make(chan Schedule),
		ProcessorCh: make(chan Schedule),
	}
	return &scheduler
}

func (s *Scheduler) WatchForNewSchedules(watch bool) {
	s.watching = watch
	folderToWatch := "./schedules"

	//TODO: remove this infinite loop looking for new files, its cpu intensive
	// for {
	// 	if s.watching {
	fileDescriptor, _ := os.Open(folderToWatch)
	files, _ := fileDescriptor.Readdir(-1)
	for _, fi := range files {
		fileName := fi.Name()
		_, ok := s.schedules[fileName]

		if !ok {
			filePath := folderToWatch + "/" + fileName
			file, err := os.Open(filePath)
			if err != nil {
				//when copying and pasting a file to create duplicates, sometimes there is a slight lock on it
				// and so we will only process a file it we were able to access it
			} else {
				data, _ := ioutil.ReadAll(file)
				file.Close()
				scheduleArr := []Schedule{}
				json.Unmarshal(data, &scheduleArr)

				s.schedules[fileName] = scheduleArr

				for _, sc := range scheduleArr {
					go func(scheduler *Scheduler, sch Schedule) {
						scheduler.ScheduleCh <- sch
					}(s, sc)
				}
			}
		}
	}
	//}
	//}
}

func (s *Scheduler) StartScheduling() {
	go func() {
		now := time.Now()
		for sch := range s.ScheduleCh {
			fmt.Println(sch)
			duration := sch.StartDate.Sub(now)
			if duration < 1 {
				duration = 0
			}

			go func(sche Schedule) {
				time.AfterFunc(duration, func() { startInterval(now, sche, s.ProcessorCh) })
			}(sch)
		}
	}()
}

func startInterval(now time.Time, sch Schedule, proCh chan Schedule) {
	totalSeconds := sch.Interval.Seconds + (sch.Interval.Minutes * 60) + (sch.Interval.Hours * 60 * 60) + (sch.Interval.Days * 24 * 60 * 60)

	ticker := time.NewTicker(time.Duration(totalSeconds) * time.Second)
	quit := make(chan bool)

	endDuration := sch.EndDate.Sub(now)
	if endDuration < 1 {
		endDuration = 0
	}
	time.AfterFunc(endDuration, func() { quit <- true })

	for {
		select {
		case <-ticker.C:
			proCh <- sch
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
