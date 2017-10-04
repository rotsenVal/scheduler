package main

import (
	"fmt"
	"scheduler/processors"
	"scheduler/types"
)

func main() {
	scheduler := types.NewScheduler()
	processor := processors.NewProcessor(scheduler.ProcessorCh)

	httpProcessor := processors.HTTPProcessor{Name: "http"}
	//httpsProcessor := processors.HTTPProcessor{Name: "https", IsSSL: true}
	processor.AddProcessors(httpProcessor)
	processor.StartProcessing()

	scheduler.StartScheduling()
	scheduler.WatchForNewSchedules(true)

	fmt.Scanln()
}
