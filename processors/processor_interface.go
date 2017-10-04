package processors

import "scheduler/types"

type IProcessor interface {
	Processing(schedule types.Schedule)
}
