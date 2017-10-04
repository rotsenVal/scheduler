package processors

import (
	"fmt"
	"scheduler/types"
)

type Processor struct {
	scheduleCh chan types.Schedule
	processors []IProcessor
}

func NewProcessor(scheduleCh chan types.Schedule) *Processor {
	processor := Processor{
		scheduleCh: scheduleCh,
		processors: []IProcessor{},
	}

	return &processor
}

func (p *Processor) StartProcessing() {
	go func(pro *Processor) {
		fmt.Println("Start processing...")
		for sched := range pro.scheduleCh {
			for _, proc := range pro.processors {
				proc.Processing(sched)
			}
		}
	}(p)

}

func (p *Processor) AddProcessors(processor ...IProcessor) {
	p.processors = append(p.processors, processor...)
}
