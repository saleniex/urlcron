package runner

import (
	"github.com/adhocore/gronx/pkg/tasker"
	"urlcron/metric"
	"urlcron/schedule"
)

// Runner runs schedule loaded from loader
type Runner struct {
	loader        schedule.Loader
	metricCounter *metric.Counter
}

// New creates new runner with given schedule.Loader
func New(loader schedule.Loader) *Runner {
	return &Runner{
		loader: loader,
	}
}

// Run loads and starts schedule
func (r Runner) Run() {
	list, err := r.loader.List()
	if err != nil {
		panic(err)
	}
	if len(list) == 0 {
		return
	}

	taskr := tasker.New(tasker.Option{
		Verbose: true,
	})

	for _, item := range list {
		taskr.Task(item.Schedule, item.Exec)
	}

	taskr.Run()
}
